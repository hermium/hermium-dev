package main

import "fmt"
import "net"
import "os"
import "math/rand"
import "strconv"
import "sync"
import "time"

import "hermium/common"
import "hermium/p2p"

type Peer struct {
    Index        uint64
    Conn         net.Conn
    StartAddress common.Address
    EndAddress   common.Address
}

type ClientState struct {
    ListenPort   uint32
    StartAddress common.Address
    EndAddress   common.Address
    Peers        []Peer
    Mux          sync.Mutex
}

func (c *ClientState) InitSettings() {
    c.ListenPort = 1330 + uint32(rand.Intn(10))
    c.Peers = make([]Peer, 0)
    fmt.Println("Set ListenPort to ", c.ListenPort)
}

func (c *ClientState) AddPeer(peer Peer) {
    peer.Index = uint64(len(c.Peers))
    c.Peers = append(c.Peers, peer)
}

func (c *ClientState) RemovePeer(p Peer) {
    c.Peers[len(c.Peers) - 1].Index = p.Index
    c.Peers[p.Index] = c.Peers[len(c.Peers) - 1]
    c.Peers = c.Peers[:len(c.Peers) - 1]
}

func (c *ClientState) HandleConnection(conn net.Conn) {
    defer conn.Close()

    fmt.Println("Connection request from peer: ", conn.RemoteAddr().String())
    peer := Peer{}
    peer.Conn = conn
    
    updateInfoMessage := &p2p.UpdateInfoMessage{
        c.StartAddress,
        c.EndAddress,
    }

    err := p2p.SendUpdateInfoMessage(peer.Conn, updateInfoMessage)

    updateInfoMessage, err = p2p.GetUpdateInfoMessage(peer.Conn)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    c.Mux.Lock()
    c.AddPeer(peer)
    c.Mux.Unlock()

    // TODO: Replace this with peer specific node logic
    for {
        time.Sleep(time.Second)
    }

    c.Mux.Lock()
    c.RemovePeer(peer)
    c.Mux.Unlock()
}

func (c *ClientState) ListenForPeers() {
    ln, err := net.Listen("tcp", ":" + strconv.Itoa(int(c.ListenPort)))
    if err != nil {
        fmt.Println(err.Error())
    }

    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println(err.Error())
            continue
        }

        go c.HandleConnection(conn)
    }
}

func (c *ClientState) ConnectToPeer(address string) {
    conn, err := net.Dial("tcp", address)
    if err != nil {
        fmt.Println(err.Error())
    }

    c.HandleConnection(conn)
}

func (c *ClientState) Run() {
    rand.Seed(time.Now().UnixNano())

    c.InitSettings()

    go c.ListenForPeers()

    if len(os.Args) > 1 {
        go c.ConnectToPeer(os.Args[1])
    }

    // TODO: Replace this with non peer specific node logic
    for {
        time.Sleep(time.Second)
    }
}

func main() {
    c := &ClientState{}
    c.Run()
}
