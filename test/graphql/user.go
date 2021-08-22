package graphql

import (
	"encoding/json"
	"fmt"
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
	resp := Fetch(userParams)
	output := UserOutput{}
	err := json.Unmarshal(resp, &output)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nJSON Output: %+v\n", output)
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
