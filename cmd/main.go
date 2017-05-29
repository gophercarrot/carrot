package main

import (
	"carrot"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	c := carrot.CreateSocket("localhost:8000", "ws")
	s := []byte(`{"body":{"code":"i","fileType":"python","line":0,"column":1,"wordToComplete":"i","offset":2}}`)
	c.WriteMessage(websocket.TextMessage, s)
	go func() {
		defer c.Close()
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()
	fmt.Scanln()
}
