package main

import (
	"log"

	"github.com/akatebi/gqltst/graphql"
)

type Message struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

func main() {
	userInput := &graphql.UserInput{Email: "me@gmail.com", Status: "ANY", First: 100}
	user, err := graphql.UserQuery(userInput)
	if err != nil {
		panic(err)
	}
	log.Printf("UserId %v", user.ID)
}
