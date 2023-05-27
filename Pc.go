package main

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

// Needs to be tested on a windows machine
func isWindows() bool {
  return exists("C:/")
}

func NewIpInfo() IpInfo {
	str := "{\"ip\":\"74.28.35.18\",\"hostname\":\"modemcable028.32-58-74.mc.videotron.ca\",\"city\":\"Montréal\",\"region\":\"Quebec\",\"country\":\"CA\",\"loc\":\"45.5984,-73.7159\",\"org\":\"AS5769VideotronTelecomLtee\",\"postal\":\"H7M\",\"timezone\":\"America/Toronto\",\"readme\":\"https://ipinfo.io/missingauth\"}"
	stuff := str

	// url := "https://ipinfo.io/json"
	// stuff := strings.TrimSpace(getRequest(url))

	var info IpInfo = IpInfo{}
	err := json.Unmarshal([]byte(stuff), &info)
	if err != nil {
		println("something happened with json fetching of ips and stuff")
	}
	return info
}

func getPcUsername() string {
	username := os.Getenv("USERNAME")
	if username != "" {
		return username
	}
	return "Unknown User"
}
func getNumProcessors() int {
	num, err := strconv.Atoi(os.Getenv("NUMBER_OF_PROCESSORS"))
	if err != nil {
		num = 0
	}
	return num / 2
}
func getEnvVars() map[string]string {
	var i int
	varsArr := os.Environ()
	var envVariables = map[string]string{}
	for i < len(varsArr) {
		environtmentVariable := varsArr[i]
		keyValue := strings.Split(environtmentVariable, "=")
		envVariables[keyValue[0]] = keyValue[1]
		i++
	}
	return envVariables
}

// IpInfo capital letters are required for unmarshal to know its a json key
type IpInfo struct {
	Ip       string
	Hostname string
	City     string
	Region   string
	Country  string
	Loc      string
	Org      string
	Postal   string
	Timezone string
	Readme   string
}

type Vars struct {
	USERDOMAIN_ROAMINGPROFILE       string
	PROCESSOR_LEVEL                 string
	SESSIONNAME                     string
	ALLUSERSPROFILE                 string
	PROCESSOR_ARCHITECTURE          string
	PSModulePath                    string
	SystemDrive                     string
	USERNAME                        string
	FPS_BROWSER_USER_PROFILE_STRING string
	python                          string
	PATHEXT                         string
	DriverData                      string
	GOPATH                          string
	ProgramData                     string
	ProgramW6432                    string
	HOMEPATH                        string
	PROCESSOR_IDENTIFIER            string
	ProgramFiles                    string
	PUBLIC                          string
	windir                          string
	LOCALAPPDATA                    string
	ChocolateyLastPathUpdate        string
	USERDOMAIN                      string
	LOGONSERVER                     string
	FPS_BROWSER_APP_PROFILE_STRING  string
	JAVA_HOME                       string
	OneDrive                        string
	APPDATA                         string
	OS                              string
	COMPUTERNAME                    string
	PROCESSOR_REVISION              string
	CommonProgramW6432              string
	ComSpec                         string
	GO111MODULE                     string
	TEMP                            string
	SystemRoot                      string
	HOMEDRIVE                       string
	USERPROFILE                     string
	TMP                             string
	NUMBER_OF_PROCESSORS            string
	IDEA_INITIAL_DIRECTORY          string
}
