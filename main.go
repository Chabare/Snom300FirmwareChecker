package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var force = flag.Bool("force", false, "Force the download.")

func main() {
	flag.Parse()
	curr := readCurrent()
	oldFirmwareNumber, oldRollupNumber := curr[0], curr[1]

	baseURL := "http://wiki.snom.com/"
	url := baseURL + "Category:Firmware:snom300"
	html := string(establishConnection(url))
	firmware, rollup := getFirmwareAndRollup(html)
	firmwareSiteLink, firmwareNumber := baseURL+firmware[0], firmware[1]
	fmt.Printf("Firmware number: %s ", firmwareNumber)
	firmwareLink := getFirmwareLink(string(establishConnection(firmwareSiteLink)))
	if firmwareNumber != oldFirmwareNumber || *force {
		fmt.Printf("(new)")
		ioutil.WriteFile(firmwareNumber+".bin", establishConnection(firmwareLink), 0644)
	} else {
		fmt.Printf("(old)")
	}
	fmt.Printf("\nFirmware link: %s\n", firmwareLink)

	rollupSiteLink, rollupNumber := baseURL+rollup[0], rollup[1]
	fmt.Printf("Rollup number: %s ", rollupNumber)
	rollupLink := getRollupLink(string(establishConnection(rollupSiteLink)))
	if rollupNumber != oldRollupNumber || *force {
		fmt.Printf("(new)")
		ioutil.WriteFile(rollupNumber+".bin", establishConnection(rollupLink), 0644)
	} else {
		fmt.Printf("(old)")
	}
	fmt.Printf("\nRollup link: %s\n", rollupLink)

	writeCurrent(firmwareNumber, rollupNumber)
}

func getFirmwareAndRollup(html string) ([]string, []string) {
	matches := regexp.MustCompile("Firmware/([^\"]+)").FindAllStringSubmatch(html, 3)

	if len(matches) >= 3 {
		return matches[1], matches[2]
	}

	return []string{""}, []string{""}
}

func getFirmwareLink(html string) string {
	matches := regexp.MustCompile("http://downloads\\.snom\\.com/fw/snom300-[0-9.]+-SIP-f.bin").FindAllStringSubmatch(html, 10)
	if len(matches) == 0 {
		return ""
	}

	return matches[0][0]
}

func getRollupLink(html string) string {
	matches := regexp.MustCompile("http://downloads\\.snom\\.com/fw/mru-preview/snom300-[0-9.]+-SIP-f.bin").FindAllStringSubmatch(html, 10)
	if len(matches) == 0 {
		return ""
	}

	return matches[0][0]
}
