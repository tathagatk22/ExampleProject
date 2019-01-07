package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Email   string `json:"email"`
	Message string `json:"message"`
}

func main() {

	url := "http://192.168.0.63:8080/message"
	var message MessageStruct
	message.Email = "tathagat.khanorkar"
	message.Message = "handleDetailss"
	b, err := json.Marshal(message)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
