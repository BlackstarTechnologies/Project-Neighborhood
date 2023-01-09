package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Version1Get(url string) {
	// "https://api.github.com/users/KrunalLathiya"
	resp, err := http.Get(url)
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	fmt.Print(string(body))
}
