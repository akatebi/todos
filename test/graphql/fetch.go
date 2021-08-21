package graphql

import (
	"io"
	"io/ioutil"
	"net/http"
)

func Fetch(reader io.Reader) ([]byte, error) {
	URL := "http://localhost:8080/query"
	req, err := http.NewRequest("POST", URL, reader)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	return bodyBytes, err
}
