package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"

	"github.com/interviewstreet/carrot"
)

var msg = []byte(`msg payload`)
var count = 1000
var httpPort = 8900

func main() {

	var host string
	flag.StringVar(&host, "host", "example.com", "Domain name of test host")

	var protocol string
	flag.StringVar(&protocol, "protocol", "wss", "Connection type")

	var request int
	flag.IntVar(&request, "request", 6000, "Total number of requests")

	var writeTime int
	flag.IntVar(&writeTime, "wtime", 100, "number of seconds to wait before writing to websockets")

	var holdTime int
	flag.IntVar(&holdTime, "htime", 30, "number of milliseconds to wait before creating new websocket connection")

	var path string
	flag.StringVar(&path, "path", "/somepath", "Specific url path")

	flag.Parse()

	runtime.GOMAXPROCS(runtime.NumCPU())
	latency := make(chan []float64)
	timeSeries := make(chan []time.Time)
	currentTest := &carrot.Base{host, protocol, request, msg, writeTime, holdTime, path}
	carrot.LoadTest(currentTest, latency, timeSeries)

	data := <-latency
	timeData := <-timeSeries
	fmt.Println(data, timeData)
	fmt.Println("Running HTTP Server, Check /latency route at Port", httpPort)
	carrot.StartHTTPServer("8900", data, timeData)
	fmt.Scanln()
}
