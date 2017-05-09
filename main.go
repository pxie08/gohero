package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
	"net"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":4000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Connected")
	ip, port, _ := net.SplitHostPort(r.RemoteAddr)

	fmt.Println(ip)
	fmt.Println(port)

	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		msgtype, msg, err := socket.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(msg))
		if err = socket.WriteMessage(msgtype, msg); err != nil {
			fmt.Println(err)
			return
		}

	}
}