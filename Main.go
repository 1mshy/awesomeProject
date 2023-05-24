package main

import (
	"encoding/json"
)

func main() {
  timeIt()
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
