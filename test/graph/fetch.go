package graph

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const URL string = "http://localhost:8080/query"

func Fetch(data interface{}) ([]byte, error) {
	client := &http.Client{}
	URL := "http://localhost:8080/query"
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(data)
	req, err := http.NewRequest("POST", URL, payloadBuf)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	return bodyBytes, err
}
