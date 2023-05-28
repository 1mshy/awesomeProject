package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

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

func postData(url string, data string) {
	resp, err := http.Post(url, "application/json",
		bytes.NewBuffer([]byte(data)))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var res map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		println(err.Error())
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
