package main

import (
	"io/ioutil"
	"os/user"
	"path/filepath"
	"strings"
)

func writeCurrent(firmwareNumber, rollupNumber string) {
	usr, _ := user.Current()
	homedir := usr.HomeDir

	data := strings.Join([]string{firmwareNumber, rollupNumber}, "\n")

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
