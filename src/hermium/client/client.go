package main

import "fmt"
import "net"
import "os"
import "path/filepath"
import "strconv"
import "time"

import "hermium/comm"
import "hermium/settings"

const SETTINGS_FILE = ".hermium_client_settings"

type Coordinator struct {
    Connection comm.Connection
    Shards     uint32
}

type ClientState struct {
    ListenPort      uint32
    CoordinatorAddr string
    Coordinator     Coordinator
}

func (c *ClientState) InitSettings() {
    settingsFilePath := filepath.Join(os.Getenv("HOME"), SETTINGS_FILE)

    s := settings.ReadClientSettings(settingsFilePath)

    c.ListenPort = s.ListenPort
    c.CoordinatorAddr = s.CoordinatorAddr

    fmt.Println("Initialized settings")
    fmt.Printf("  Listen Port: %v\n", c.ListenPort)
    fmt.Printf("  Coordinator Address: %v\n", c.CoordinatorAddr)
}

func (c *ClientState) ConnectToCoordinator() {
    conn, err := net.Dial("tcp", c.CoordinatorAddr)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    c.Coordinator.Connection = comm.NewConnection(conn)

    messageType, err := c.Coordinator.Connection.ReadMessageType()
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
    if messageType != comm.POOL_INFO_TYPE {
        fmt.Println("Expected message type %s but received message type %s",
            comm.POOL_INFO_TYPE, messageType)
        os.Exit(1)
    }

    poolInfoMessage, err := c.Coordinator.Connection.ReadPoolInfoMessage()
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    c.Coordinator.Shards = poolInfoMessage.Shards

    fmt.Println("Connected to coordinator")
    fmt.Printf("  Coordinator Address: %v\n", c.Coordinator.Connection.Conn.RemoteAddr().String())
    fmt.Printf("  Shards: %v\n", c.Coordinator.Shards)
}

func (c *ClientState) HandleConnection(conn net.Conn) {
    fmt.Println("Connection request from peer: ", conn.RemoteAddr().String())
}

func (c *ClientState) ListenForPeers() {
    fmt.Println("Listening for peers")

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
    fmt.Println("Started Hermium client")

    c.InitSettings()

    c.ConnectToCoordinator()

    go c.ListenForPeers()

    for {
        time.Sleep(time.Second)
    }
}

func main() {
    c := &ClientState{}
    c.Run()
}
