package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"github.com/heroku/rollbar"
	"strings"
)

var force = flag.Bool("force", false, "Force the download.")

func main() {
	flag.Parse()
	rollbar.SetToken(readToken("token"))
	rollbar.SetServerRoot("github.com/chabare/Snom300FirmwareChecker")
	if *force {
		rollbar.Message(rollbar.WARN, "Forcing download.")
	}

	curr := readCurrent()
	oldFirmwareNumber, oldRollupNumber := curr[0], curr[1]

	baseURL := "http://wiki.snom.com/"
	url := baseURL + "Category:Firmware:snom300"
	html := string(establishConnection(url))
	firmware, rollup := getFirmwareAndRollup(html)
	firmwareSiteLink, firmwareNumber := baseURL+firmware[0], firmware[1]
	fmt.Printf("Firmware number: %s ", firmwareNumber)
	rollbar.Message(rollbar.INFO, "Firmware number: " + firmwareNumber)
	if firmwareNumber != oldFirmwareNumber || *force {
		rollbar.Message(rollbar.INFO, "Found new firmware: " + string(firmwareNumber))
		fmt.Printf("(new)")
		link := getFirmwareLink(string(establishConnection(firmwareSiteLink)))
		ioutil.WriteFile(firmwareNumber+".bin", establishConnection(link), 0644)
	} else {
		fmt.Printf("(old)")
	}
	fmt.Println()

	rollupSiteLink, rollupNumber := baseURL+rollup[0], rollup[1]
	fmt.Printf("Rollup number: %s ", rollupNumber)
	if rollupNumber != oldRollupNumber || *force {
		rollbar.Message(rollbar.INFO, "Found new rollup: " + string(rollupNumber))
		fmt.Printf("(new)")
		link := getRollupLink(string(establishConnection(rollupSiteLink)))
		ioutil.WriteFile(rollupNumber+".bin", establishConnection(link), 0644)
	} else {
		fmt.Printf("(old)")
	}
	fmt.Println()

	writeCurrent(firmwareNumber, rollupNumber)
	rollbar.Wait()
}

func getFirmwareAndRollup(html string) ([]string, []string) {
	matches := regexp.MustCompile("Firmware/([^\"]+)").FindAllStringSubmatch(html, 3)

	if len(matches) >= 3 {
		return matches[1], matches[2]
	}

	rollbar.Message(rollbar.ERR, "No Firmware(and/or) Rollup found")
	return []string{""}, []string{""}
}

func getFirmwareLink(html string) string {
	matches := regexp.MustCompile("http://downloads\\.snom\\.com/fw/snom300-[0-9.]+-SIP-f.bin").FindAllStringSubmatch(html, 10)
	if len(matches) == 0 {
		rollbar.Message(rollbar.ERR, "Couldn't find link to download firmware.")
		return ""
	}

	return matches[0][0]
}

func getRollupLink(html string) string {
	matches := regexp.MustCompile("http://downloads\\.snom\\.com/fw/mru-preview/snom300-[0-9.]+-SIP-f.bin").FindAllStringSubmatch(html, 10)
	if len(matches) == 0 {
		rollbar.Message(rollbar.ERR, "Couldn't find link to download rollup.")
		return ""
	}

	return matches[0][0]
}

func readToken(filename string) string {
	str, _ := ioutil.ReadFile(filename)
	return strings.TrimSpace(string(str))
}
