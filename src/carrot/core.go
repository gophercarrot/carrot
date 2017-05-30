package carrot

import (
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

func receiveMsg(wsconn *websocket.Conn, done chan *Routine, rout *Routine) {
	for {
		_, message, err := wsconn.ReadMessage()
		rout.ReceiveTime = time.Now()
		rout.Diff = rout.ReceiveTime.Sub(rout.SendTime)
		rout.ReceivedMsg = string(message)
		if err != nil {
			log.Println("read:", err)
			return
		}
		done <- rout
	}
}

func writeMsg(wsconn *websocket.Conn, base *Base, rout *Routine) {
	time.Sleep(time.Second * time.Duration(base.Delay))
	rout.SendTime = time.Now()
	wsconn.WriteMessage(websocket.TextMessage, base.Msg)
}

func singleTest(counter *Counter, queue chan *Routine, base *Base, rout *Routine) {
	doneCh := make(chan *Routine)
	conn, err := CreateSocket(base.URL, base.Proto, counter)
	if err != nil {
		return
	}
	go writeMsg(conn, base, rout)
	go receiveMsg(conn, doneCh, rout)
	queue <- <-doneCh
}

func LoadTest(base *Base, latencyCh chan []float64, timeCh chan []time.Time) {

	queue := make(chan *Routine, 1)
	globalCounter := &Counter{0, sync.Mutex{}, 0, 0}
	localCounter := 0

	var latency []float64
	var timeSeries []time.Time

	for range time.Tick(time.Millisecond * time.Duration(base.TickDelay)) {
		routine := &Routine{time.Now(), time.Now(), 0, ""}
		go singleTest(globalCounter, queue, base, routine)
		localCounter++
		if localCounter == base.Count {
			break
		}
	}

	go func() {
		bufferLimit := 0
		for req := range queue {
			latency = append(latency, req.Diff.Seconds()*1000)
			timeSeries = append(timeSeries, req.SendTime)
			bufferLimit++
			if bufferLimit == base.Count {
				latencyCh <- latency
				timeCh <- timeSeries
			}
		}
	}()

}
