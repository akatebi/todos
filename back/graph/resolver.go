package graph

import (
	"encoding/json"
	"log"

	"github.com/akatebi/todos/graph/model"
	"github.com/graphql-go/relay"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver ...
type Resolver struct {
	users map[string]*model.User
	todos map[string][]*model.Todo
}

var usersData = make(map[string]*model.User)
var todosData = make(map[string][]*model.Todo)

func initUser(userID string, todos []*model.Todo) {
	log.Printf("##### userID = %v", userID)
	ID := relay.ToGlobalID("User", userID)
	log.Printf("##### ID = %v", ID)
	usersData[userID] = &model.User{
		ID:             ID,
		UserID:         userID,
		TotalCount:     len(todos),
		CompletedCount: 0,
	}
	todosData[userID] = todos
	for _, todo := range todos {
		if todo.Complete == true {
			usersData[userID].CompletedCount++
		}
	}
}

// Initialize ...
func Initialize() *Resolver {
	todos := []*model.Todo{
		{
			ID:       newID(),
			Text:     "Taste JavaScript",
			Complete: true,
		},
		{
			ID:       newID(),
			Text:     "Buy a unicorn",
			Complete: false,
		},
	}
	initUser("me@gmail.com", todos)
	return &Resolver{usersData, todosData}
}

var idCounter int

func newID() string {
	idCounter++
	json, err := json.Marshal(idCounter)
	if err != nil {
		panic(err)
	}
	return string(json)
}
