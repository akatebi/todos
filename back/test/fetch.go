package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/akatebi/todos/graph/model"
)

func main() {
	query := `query User($email: String!) {
        user(email: $email) {
          id
          email
          completedCount
          totalCount
          todos(first: 100) {
            edges {
              cursor
              node {
                id
                complete
                __typename
              }
            }
            pageInfo {
              endCursor
              hasNextPage
            }
          }
        }
      }`
	type Variables struct {
		Email string `json:"email"`
	}
	type GraphQL struct {
		query      string `json:"query"`
		variables  Variables `json:"email"`
	}
	variables := []{email: "me@gmail.com"}
	fmt.Println("Calling API...")
	client := &http.Client{}
	// req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	URL := "http://localhost:8080/query"
	body := bytes.NewBuffer([]byte("HEY"))
	req, err := http.NewRequest("POST", URL, body)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	// type Response struct {
	// 	ID     string `json:"id"`
	// 	Joke   string `json:"joke"`
	// 	Status int    `json:"status"`
	// }
	// var responseObject Response
	// json.Unmarshal(bodyBytes, &responseObject)
	// fmt.Printf("API Response as struct %+v\n", responseObject)

	user := &model.User{}
	json.Unmarshal(bodyBytes, user)
	fmt.Printf("User %+v\n", user)
}
