package main

import (
	"fmt"
	"io/ioutil"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	curr := readCurrent()
	oldFirmwareNumber, oldRollupNumber := curr[0], curr[1]

	baseURL := "http://wiki.snom.com/"
	url := baseURL + "Category:Firmware:snom300"
	html := string(establishConnection(url))
	firmware, rollup := getFirmwareAndRollup(html)
	firmwareSiteLink, firmwareNumber := baseURL+firmware[0], firmware[1]
	fmt.Printf("Firmware number: %s\n", firmwareNumber)
	if firmwareNumber != oldFirmwareNumber {
		link := getFirmwareLink(string(establishConnection(firmwareSiteLink)))
		ioutil.WriteFile(firmwareNumber+".bin", establishConnection(link), 0644)
	}

	rollupSiteLink, rollupNumber := baseURL+rollup[0], rollup[1]
	fmt.Printf("Rollup number: %s\n", rollupNumber)
	if rollupNumber != oldRollupNumber {
		link := getRollupLink(string(establishConnection(rollupSiteLink)))
		ioutil.WriteFile(rollupNumber+".bin", establishConnection(link), 0644)
	}

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

func writeCurrent(firmwareNumber, rollupNumber string) {
	usr, _ := user.Current()
	homedir := usr.HomeDir

	data := firmwareNumber + "\n"
	data += rollupNumber

	ioutil.WriteFile(filepath.Join(homedir, ".snom"), []byte(data), 0644)
}

func readCurrent() []string {
	usr, _ := user.Current()
	homedir := usr.HomeDir

	data, _ := ioutil.ReadFile(filepath.Join(homedir, ".snom"))

	numbers := strings.Split(string(data), "\n")
	if len(numbers) < 2 {
		return []string{"", ""}
	}

	return numbers
}
