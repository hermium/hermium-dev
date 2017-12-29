package main

import "fmt"
import "time"
import "strconv"
import "errors"

import "crypto/sha256"
import "encoding/hex"

type Chain struct {
	/* Struct for the Chain type. 
	VARIABLES
		chain : an array of pointers to blocks in the chain
	*/
	chain []*Block
	length int64
}

type Block struct {
	/* Struct for the Block type. 
	VARIABLES
		index : the height of the block in the chain
		timestamp : the timestamp at which the block was created
		data : the contents of the block
		hash : the SHA256 created from data
		prev : the SHA256 of the previous block
	*/
	index int64
	time int64
	data string
	hash string
	prev string
}

func newChain() *Chain {
	/* Generates a new Chain object with a genesis Block. */
	c := new(Chain)
	genesisBlock := Block{0, time.Now().Unix(), "First block", "", ""}
	genesisBlock.hash = genesisBlock.getHash()
	c.chain = append(c.chain, &genesisBlock)
	c.length = 1
	return c
}

func (c Chain) getPrevBlock() *Block {
	/* Returns the last Block in the current Chain */
	return c.chain[c.length-1]
}

func newBlock(c Chain, data string) *Block {
	/* Generates a new Block object given a Chain and a data string. */
	b := new(Block)
	p := c.getPrevBlock()
	b.index = c.length
	b.time = time.Now().Unix()
	b.data = data
	b.prev = p.hash
	b.hash = b.getHash()
	return b
}

func (b Block) getHash() string {
	/* Generates the SHA256 hash of the Block. */
	indexStr := strconv.FormatInt(b.index, 10)
	timeStr := strconv.FormatInt(b.time, 10)
	s := indexStr + timeStr + b.data + b.prev
    h := sha256.New()
    h.Write([]byte(s))
    o := hex.EncodeToString(h.Sum(nil))
	return o
}

func (c Chain) isValidBlock(b Block) (bool,error) {
	/* Checks if a new Block is valid w.r.t a Chain c. */
	p := c.getPrevBlock()
	if p.index + 1 != b.index {
		return false, errors.New("New block has incorrect index")
	} 
	if p.hash != b.prev {
		return false, errors.New("New block has incorrect previous hash")
	}
	if b.hash != b.getHash() {
		return false, errors.New("New block has incorrect hash")
	}
	return true, nil
}

func (c Chain) push(b Block) *Chain {
	/* Pushes a new Block onto the Chain. */
	valid, err := c.isValidBlock(b)
	if err != nil {
		panic(err)
	}
	if valid {
		c.chain = append(c.chain, &b)
		c.length = c.length + 1
	}
	return &c
}

/* TODOS:
1) Write a func that checks if a Chain is valid
2) We need to maintain a global Chain -- not sure how to do this in Go.
3) Communicating the global Chain across {coordinator,client} nodes
*/

func main() {
	/* Dumbass test to make sure everything compiles. */
	BLOCKCHAIN := newChain()
	fmt.Println(BLOCKCHAIN.chain)
	BLOCK := newBlock(*BLOCKCHAIN, "Nilai")
	BLOCKCHAIN = BLOCKCHAIN.push(*BLOCK)
	fmt.Println(BLOCKCHAIN.chain)


}