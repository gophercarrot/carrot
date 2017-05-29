package carrot

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

func receiveMsg(wsconn *websocket.Conn, done chan string) {
	for {
		_, message, err := wsconn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		done <- string(message)
	}
}

func writeMsg(wsconn *websocket.Conn, tm int, msg []byte, delay int) {
	time.Sleep(time.Second * time.Duration(delay))
	wsconn.WriteMessage(tm, msg)
}

func LoadTest(url string, proto string, count int, msg []byte, delay int) {
	//latency := make([]int, count)

	var wg sync.WaitGroup
	queue := make(chan string, 1)

	wg.Add(count)

	for i := 0; i < count; i++ {
		go func() {
			doneCh := make(chan string)
			conn := CreateSocket(url, proto)
			go writeMsg(conn, websocket.TextMessage, msg, delay)
			go receiveMsg(conn, doneCh)
			queue <- <-doneCh
		}()
	}

	go func() {
		for req := range queue {
			fmt.Println(req)
		}
		wg.Done()
	}()
	wg.Wait()
}
