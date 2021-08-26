package graph

import (
	"encoding/json"
	"log"

	"github.com/akatebi/gqltst/graph/model"
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
	After  string `json:"after"`
	Before string `json:"before"`
	First  int    `json:"first"`
	Last   int    `json:"last"`
}

type UserParams struct {
	Query     string `json:"query"`
	Variables UserInput
}

type UserResponse struct {
	Data struct {
		User model.User `json:"user"`
	}
	// Data  interface{}
	Error interface{}
}

func UserQuery(userInput *UserInput) (*UserResponse, error) {
	//
	userParams := &UserParams{Query: query, Variables: *userInput}
	resp, err := Fetch(userParams)
	if err != nil {
		return nil, err
	}
	log.Printf("resp: %+v\n\n", string(resp))
	userResponse := &UserResponse{}
	err = json.Unmarshal(resp, userResponse)
	if err != nil {
		return nil, err
	}
	return userResponse, nil
}
