package main

import (
	"errors"
	"os/exec"
	"strings"
)

func getChromeVersion() (string, error) {
	cmd := exec.Command("wmic datafile where name=\"C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe\" get Version /value")
	out, err := cmd.Output()
	if err != nil {
		return "", errors.New("Chrome is not present on this computer")
	}
	return strings.Split(string(out), "=")[1], nil
}
