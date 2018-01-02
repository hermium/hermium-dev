package main

import "fmt"
import "net"
import "os"
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
    c.ListenPort = 1337
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

func (c *ClientState) Run() {
    c.InitSettings()

    go c.ListenForPeers()

    for {
        time.Sleep(time.Second)
    }
}

func main() {
    c := &ClientState{}
    c.Run()
}
