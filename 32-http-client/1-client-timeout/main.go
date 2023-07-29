package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	c := http.Client{
		Timeout: time.Second,
	}
	resp, err := c.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf(string(body))
}
