package main

import (
	"carrot"
	"fmt"
	"runtime"
	"time"
)

var msg = []byte(`{"body":{"code":"i","fileType":"python","line":0,"column":1,"wordToComplete":"i","offset":2}}`)
var count = 1000
var httpPort = 8900

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Running HTTP Server at", httpPort)
	latency := make(chan []float64)
	timeSeries := make(chan []time.Time)

	//currentTest := &carrot.Base{"localhost:8000", "ws", 600, msg, 2}
	currentTest := &carrot.Base{"autosuggest.hackerrank.com", "wss", 1000, msg, 2, 30}
	carrot.LoadTest(currentTest, latency, timeSeries)

	data := <-latency
	timeData := <-timeSeries
	fmt.Println(data, timeData)
	carrot.StartHTTPServer("8900", data, timeData)
	fmt.Scanln()
}
