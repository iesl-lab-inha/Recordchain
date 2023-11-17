package main

import (
	"fmt"
	"main/cli"
	"main/core"
	"strconv"
	"time"
)

func main() {
	cli := cli.CLI{}
	cli.Run()
	// testBCC()
}

func testBCC() {
	bc := core.NewBlockchain()
	var block *core.Block

	bc.NewRecord([]byte("AAABBB"), false, time.Now().Unix()+85)   // 85 seconds
	bc.NewRecord([]byte("BBBCCC"), false, time.Now().Unix()+95)   // 95 seconds
	bc.NewRecord([]byte("BBBDDD"), false, time.Now().Unix()+900)  // 15 minutes
	bc.NewRecord([]byte("AAAIII"), false, time.Now().Unix()+3600) // 1 hour
	bc.NewRecord([]byte("AAAJJJ"), false, time.Now().Unix()+90)   // 90 seconds
	bc.NewRecord([]byte("AAAKKK"), false, time.Now().Unix()+1800) // 30 minutes
	bc.NewRecord([]byte("AAAHHH"), false, time.Now().Unix()+1200) // 20 minutes
	bc.NewRecord([]byte("AAATTT"), false, time.Now().Unix()+100)  // 100 seconds
	block = bc.NewBlock()
	bc.AddBlock(block)

	bc.NewRecord([]byte("AAACCC"), false, time.Now().Unix()+90)   // 90 seconds
	bc.NewRecord([]byte("CCCDDD"), false, time.Now().Unix()+1800) // 30 minutes
	bc.NewRecord([]byte("DDDFFF"), false, time.Now().Unix()+20)   // 20 seconds
	bc.NewRecord([]byte("DDDFFF"), false, time.Now().Unix()+3600) // 1 hour
	block = bc.NewBlock()
	bc.AddBlock(block)

	bc.NewRecord([]byte("AAAEEE"), false, time.Now().Unix()+30)   // 30 seconds
	bc.NewRecord([]byte("AAAGGG"), false, time.Now().Unix()+1800) // 30 minutes
	bc.NewRecord([]byte("AAAHHH"), false, time.Now().Unix()+1200) // 20 minutes
	block = bc.NewBlock()
	bc.AddBlock(block)

	block = bc.NewBlock()
	bc.AddBlock(block)

	for _, block := range bc.Chain {
		fmt.Printf("\n\nPrev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Records size: %d\n", len(block.Records))
		fmt.Printf("Hash: %x\n", block.Hash)
		//pow := core.NewProofOfWork(block)
		//fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		//fmt.Println()
	}
}

func testPOW() {
	bc := core.NewBlockchain()
	var block *core.Block

	block = bc.NewBlock()
	bc.AddBlock(block)
	bc.AddBlock(block)

	for _, block := range bc.Chain {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Mempool size: %d\n", len(block.Records))
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := core.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
