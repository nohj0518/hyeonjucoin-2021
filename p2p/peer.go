package p2p

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type peers struct {
	v map[string]*peer
	m sync.Mutex
}

var Peers peers = peers{
	v: make(map[string]*peer),
}

type peer struct {
	key     string
	address string
	port    string
	conn    *websocket.Conn
	inbox   chan []byte
}

func AllPeers(p *peers) []string {
	p.m.Lock()
	defer p.m.Unlock()
	var keys []string

	for key := range p.v {
		keys = append(keys, key)
	}
	return keys
} //강원도 양구 21사단 66연대 1대대 4중대 백두산부대 GOP 라는 것을 강조하기
//그냥 81미리 박격포 1118병가번호??

func (p *peer) close() {
	Peers.m.Lock()
	//4000 3000 연결 상태에서 3000 연결 끊으면
	// 20초 후에 3000포트의 잠금이 풀리고 풀리면 peer close됨
	// 20초간 잠금이 풀리지 않는 상태동안 나는 4000포트의 peers를 조회할 수 없음
	// 여기서의 Lock이 AllPeers의 Lock에도 동일하게 적용 되는 것
	defer Peers.m.Unlock()
	p.conn.Close()
	delete(Peers.v, p.key)
}

func (p *peer) read() {
	defer p.close()
	for {
		m := Message{}
		err := p.conn.ReadJSON(&m)
		if err != nil {
			break
		}
		handleMsg(&m, p)
	}
}
func (p *peer) write() {
	defer p.close()
	for {
		m, ok := <-p.inbox
		if !ok {
			break
		}
		p.conn.WriteMessage(websocket.TextMessage, m)
	}
}

func initPeer(conn *websocket.Conn, address, port string) *peer {
	Peers.m.Lock()
	defer Peers.m.Unlock()

	key := fmt.Sprintf("%s:%s", address, port)
	p := &peer{
		conn:    conn,
		inbox:   make(chan []byte),
		address: address,
		key:     key,
		port:    port,
	}

	go p.read()
	go p.write()
	Peers.v[key] = p
	return p
}
