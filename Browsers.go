package main

import "os"

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}
func existsFromC(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}
func existsFromProgramFiles(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}
