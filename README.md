#Carrot
Distributed WebSocket and HTTP Load Testing Framework in Go

### Highly Concurrent Load Testing Framework

## Setup
Properly setup $GOPATH

## Fetch Dependencies

```
go get github.com/gorilla/websocket
go get github.com/wcharczuk/go-chart
```

## Fetch Carrot

```
go get -v github.com/interviewstreet/carrot
```

## Running

Go to $GOPATH/go/src/github.com/interviewstreet/carrot and run

```
go run cmd/main.go
```

## Config at Run time
```
go run main.go -host=example.com -protocol=wss -htime=40 -request=5000 -wtime=1 -htime=30 -path=/somepath

wtime -> number of seconds to wait before writing to websockets
htime -> number of milliseconds to wait before creating new websocket connection
```

## Parameters in main.go
```	
currentTest := &carrot.Base{"example.com", "wss", 1000, msg, 2, 30}

wss -> protocol use 'ws' for localhost
1000 -> number of requests
msg -> payload to send
2 -> number of seconds to wait before writing to websockets
30 -> number of milliseconds to wait before creating new websocket connection
```

### After stats you should see screenshot which can be looked at http://localhost:8900/latency

![ScreenShot](https://github.com/interviewstreet/Carrot/blob/master/latency.png)
