package main

import (
	"flag"

	"github.com/feel-easy/go-example/termui/ui"
)

var (
	serverIP   string
	serverPort string
)

func init() {
	flag.StringVar(&serverIP, "ip", "127.0.0.1", "设置服务器的IP地址")
	flag.StringVar(&serverPort, "port", "8888", "设置服务器的连接端口号")
}

func main() {
	flag.Parse()
	maxChanSize := 10000

	//log.SetLevel(log.DebugLevel)
	msgIn := make(chan ui.Message, maxChanSize)
	msgOut := make(chan ui.MessageOut, maxChanSize)
	autoReply := make(chan int, maxChanSize)
	closeChan := make(chan int, maxChanSize)

	layout := ui.NewLayout("张三", "aa", msgIn, msgOut, closeChan, autoReply, nil)
	layout.Init()
}
