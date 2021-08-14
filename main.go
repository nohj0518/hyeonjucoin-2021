package main

import (
	"fmt"

	"github.com/nohj0518/hyeonjucoin-2021/blockchain"
)

func main() {
	fmt.Println("4.5 Refactoring part Two!")
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	chain.AddBlock("Fourth Block")
	for _, block := range chain.AllBlocks() {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
	}
}
