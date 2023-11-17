package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

type Record struct {
	ID        []byte
	Hash      []byte
	Status    []byte
	Expire    int64
	Signature []byte
}

// SetID sets ID of a record
func (tx *Record) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

// Serialize returns a serialized Transaction
func (tx Record) Serialize() []byte {
	var encoded bytes.Buffer

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	return encoded.Bytes()
}

// DeserializeTransaction deserializes a transaction
func DeserializeRecord(data []byte) Record {
	var record Record

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&record)
	if err != nil {
		log.Panic(err)
	}

	return record
}
