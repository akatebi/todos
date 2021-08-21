package graphql

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func Fetch(bytes io.Reader) ([]byte, error) {
	URL := "http://localhost:8080/query"
	req, err := http.NewRequest("POST", URL, bytes)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{
		Transport: nil,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
		},
		Jar:     nil,
		Timeout: 0,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	return bodyBytes, err
}
