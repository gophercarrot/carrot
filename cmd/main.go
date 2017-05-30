package main

import (
	"carrot"
	"fmt"
	"runtime"
)

var msg = []byte(`{"body":{"code":"i","fileType":"python","line":0,"column":1,"wordToComplete":"i","offset":2}}`)
var count = 1000
var httpPort = 8900

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Running HTTP Server at", httpPort)
	// for i := 0; i <= 20; i++ {
	// 	go runSockets()
	// }
	currentTest := &carrot.Base{"localhost:8000", "ws", 10, msg, 2}
	//currentTest1 := &carrot.Base{"autosuggest.hackerrank.com", "wss", 10000, msg, 50000}
	carrot.LoadTest(currentTest)
	carrot.StartHTTPServer("8900")
	fmt.Scanln()
}
