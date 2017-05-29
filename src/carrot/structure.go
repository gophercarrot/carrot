package carrot

import (
	"github.com/gorilla/websocket"
)

type Completion struct {
	Conn    *websocket.Conn
	Count   int
	Latency []int
}
