// API to get address by zipcode
package main

import (
	"net/http"

	"github.com/kthanasio/golang/cep/controller"
)

func main() {
	http.HandleFunc("/", controller.CEPController)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
