package main

import (
	"github.com/webmamba/my-crypto/blockchain"
	"fmt"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("Second")
	chain.AddBlock("Third")
	chain.AddBlock("LOL")

	for i, block := range chain.Blocks {
		fmt.Printf("Block number: %d\n", i)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Timestamp: %s\n", block.Time)
		fmt.Printf("Validation: %t\n", block.Validate())
		fmt.Printf("----------------\n")
	}
}