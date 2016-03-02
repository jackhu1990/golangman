package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func httpGet() {
	resp, err := http.Get("http://localhost:8001/hello")
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
