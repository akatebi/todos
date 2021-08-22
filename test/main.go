package main

import (
	"github.com/akatebi/gqltst/graphql"
)

type Message struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

func main() {
	userInput := &graphql.UserInput{Email: "me@gmail.com", Status: "ANY", First: 100}
	graphql.UserQuery(userInput)
}
