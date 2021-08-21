package graphql

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
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
  `

type UserInput struct {
	Email  string `json:"email"`
	Status string `json:"status"`
	First  string `json:"first"`
	After  string `json:"first"`
	Last   string `json:"first"`
	before string `json:"first"`
}

type UserParams struct {
	Query     string `json:"query"`
	Variables UserInput
}

type UserOutput struct {
	Data  interface{}
	Error interface{}
}

func UserQuery(userInput *UserInput) (UserOutput, error) {
	// test(userInput)
	userParams := &UserParams{Query: query, Variables: *userInput}
	data, err := json.Marshal(userParams)
	if err != nil {
		panic(err)
	}
	reader := bytes.NewReader(data)
	resp, err := Fetch(reader)
	if err != nil {
		panic(err)
	}
	var userOutput UserOutput
	err = json.Unmarshal(resp, &userOutput)
	fmt.Printf("%#v\n", userOutput)
	return userOutput, err
}

func test(input *UserInput) {
	b, err := json.Marshal(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("JSON Input: %+v\n", b)
	os.Stdout.Write(b)

	output := UserInput{}
	err = json.Unmarshal(b, &output)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nJSON Output: %+v\n", output)
}
