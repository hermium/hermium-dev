package common

const BlockHashSize = 32
const MerkleHashSize = 32

type BlockDifficulty float64

type BlockHash [BlockHashSize]byte

type BlockNonce uint64

type BlockWeight float64

type MerkleHash [MerkleHashSize]byte

type Timestamp uint64

type BlockHeader struct {
    PrevBlockHash BlockHash
    Timestamp     Timestamp
    Difficulty    BlockDifficulty
    PrevAvgWeight BlockWeight
    MerkleRoot    MerkleHash
    StateCommit   MerkleHash
    Nonce         BlockNonce
}

type BlockEntry struct {
    Address      Address
    Transactions []Transaction
}

type BlockStubbedMerkle struct {
    StartPos    uint64
    EndPos      uint64
    NumLeaves   uint64
    Entries     []BlockEntry      
    LeftHashes  MerkleHash
    RightHashes MerkleHash
}
