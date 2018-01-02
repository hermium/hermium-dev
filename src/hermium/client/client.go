package main

import "fmt"
import "net"
import "os"
import "math/rand"
import "strconv"
import "time"

import "hermium/common"

type Peer struct {
    Conn         net.Conn
    Version      string
    StartAddress common.Address
    EndAddress   common.Address
}

type ClientState struct {
    ListenPort uint32
    Peers      []Peer
}

func (c *ClientState) InitSettings() {
    c.ListenPort = 1330 + uint32(rand.Intn(10))
    fmt.Println("Set ListenPort to ", c.ListenPort)
}

func (c *ClientState) HandleConnection(conn net.Conn) {
    fmt.Println("Connection request from peer: ", conn.RemoteAddr().String())
    peer := Peer{}
    peer.Conn = conn

    // TODO: exchange UpdateInfoMessage with new peer and add to c.Peers
}

func (c *ClientState) ListenForPeers() {
    ln, err := net.Listen("tcp", ":" + strconv.Itoa(int(c.ListenPort)))
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
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
    // TODO
}

func (c *ClientState) Run() {
    rand.Seed(time.Now().UnixNano())

    c.InitSettings()

    go c.ListenForPeers()

    if len(os.Args) > 1 {
        c.ConnectToPeer(os.Args[1])
    }

    for {
        time.Sleep(time.Second)
    }
}

func main() {
    c := &ClientState{}
    c.Run()
}
