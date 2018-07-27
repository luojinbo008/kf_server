package main

import (
	"fmt"
	"net/http"
	"kf_server/source"
)

var manager = client.Manager{
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	clients:    make(map[*Client]bool),
}

func main() {
	fmt.Println("Starting application...")

	go manager.start()

	http.HandleFunc("/ws", client.WsPage)
	http.ListenAndServe(":12345", nil)
}