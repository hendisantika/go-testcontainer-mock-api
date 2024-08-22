package test_server

import (
	"encoding/json"
	"io"
	"net/http"
)

type AgeResponse struct {
	Count uint64
	Name  string
	Age   int
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	ageResponse := AgeResponse{
		Count: 1000,
		Name:  name,
		Age:   62,
	}
	json, err := json.Marshal(ageResponse)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	io.WriteString(w, string(json))
}
