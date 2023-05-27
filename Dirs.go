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

func appdataPath() string {
	return os.Getenv("APPDATA")
}
func startMenuPath() string {
	return appdataPath() + "\\Microsoft\\Windows\\Start Menu\\Programs"
}

func readStartMenuProgs() {
	//entries, err := os.ReadDir(appdataPath())
}
