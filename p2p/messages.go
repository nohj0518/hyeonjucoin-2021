package p2p

import (
	"encoding/json"
	"fmt"

	"github.com/nohj0518/hyeonjucoin-2021/blockchain"
	"github.com/nohj0518/hyeonjucoin-2021/utils"
)

type MessageKind int

const (
	MessageNewestBlock MessageKind = iota
	MessageAllBlocksRequest
	MessageAllBlocksResponse
)

type Message struct {
	Kind    MessageKind
	Payload []byte
}

func makeMessage(kind MessageKind, payload interface{}) []byte {
	m := Message{
		Kind:    kind,
		Payload: utils.ToJSON(payload),
	}
	return utils.ToJSON(m) // JSON byte로 변환(Marshal)
}

func sendNewestBlock(p *peer) {
	fmt.Printf("Sending newest block to %s\n", p.key)
	b, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
	utils.HandleErr(err)
	m := makeMessage(MessageNewestBlock, b)
	p.inbox <- m
}

func requestAllBlocks(p *peer) {
	m := makeMessage(MessageAllBlocksRequest, nil)
	p.inbox <- m
}

func sendAllBlocks(p *peer) {
	m := makeMessage(MessageAllBlocksResponse, blockchain.Blocks(blockchain.Blockchain()))
	p.inbox <- m
}

func handleMsg(m *Message, p *peer) {
	switch m.Kind {
	case MessageNewestBlock:
		fmt.Printf("Received the newest block from %s\n", p.key)
		var payload blockchain.Block
		utils.HandleErr(json.Unmarshal(m.Payload, &payload))
		b, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
		utils.HandleErr(err)
		if payload.Height >= b.Height {
			fmt.Printf("Requesting all blocks from %s\n", p.key)
			// request all the blocks from 4000
			requestAllBlocks(p)
		} else {
			// send 4000 our block
			sendNewestBlock(p)
		}
	case MessageAllBlocksRequest:
		fmt.Printf("%s wants all the blocks\n", p.key)
		// 4000 번으로 모든 block 을 달라고 Request 요청함
		sendAllBlocks(p)
	case MessageAllBlocksResponse:
		fmt.Printf("Received all the blocks from %s\n", p.key)
		// 4000 번으로 부터 모든 block 을 Response 받음
		var payload []*blockchain.Block
		utils.HandleErr(json.Unmarshal(m.Payload, &payload))
	}
}
