package main

import "os"

func getPcUsername() string {
	return os.Getenv("USERNAME")
}

func getIp() string {
  url := "https://ipinfo.io/json"
  stuff := getRequest(url)
  println(stuff)
  return stuff
}
