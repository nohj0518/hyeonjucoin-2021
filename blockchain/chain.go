package blockchain

import (
	"sync"
)

type blockchain struct {
	// blocks []*Block
	NewestHash string `json:"newestHash"`
	Height 	   int    `json:"height"`
}

var b *blockchain
var once sync.Once

func (n *blockchain) AddBlock(data string){
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash =block.Hash
	b.Height = block.Height
}


func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"",0}
			b.AddBlock("Genesis")
		})
	}
	return b
}
