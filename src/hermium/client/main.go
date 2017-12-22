package main

import "fmt"
import "net"
import "os"
import "strconv"
import "time"

import utils "hermium/utils"

const LISTEN_PORT = 1337

func main() {
    fmt.Println("Started Hermium client")

    go listenForPeers()

    for {
        time.Sleep(time.Second)
    }
}

func listenForPeers() {
    fmt.Println("Listening for peers")

    ln, err := net.Listen("tcp", ":" + strconv.Itoa(LISTEN_PORT))
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println(err)
            continue
        }

        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    fmt.Println("Connected to client: ", conn.RemoteAddr().String())
    client := utils.Client{conn}
    fmt.Println(client)
}
