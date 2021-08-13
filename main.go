package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

func main() {
	fmt.Println("4.1 Our Fist Block!")
	genesisBlock := block{"Genesis Block", "", ""} // First Block
	hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash))
	hexHash := fmt.Sprintf("%x", hash)
	genesisBlock.hash = hexHash

	fmt.Print(genesisBlock)
}
