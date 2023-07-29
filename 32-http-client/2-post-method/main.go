package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	jsonData := bytes.NewBuffer([]byte(`{ "name": "kleber" }`))
	resp, err := c.Post("https://64c58173c853c26efadade93.mockapi.io/api/users", "application/json", jsonData)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// fmt.Printf(string(body))
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
