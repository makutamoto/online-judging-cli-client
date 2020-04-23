package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type statusType struct {
	WholeResult int    `json:"whole_result"`
	Result      int    `json:"result"`
	Time        int64  `json:"time"`
	Memory      int64  `json:"memory"`
	CurrentCase int    `json:"current_case"`
	WholeCase   int    `json:"whole_case"`
	Description string `json:"description"`
}

func sendData(data []byte) {
	var result resultType
	var status statusType
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
		if err := json.Unmarshal(bytes, &status); err != nil {
			log.Fatalln(err)
		}
		res := resultType(status.Result)
		fmt.Printf("%d/%d %v %dms %dkb\n", status.CurrentCase, status.WholeCase, res, status.Time, status.Memory)
		result.update(res)
	}
	fmt.Println(result)
	if resultType(status.WholeResult) == resultCompileError {
		fmt.Println(status.Description)
	}
}
