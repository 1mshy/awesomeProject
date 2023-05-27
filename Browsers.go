package main

import ("os"
       "os/exec"
       "errors"
       "strings"
       )

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

func getChromeVersion() (string, error) {
  cmd := exec.Command("wmic datafile where name=\"C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe\" get Version /value")
  out, err := cmd.Output()
  if err != nil {
    return "", errors.New("Chrome is not present on this computer")
  }
  return strings.Split(string(out), "=")[1], nil
}