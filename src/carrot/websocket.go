package carrot

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

//unused for later purpose
type completion struct {
	body struct {
		code           string
		fileType       string
		line           int
		column         int
		wordToComplete string
		offset         int
	}
}

func getAddr(addr string) string {
	return addr
}

// CreateSocket returns a socket instance
func CreateSocket(addr string, urlProto string) *websocket.Conn {
	wsaddr := url.URL{Scheme: urlProto, Host: getAddr(addr), Path: "/"}

	c, _, err := websocket.DefaultDialer.Dial(wsaddr.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	return c
}
