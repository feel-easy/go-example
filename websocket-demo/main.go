package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)
var upgrader = websocket.Upgrader{}

type Message struct {
	Message string `json:"message"`
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	http.HandleFunc("/ws", handleConnections)
	go handleMessages()
	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

//注册成为 websocket
func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	clients[ws] = true
	//不断的从页面上获取数据 然后广播发送出去
	for {
		// 发送数据
		time.Sleep(time.Second * 3)
		msg := Message{Message: "这是向页面发送的数据 " + time.Now().Format("2006-01-02 15:04:05")}
		broadcast <- msg
		// 接收数据
		var msgRecv Message
		err := ws.ReadJSON(&msgRecv)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		log.Println(msgRecv)
	}
}

//广播发送至页面
func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("client.WriteJSON error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
