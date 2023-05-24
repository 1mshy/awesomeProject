package main

import "os"

func getPcUsername() string {
	return os.Getenv("USERNAME")
}
