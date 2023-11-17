package core

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"main/util"
	"time"
)

// BCC represents a proof-of-work
type BCC struct {
	block        *Block
	startTime    int64
	rc           *Recordchain
	minedRecords []*Record
}

// NewBCC builds and returns a BCC
func NewBCC(b *Block, rc *Recordchain) *BCC {
	bcc := BCC{b, 0, rc, []*Record{}}
	prevBlock := rc.Chain[len(rc.Chain)-1]
	bcc.startTime = prevBlock.Timestamp
	return &bcc
}

func (pow *BCC) parepareData() []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.HashRecords(),
			util.IntToHex(pow.block.Timestamp),
			util.IntToHex(int64(pow.block.Index)),
		},
		[]byte{},
	)

	return data
}

func (bcc *BCC) traverseChain() {
	for _, b := range bcc.rc.Chain {
		for _, r := range b.Records {
			if r.Expire <= time.Now().Unix() {
				checked := false
				for _, rx := range bcc.minedRecords {
					if bytes.Compare(rx.Hash, r.Hash) == 0 {
						checked = true
					}
				}
				if !checked {
					rx := bcc.rc.NewRecord(r.Hash, true, r.Expire)
					bcc.minedRecords = append(bcc.minedRecords, rx)
					fmt.Printf("Found: %x\n", rx.ID)
				}
			}
		}
	}
}

func (bcc *BCC) mineBlock() []byte {
	bcc.block.Records = append(bcc.block.Records, bcc.minedRecords...)
	bcc.block.Timestamp = time.Now().Unix()
	data := bcc.parepareData()
	hash := sha256.Sum256(data)
	fmt.Printf("Hash: %x\n", hash)
	return hash[:]
}

// Run performs a BCC
func (bcc *BCC) Run() []byte {

	fmt.Printf("\n\n*** Mining the block %d containing %d records ***\n", bcc.block.Index, len(bcc.block.Records))

	fmt.Printf("Step 1: Traversing expired records\n")
	countingTime := time.Now().Unix()
	for bcc.startTime+10 > time.Now().Unix() {
		bcc.traverseChain()
	}
	fmt.Printf("Time passed: %d seconds\n\n", time.Now().Unix()-countingTime)

	fmt.Printf("Step 2: Competition\n")
	countingTime = time.Now().Unix()
	for bcc.startTime+60 > time.Now().Unix() {
		if len(bcc.minedRecords) > 0 {
			fmt.Printf("Time passed: %d seconds\n\n", time.Now().Unix()-countingTime)
			fmt.Printf("Step 3: Mining\n")
			return bcc.mineBlock()
		} else {
			bcc.traverseChain()
		}
	}
	fmt.Printf("Time passed: %d seconds\n\n", time.Now().Unix()-countingTime)
	fmt.Printf("Step 3: Mining\n")
	return bcc.mineBlock()
}
