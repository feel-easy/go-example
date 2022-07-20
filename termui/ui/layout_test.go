package ui

import (
	"flag"
	"testing"
)

func Test_UI(t *testing.T) {
	flag.Parse()
	maxChanSize := 10000

	//log.SetLevel(log.DebugLevel)
	msgIn := make(chan Message, maxChanSize)
	msgOut := make(chan MessageOut, maxChanSize)
	autoReply := make(chan int, maxChanSize)
	closeChan := make(chan int, maxChanSize)

	layout := NewLayout("", "", msgIn, msgOut, closeChan, autoReply, nil)
	layout.Init()
}
