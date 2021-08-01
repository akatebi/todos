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

func (r *Resolver) initUser(userID string, todos []*model.Todo) {
	log.Printf("## initUser ## userID = %v", userID)
	ID := relay.ToGlobalID("User", userID)
	log.Printf("## initUser ## ID = %v", ID)
	if r.users[ID] == nil {
		r.users = make(map[string]*model.User)
		r.todos = make(map[string][]*model.Todo)
	}
	r.users[ID] = &model.User{
		ID:             ID,
		UserID:         userID,
		TotalCount:     len(todos),
		CompletedCount: 0,
	}
	r.todos[ID] = todos
	for _, todo := range todos {
		if todo.Complete == true {
			r.users[ID].CompletedCount++
		}
	}
	for k, v := range r.users {
		log.Printf("## initUser ## Users %v = %v", k, v)
	}
	for k, v := range r.todos[ID] {
		log.Printf("## initUser ## Todos %v = %v", k, v)
	}
}

// Initialize ...
func (r *Resolver) initalize(userID string) *Resolver {
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
	r.initUser(userID, todos)
	return r
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
