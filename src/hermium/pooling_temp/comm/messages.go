package comm

import "encoding/binary"
import "fmt"

type MessageType uint8
const (
    POOL_INFO_TYPE MessageType = iota
)

type PoolInfoMessage struct {
    Shards uint32
}

func (conn *Connection) ReadMessageType() (MessageType, error) {
    m, err := conn.Reader.ReadByte()
    messageType := MessageType(m)
    if err != nil {
        fmt.Println(err.Error())
        return messageType, err
    }

    return messageType, nil
}

func (conn *Connection) ReadPoolInfoMessage() (PoolInfoMessage, error) {
    bytes := make([]byte, 4)
    _, err := conn.Reader.Read(bytes)

    if err != nil {
        fmt.Println(err.Error())
        return PoolInfoMessage{}, err
    }

    shards := binary.BigEndian.Uint32(bytes)
    return PoolInfoMessage{shards}, nil
}
