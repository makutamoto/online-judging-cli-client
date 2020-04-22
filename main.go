package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 6 {
		log.Fatalln("Expected arguments: lang, code, testcases, limit, accuracy")
	}
	limit, err := strconv.Atoi(os.Args[4])
	if err != nil {
		log.Fatalln(err)
	}
	accuracy, err := strconv.Atoi(os.Args[5])
	json := makeJSON(os.Args[1], os.Args[2], os.Args[3], limit, accuracy)
	sendData(json)
}
