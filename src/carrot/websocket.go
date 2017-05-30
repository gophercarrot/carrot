package carrot

import (
	"fmt"
	"net/url"

	"github.com/gorilla/websocket"
)

//unused for later purpose
type completion struct {
	body struct {
		code           string
		fileType       string
		line           int
		column         int
		wordToComplete string
		offset         int
	}
}

func getAddr(addr string) string {
	return addr
}

// CreateSocket returns a socket instance
func CreateSocket(addr string, urlProto string, counter *Counter) (*websocket.Conn, error) {
	wsaddr := url.URL{Scheme: urlProto, Host: getAddr(addr), Path: "/"}
	c, _, err := websocket.DefaultDialer.Dial(wsaddr.String(), nil)
	counter.Increment() // increment global counter
	if err != nil {
		//log.Fatal("dial:", err)
		fmt.Println("Broken WebSocket Conn:", counter.val)
		counter.Failure()
	} else {
		fmt.Println("Created WebSocket Conn:", counter.val)
		counter.Success()
	}
	fmt.Println("Success and Failures", counter.success, counter.failure)
	return c, err
}
