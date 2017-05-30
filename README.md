# Carrot
Distributed WebSocket and XHR Load Testing Framework in Go

### Highly Concurrent Load Testing Framework
Lots of fan in Go Routines that runs to provide better load details

## Setup
Add repo to gopath

```
GOPATH=$GOPATH:$PWD
go get github.com/gorilla/websocket
go get github.com/wcharczuk/go-chart
```

## Running
```
go run cmd/main.go
```

## Parameters in main.go
```	
currentTest := &carrot.Base{"autosuggest.hackerrank.com", "wss", 1000, msg, 2, 30}

1000 -> number of requests
msg -> payload to send
2 -> number of seconds to wait before writing to websockets
30 -> number of milliseconds to wait before creating new websocket connection
```

## After stats you should see screenshot which can be looked at http://localhost:8900/latency

![ScreenShot](https://github.com/interviewstreet/Carrot/blob/master/latency.png)
