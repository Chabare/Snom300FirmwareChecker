package main

import (
	"os/user"
	"io/ioutil"
	"path/filepath"
	"strings"
)

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
