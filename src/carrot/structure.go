package carrot

import (
	"time"
)

type Base struct {
	URL, Proto string
	Count      int
	Msg        []byte
	Delay      int
	TickDelay  int
}

type Routine struct {
	SendTime    time.Time
	ReceiveTime time.Time
	Diff        time.Duration // milliseconds
	ReceivedMsg string
}
