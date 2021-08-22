package graphql

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const query string = `
query user($email: String!, $status: Status, $first: Int, $after: String, $last: Int, $before: String) {
	user(email: $email) {
	  id
	  email
	  completedCount
	  totalCount
	  todos(status: $status, first: $first, after: $after, last: $last, before: $before) {
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

type UserInput struct {
	Email  string `json:"email"`
	Status string `json:"status"`
	First  int    `json:"first"`
	After  string `json:"after"`
	Last   int    `json:"last"`
	Before string `json:"before"`
}

type UserParams struct {
	Query     string `json:"query"`
	Variables UserInput
}

type UserOutput struct {
	Data  interface{}
	Error interface{}
}

func UserQuery(userInput *UserInput) {
	//
	userParams := &UserParams{Query: query, Variables: *userInput}
	client := &http.Client{}
	// req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	// body := bytes.NewBuffer([]byte(graphql))
	URL := "http://localhost:8080/query"
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(userParams)
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
	fmt.Printf("%v", string(bodyBytes))
}

// func test(input *UserInput) {
// 	b, err := json.Marshal(input)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("JSON Input: %+v\n", b)
// 	os.Stdout.Write(b)

// 	output := UserInput{}
// 	err = json.Unmarshal(b, &output)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("\nJSON Output: %+v\n", output)
// }
