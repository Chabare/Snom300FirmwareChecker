package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func establishConnection(u string) []byte {
	res, err := http.Get(u)

	var body []byte

	if err != nil {
		log.Print(err)
		return []byte("101")
	}

	// read body
	body, err = ioutil.ReadAll(res.Body)

	res.Body.Close()

	if err != nil {
		log.Print(err)
		return []byte("101")
	}

	if res.StatusCode != 200 {
		if res.StatusCode == 404 {
			return []byte("404")
		}
		if res.StatusCode == 403 {
			return []byte("403")
		}
		log.Printf("Unexpected status code %d from %s", res.StatusCode, u)
		return []byte("101")
	}

	return body
}
