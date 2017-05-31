package main

import (
	"carrot"
	"fmt"
	"runtime"
	"time"
)

var msg = []byte(`msg payload`)
var count = 1000
var httpPort = 8900

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	latency := make(chan []float64)
	timeSeries := make(chan []time.Time)

	currentTest := &carrot.Base{"example.com", "wss", 6000, msg, 100, 30, "/somepath"}
	carrot.LoadTest(currentTest, latency, timeSeries)

	data := <-latency
	timeData := <-timeSeries
	fmt.Println(data, timeData)
	fmt.Println("Running HTTP Server, Check /latency route at Port", httpPort)
	carrot.StartHTTPServer("8900", data, timeData)
	fmt.Scanln()
}
