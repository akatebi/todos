package graphql

import (
	"encoding/json"
	"fmt"
)

const query string = `
query user($email: String!, $status: Status, $first: Int, $after:String) {
	user(email: $email) {
	  id
	  email
	  completedCount
	  totalCount
	  todos(first: $first, status:$status, after:$after) {
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
	Data struct {
		User User `json:"user"`
	}
	// Data  interface{}
	Error interface{}
}

func UserQuery(userInput *UserInput) (User, error) {
	//
	userParams := &UserParams{Query: query, Variables: *userInput}
	resp := Fetch(userParams)
	fmt.Printf("\nresp: %+v\n", string(resp))
	output := UserOutput{}
	err := json.Unmarshal(resp, &output)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\noutput: %+v\n", output)
	return output.Data.User, nil
}
