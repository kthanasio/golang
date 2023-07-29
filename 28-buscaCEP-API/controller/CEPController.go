package controller

import (
	"28-buscaCEP-API/service"
	"encoding/json"
	"net/http"
)

// CEPController will handle new requests
func CEPController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{ "error": "Route not found!"}`))
		return
	}
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.Write([]byte(`{ "error": "Missing parameter!"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, e := service.CEPService(cepParam)

	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "error": "Something wrent wrong!"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
