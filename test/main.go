package main

import (
	"fmt"

	"github.com/akatebi/gqltst/graphql"
)

type Message struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

func main() {
	userInput := &graphql.UserInput{Email: "me@gmail.com"}
	userOutput, err := graphql.UserQuery(userInput)
	if err != nil {
		panic(err)
	}
	fmt.Printf("userOutput %v", userOutput)
}
