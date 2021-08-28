package main

import (
	"log"

	"github.com/akatebi/gqltst/graph"
)

type Message struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

func Main2() {
	userInput := &graph.UserInput{Email: "me@gmail.com", Status: "ANY", First: 100}
	resp, err := graph.UserQuery(userInput)
	if err != nil {
		panic(err)
	}
	log.Printf("### %+v", resp)
}
