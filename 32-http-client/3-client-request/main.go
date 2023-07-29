package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	c := http.Client{}
	req, err := http.NewRequest(
		"GET",
		"https://google.com",
		nil,
	)
	req.Header.Set("Accept", "application/json")
	if err != nil {
		panic(err)
	}
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf(string(body))
}
