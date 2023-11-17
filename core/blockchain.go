package core

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"main/util"
	"time"
)

type Recordchain struct {
	Chain   []*Block
	Mempool []*Record
}

// NewRecord creates a new record
func (rc *Recordchain) NewRecord(hash []byte, delete bool, expire int64) *Record {
	var status []byte
	if delete {
		status = bytes.Repeat([]byte{0}, 32)
	} else {
		status = bytes.Repeat([]byte{1}, 32)
	}
	rx := Record{nil, hash, status, expire, nil}
	rx.SetID()
	rc.Mempool = append(rc.Mempool, &rx)
	return &rx
}

// NewBlock creates and returns Block
func (rc *Recordchain) NewBlockPOW(prevBlockHash []byte, index int) *Block {
	block := &Block{time.Now().Unix(), rc.Mempool, prevBlockHash, []byte{}, index, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	rc.Mempool = []*Record{}
	return block
}

// AddBlock saves provided data as a block in the blockchain
func (rc *Recordchain) AddBlockPOW(data string) {
	prevBlock := rc.Chain[len(rc.Chain)-1]
	newBlock := rc.NewBlockPOW(prevBlock.Hash, prevBlock.Index+1)
	rc.Chain = append(rc.Chain, newBlock)
}

// NewBlock creates and returns Block
func (rc *Recordchain) NewBlock() *Block {
	block := &Block{time.Now().Unix(), rc.Mempool, nil, []byte{}, 0, 0}

	rc.Mempool = []*Record{}
	return block
}

// AddBlock saves provided data as a block in the blockchain
func (rc *Recordchain) AddBlock(block *Block) {
	prevBlock := rc.Chain[len(rc.Chain)-1]
	block.PrevBlockHash = prevBlock.Hash
	block.Index = prevBlock.Index + 1

	bcc := NewBCC(block, rc)
	hash := bcc.Run()
	block.Hash = hash[:]
	// newBlock := rc.NewBlock(prevBlock.Hash, prevBlock.Index+1)
	rc.Chain = append(rc.Chain, block)
}

// NewBlockchain creates a new Blockchain with genesis Block
func NewBlockchain() *Recordchain {
	rc := &Recordchain{
		Chain:   []*Block{},
		Mempool: []*Record{},
	}

	genesis := &Block{Timestamp: time.Now().Unix(), PrevBlockHash: bytes.Repeat([]byte{0}, 32), Index: 1}
	data := bytes.Join(
		[][]byte{
			genesis.PrevBlockHash,
			genesis.HashRecords(),
			util.IntToHex(genesis.Timestamp),
			util.IntToHex(int64(genesis.Index)),
		},
		[]byte{},
	)
	hash := sha256.Sum256(data)
	genesis.Hash = hash[:]
	rc.Chain = append(rc.Chain, genesis)

	fmt.Printf("Creating a block. Chain size: %d\n", len(rc.Chain))

	return rc
}

// GetBlockHashes returns a list of hashes of all the blocks in the chain
func (rc *Recordchain) GetBlockHashes() [][]byte {
	var blocks [][]byte

	for _, b := range rc.Chain {
		blocks = append(blocks, b.Hash)
	}

	return blocks
}
