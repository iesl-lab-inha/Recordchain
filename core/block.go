package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

// Block keeps block headers
type Block struct {
	Timestamp     int64
	Records       []*Record
	PrevBlockHash []byte
	Hash          []byte
	Index         int
	Nonce         int
}

// HashRecords returns a hash of the transactions in the block
func (b *Block) HashRecords() []byte {
	var rxHashes [][]byte
	var rxHash [32]byte

	for _, rx := range b.Records {
		rxHashes = append(rxHashes, rx.Hash)
	}
	rxHash = sha256.Sum256(bytes.Join(rxHashes, []byte{}))

	return rxHash[:]
}

// Serialize serializes the block
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// DeserializeBlock deserializes a block
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}
