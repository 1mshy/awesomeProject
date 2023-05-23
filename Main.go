package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func main() {
	rep, err := http.Get("http://10.0.0.32:8080/discord")
	DoErr(&err)
	body, err := ioutil.ReadAll(rep.Body)
	DoErr(&err)
	ss := string(body)
	dojson(ss)

	println()
	var i int
	for i < 1000 {
		isPrime(i)
		i++
	}
	println()
}

func dojson(str string) {
	var result map[string]string
	err := json.Unmarshal([]byte(str), &result)
	DoErr(&err)

	for key, value := range result {
		println(key, ": ", value)
	}
}

func isPalindrome(x int) bool {

	var reversed int
	temp := x
	for temp > 0 {
		reversed = reversed*10 + temp%10
		temp /= 10
	}
	return reversed == x
}

func DoErr(err *error) {
	if *err != nil {
		print((*err).Error())
	}
}
