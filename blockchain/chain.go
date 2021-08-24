package blockchain

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"sync"

	"github.com/nohj0518/hyeonjucoin-2021/db"
	"github.com/nohj0518/hyeonjucoin-2021/utils"
)
type blockchain struct {
	// blocks []*Block
	NewestHash string `json:"newestHash"`
	Height 	   int    `json:"height"`
}
var b *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte){
	decoder := gob.NewDecoder(bytes.NewReader(data))
	utils.HandleErr(decoder.Decode(b))
}

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string){
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash =block.Hash
	b.Height = block.Height
	b.persist()
}

func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"",0}
			fmt.Printf("*Newest Hash: %s\n*Height: %d", b.NewestHash, b.Height)
			// search for check point on the db
			checkpoint := db.Checkpoint() 
			if checkpoint == nil {
				b.AddBlock("Genesis")
			}else{
				// restore b from bytes
				fmt.Println("\nRestoring...")
				b.restore(checkpoint)
			}
		})
	}
	fmt.Printf("Newest Hash: %s\n Height: %d", b.NewestHash, b.Height)
	return b
}
