package main

import (
	"log"

	"github.com/gorilla/websocket"
)

func sendData(data []byte) {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:7867/submit", nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
		log.Fatalln(err)
	}
	for {
		_, bytes, err := conn.ReadMessage()
		if err != nil {
			break
		}
		log.Println(string(bytes))
	}
}
