package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
				text
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
		Query     string    `json:"query"`
		Variables Variables `json:"variables"`
	}
	body := &GraphQL{Query: query, Variables: Variables{Email: "me@gmail.com"}}
	fmt.Println("Calling API...")
	client := &http.Client{}
	// req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	// body := bytes.NewBuffer([]byte(graphql))
	URL := "http://localhost:8080/query"
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(body)
	req, err := http.NewRequest("POST", URL, payloadBuf)
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
	// type GraphQL struct {
	// 	Data model.User
	// 	Error error
	// }

	type GraphQLResp struct {
		Data  interface{}
		Error interface{}
	}
	// var res interface{}
	var res GraphQLResp
	json.Unmarshal(bodyBytes, &res)
	fmt.Printf("%#v\n", res)
}
