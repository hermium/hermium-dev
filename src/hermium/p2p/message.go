package p2p

import "net"

import "hermium/common"

type MessageType uint64

const (
    UpdateInfoType MessageType = iota
    NodesGetType
    NodesSendType
    MempoolGetType
    MempoolSendType
    HeadersGetType
    HeadersSendType
    BlockGetType
    BlockSendType
)

type UpdateInfoMessage struct {
    StartAddress common.Address
    EndAddress   common.Address
}

type NodesGetMessage struct {
    StartAddress common.Address
    EndAddress   common.Address
    MaxCount     uint64
}

type NodesSendMessage struct {
    Nodes []net.TCPAddr
}

type MempoolGetMessage struct {
    StartAddress common.Address
    EndAddress   common.Address
}

type MempoolSendMessage struct {
    Transactions []common.Transaction
}

type HeadersGetMessage struct {
    StartHash      common.BlockHash
    MaxCount       uint64
}

type HeadersSendMessage struct {
    Headers []common.BlockHeader
}

type BlockGetMessage struct {
    Hash         common.BlockHash
    StartAddress common.Address
    EndAddress   common.Address
}

type BlockSendMessage struct {
    Header        common.BlockHeader
    StubbedMerkle common.BlockStubbedMerkle
}
