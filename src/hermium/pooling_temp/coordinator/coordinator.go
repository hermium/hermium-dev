package main

import "encoding/binary"
import "fmt"
import "net"
import "os"
import "path/filepath"
import "strconv"
import "time"

import "hermium/comm"
import "hermium/settings"

const SETTINGS_FILE = ".hermium_coordinator_settings"

type CoordinatorState struct {
    ListenPort uint32
    Shards     uint32
}

func (c *CoordinatorState) InitSettings() {
    settingsFilePath := filepath.Join(os.Getenv("HOME"), SETTINGS_FILE)

    s := settings.ReadCoordinatorSettings(settingsFilePath)

    c.ListenPort = s.ListenPort
    c.Shards = s.Shards

    fmt.Println("Initialized settings")
    fmt.Printf("  Listen Port: %v\n", c.ListenPort)
    fmt.Printf("  Shards: %v\n", c.Shards)
}

func (c *CoordinatorState) HandleConnection(conn net.Conn) {
    fmt.Println("Connection request from client: ", conn.RemoteAddr().String())
    connection := comm.NewConnection(conn)
    connection.Writer.WriteByte(byte(comm.POOL_INFO_TYPE))
    bs := make([]byte, 4)
    binary.BigEndian.PutUint32(bs, c.Shards)
    connection.Writer.Write(bs)
    connection.Writer.Flush()
}


func (c *CoordinatorState) ListenForClients() {
    fmt.Println("Listening for clients")

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

func (c *CoordinatorState) Run() {
    fmt.Println("Started Hermium coordinator")

    c.InitSettings()

    go c.ListenForClients()

    for {
        time.Sleep(time.Second)
    }
}

func main() {
    c := &CoordinatorState{}
    c.Run()
}
