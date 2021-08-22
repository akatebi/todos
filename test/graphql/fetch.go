package graphql

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const URL string = "http://localhost:8080/query"

func Fetch(data interface{}) (*http.Response, error) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(data)
	resp, err := http.Post(URL, "application/json; charset=utf-8", b)
	if err != nil {
		panic(err)
	}
	return resp, nil
}
