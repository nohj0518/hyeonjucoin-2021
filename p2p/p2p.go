package p2p

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/nohj0518/hyeonjucoin-2021/utils"
)

var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	_, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)
}
