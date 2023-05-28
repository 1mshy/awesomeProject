package main

import (
	"fmt"
	"reflect"
	"strings"
)

const url = "https://discord.com/api/webhooks/1111739259255275555/3TdXYkyqFyEIwBph1di8dPZG5C6mVKQn1YoL_jeITZZ5yGqGJYJKYikxlwecFsrsqcMd"

func webhookRaw(message string) {

	values := map[string]string{"content": message, "username": "why"}
	postJson(url, values)
}
func webhook(data string) {
	postData(url, data)
}

type DiscordMessage struct {
	content string
	embeds  Embeds
	footer  Footer
	author  Author
}

type Embed struct {
	title       string
	description string
	url         string
	footer      Footer
}
type Embeds struct {
	embeds []Embed
}

type Footer struct {
	icon_url string
	text     string
}
type Author struct {
	name     string
	url      string
	icon_url string
}

// tostrings
func (e Embed) String() string {
	stringed := iterOver(e, 1)
	if stringed == "{}" {
		return ""
	}
	return stringed
}

func (f Footer) String() string {
	stringed := iterOver(f, 0)
	if stringed == "{}" {
		return ""
	}
	return stringed
}
func (a Author) String() string {
	stringed := iterOver(a, 0)
	if stringed == "{}" {
		return ""
	}
	return stringed
}
func (e Embeds) String() string {
	sb := new(strings.Builder)
	sb.WriteString("[")
	firstItteration := true
	for _, embed := range e.embeds {
		if !firstItteration {
			sb.WriteString(",")
		} else {
			firstItteration = false
		}
		sb.WriteString(embed.String())
	}
	sb.WriteString("]")
	return sb.String()
}
func (d DiscordMessage) String() string {
	sb := new(strings.Builder)
	embedStr := d.embeds.String()
	footerStr := d.footer.String()
	authorStr := d.author.String()
	sb.WriteString("{")
	sb.WriteString(fmt.Sprintf("\"content\":\"%s\"", d.content))
	if embedStr != "" {
		sb.WriteString(fmt.Sprintf(",\"embeds\":%s", embedStr))
	}
	if footerStr != "" {
		sb.WriteString(fmt.Sprintf(",\"footer\":%s", footerStr))
	}
	if authorStr != "" {
		sb.WriteString(fmt.Sprintf(",\"author\":%s", authorStr))
	}
	sb.WriteString("}")
	return sb.String()
}

// end tostrings

// shorted cuts off from the last value
func iterOver(e interface{}, shorted int) string {
	sb := new(strings.Builder)
	v := reflect.ValueOf(e)
	typeOf := v.Type()
	firstIteration := true

	sb.WriteString("{")
	for i := 0; i < v.NumField()-shorted; i++ {
		currentTitle := typeOf.Field(i).Name
		currentValue := v.Field(i).String()
		if currentValue == "" {
			continue
		}
		if !firstIteration {
			sb.WriteString(",")
		}
		sb.WriteString("\"")
		sb.WriteString(currentTitle)
		sb.WriteString("\"")
		sb.WriteString(":")
		sb.WriteString("\"")
		sb.WriteString(currentValue)
		sb.WriteString("\"")

		if firstIteration {
			firstIteration = false
		}
	}
	sb.WriteString("}")
	stringed := sb.String()
	return stringed
}
