package main

import (
	"encoding/json"
	"os"
)

func getPcUsername() string {
	return os.Getenv("USERNAME")
}

func getIp() string {
	url := "https://ipinfo.io/json"
	stuff := getRequest(url)
	var js map[string]string
	err := json.Unmarshal([]byte(stuff), &js)
	if err != nil {
		return ""
	}
	return js["ip"]
}
