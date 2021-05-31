package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func main() {
	go Sender()
	time.Sleep(180 * time.Second)
}

func Sender() {
	count := 1
	ticker := time.NewTicker(5 * time.Second)
	for {
		var data = []byte(`{"value":"94"}`)
		req, err := http.NewRequest("POST", "http://localhost:8000/rate", bytes.NewBuffer(data))
		client := http.DefaultClient
		resp, err := client.Do(req)
		fmt.Println("resp:", count, resp, err)
		if count%5 == 0 {
			time.Sleep(30 * time.Second)
		}
		count++
		<-ticker.C
	}
}
