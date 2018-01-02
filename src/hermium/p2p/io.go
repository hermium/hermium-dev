package p2p

import "encoding/binary"
import "errors"
import "net"

func SendUpdateInfoMessage(conn net.Conn, m *UpdateInfoMessage) error {
    b := make([]byte, 8)
    binary.BigEndian.PutUint64(b, uint64(UpdateInfoType))
    conn.Write(b)
    conn.Write(m.StartAddress[:])
    conn.Write(m.EndAddress[:])

    return nil
}

func GetUpdateInfoMessage(conn net.Conn) (*UpdateInfoMessage, error) {
    buf := make([]byte, 8)
    _, err := conn.Read(buf)
    if err != nil {
        return nil, err
    }
    messageType := MessageType(binary.BigEndian.Uint64(buf))
    if messageType != UpdateInfoType {
        return nil, errors.New("Incorrect message type")
    }

    updateInfoMessage := &UpdateInfoMessage{}

    _, err = conn.Read(updateInfoMessage.StartAddress[:])
    if err != nil {
        return nil, err
    }
    _, err = conn.Read(updateInfoMessage.EndAddress[:])
    if err != nil {
        return nil, err
    }

    return updateInfoMessage, nil
}
