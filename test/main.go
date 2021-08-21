package main

import (
	"encoding/json"
	"fmt"
	"os"

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

func test() {
	input := Message{First: "Karmen", Last: "Rapouchi"}
	b, err := json.Marshal(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("JSON Input: %+v\n", b)
	os.Stdout.Write(b)

	output := Message{}
	err = json.Unmarshal(b, &output)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nJSON Output: %+v\n", output)
}
