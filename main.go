package main

import (
	"fmt"
	"regexp"
)

func main() {
	url := "http://wiki.snom.com/Category:Firmware:snom300"
	html := string(establishConnection(url))
	firmware, rollup := getFirmwareAndRollup(html)
	fmt.Printf("Firmware: %s\nMaintenance: %s\n", firmware, rollup)
}

func getFirmwareAndRollup(html string) (string, string) {
	matches := regexp.MustCompile("/Firmware/([^\"]+)").FindAllStringSubmatch(html, 3)

	if len(matches) >= 3 {
		return matches[1][1], matches[2][1]
	}

	return "", ""
}
