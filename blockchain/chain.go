package blockchain

import (
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
	utils.FromBytes(b, data)
}

func (b *blockchain) persist() {
	db.SaveCheckpoint(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string){
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash =block.Hash
	b.Height = block.Height
	b.persist()
}

func (b *blockchain) Blocks() []*Block{
	var blocks []*Block
	hashCursor := b.NewestHash
	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)
		if block.PrevHash != ""{
			hashCursor = block.PrevHash
		} else {
			break
		}
	}
	return blocks
}

func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"",0}
			// search for check point on the db
			checkpoint := db.Checkpoint() 
			if checkpoint == nil {
				b.AddBlock("Genesis")
			}else{
				// restore b from bytes
				b.restore(checkpoint)
			}
		})
	}
	fmt.Println(b.NewestHash)
	return b
}
