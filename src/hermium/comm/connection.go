package comm

import "bufio"
import "net"

type Connection struct {
    Conn   net.Conn
    Reader *bufio.Reader
    Writer *bufio.Writer
}

func NewConnection(conn net.Conn) Connection {
    return Connection{conn, bufio.NewReader(conn), bufio.NewWriter(conn)}
}

func (conn *Connection) Close() {
    // TODO: close connection
}
