package main

import (
	"carrot"
	"fmt"
	"log"
	"runtime"

	"github.com/gorilla/websocket"
)

var msg = []byte(`{"body":{"code":"i","fileType":"python","line":0,"column":1,"wordToComplete":"i","offset":2}}`)
var count = 1000

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	latency := make([]int, count)

	conn := carrot.CreateSocket("autosuggest.hackerrank.com", "wss")
	iface := carrot.Completion{conn, 0, latency}

	iface.Conn.WriteMessage(websocket.TextMessage, msg)

	go func() {
		for {
			_, message, err := iface.Conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()
	fmt.Scanln()
}
