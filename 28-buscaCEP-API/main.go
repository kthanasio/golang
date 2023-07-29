// API to get address by zipcode
package main

import (
	"28-buscaCEP-API/controller"
	"net/http"
)

func main() {
	http.HandleFunc("/", controller.CEPController)
	http.ListenAndServe(":8080", nil)
}
