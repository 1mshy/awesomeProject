package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func webhook(message string) {
	url := "https://discord.com/api/webhooks/1111739259255275555/3TdXYkyqFyEIwBph1di8dPZG5C6mVKQn1YoL_jeITZZ5yGqGJYJKYikxlwecFsrsqcMd"

	values := map[string]string{"content": message, "username": "why"}
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

func getRequest(url string) string {
	response, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	str := string(body)
	return str
}
