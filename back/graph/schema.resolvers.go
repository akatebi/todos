package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/akatebi/todos/graph/generated"
	"github.com/akatebi/todos/graph/model"
	"github.com/graphql-go/relay"
)

func (r *mutationResolver) AddTodo(ctx context.Context, input model.AddTodoInput) (*model.AddTodoPayload, error) {
	todo := &model.Todo{
		Text:     input.Text,
		ID:       newID(),
		Complete: false,
	}
	ID := input.UserID
	log.Printf("#### ID %v", ID)
	if r.users[ID] == nil {
		initUser(ID, []*model.Todo{})
	}
	todos := r.todos[ID]
	user := r.users[ID]
	user.TotalCount++
	log.Printf("#### User %v", *user)
	r.todos[ID] = append(todos, todo)
	log.Printf("### Todos %v", todos)
	cursor := *encodeCursor(len(todos))
	// log.Printf("cursor %v", cursor)
	payload := &model.AddTodoPayload{
		ClientMutationID: input.ClientMutationID,
		User:             user,
		TodoEdge: &model.TodoEdge{
			Cursor: cursor,
			Node:   todo,
		},
	}
	return payload, nil
}

func (r *mutationResolver) ChangeTodoStatus(ctx context.Context, input model.ChangeTodoStatusInput) (*model.ChangeTodoStatusPayload, error) {
	ID := input.UserID
	if r.users[ID] == nil {
		return nil, fmt.Errorf("User %s not exists", ID)
	}
	todos := r.todos[ID]
	user := r.users[ID]
	var payload *model.ChangeTodoStatusPayload
	ID = relay.FromGlobalID(input.ID).ID
	for _, todo := range todos {
		if todo.ID == ID {
			// log.Printf("Status Todo %#v", todo)
			if todo.Complete != input.Complete {
				if input.Complete == true {
					user.CompletedCount++
				} else {
					user.CompletedCount--
				}
			}
			todo.Complete = input.Complete
			payload = &model.ChangeTodoStatusPayload{
				Todo:             todo,
				User:             user,
				ClientMutationID: input.ClientMutationID,
			}
			break
		}
	}
	log.Printf("Status payload %#v, %#v", payload.Todo, payload.User)
	return payload, nil
}

func (r *mutationResolver) MarkAllTodos(ctx context.Context, input model.MarkAllTodosInput) (*model.MarkAllTodosPayload, error) {
	ID := input.UserID
	if r.users[ID] == nil {
		return nil, fmt.Errorf("User %s not exists", ID)
	}
	todos := r.todos[ID]
	user := r.users[ID]
	changedTodos := []*model.Todo{}
	for _, todo := range todos {
		log.Printf("Mark Todo %#v", todo)
		if todo.Complete != input.Complete {
			if input.Complete == true {
				user.CompletedCount++
			} else {
				user.CompletedCount--
			}
			changedTodos = append(changedTodos, todo)
		}
		todo.Complete = input.Complete
	}
	payload := model.MarkAllTodosPayload{
		ChangedTodos:     changedTodos,
		User:             user,
		ClientMutationID: input.ClientMutationID,
	}
	log.Printf("Mark All payload %#v, %#v", payload.ChangedTodos, payload.User)
	return &payload, nil
}

func (r *mutationResolver) RemoveCompletedTodos(ctx context.Context, input model.RemoveCompletedTodosInput) (*model.RemoveCompletedTodosPayload, error) {
	todos := r.todos[input.UserID]
	deletedTodoIds := []string{}
	for i, todo := range todos {
		if todo.Complete == true {
			id := relay.ToGlobalID("Todo", todo.ID)
			deletedTodoIds = append(deletedTodoIds, id)
			r.todos[input.UserID] = append(todos[:i], todos[i+1:]...)
		}
	}
	user := r.users[input.UserID]
	payload := model.RemoveCompletedTodosPayload{
		DeletedTodoIds:   deletedTodoIds,
		User:             user,
		ClientMutationID: input.ClientMutationID,
	}
	return &payload, nil
}

func (r *mutationResolver) RemoveTodo(ctx context.Context, input model.RemoveTodoInput) (*model.RemoveTodoPayload, error) {
	ID := relay.FromGlobalID(input.ID).ID
	todos := r.todos[input.UserID]
	index := -1
	for i, todo := range todos {
		if todo.ID == ID {
			index = i
			r.todos[input.UserID] = append(todos[:i], todos[i+1:]...)
			if todo.Complete == true {
				r.users[input.UserID].CompletedCount--
			}
			r.users[input.UserID].TotalCount--
			break
		}
	}
	if index == -1 {
		return nil, fmt.Errorf("Todo ID %v Not Found", input.ID)
	}
	user := r.users[input.UserID]
	payload := model.RemoveTodoPayload{
		DeletedTodoID:    input.ID,
		User:             user,
		ClientMutationID: input.ClientMutationID,
	}
	return &payload, nil
}

func (r *mutationResolver) RenameTodo(ctx context.Context, input model.RenameTodoInput) (*model.RenameTodoPayload, error) {
	ID := relay.FromGlobalID(input.ID).ID
	var renamed *model.Todo
	index := -1
	for user := range r.users {
		todos := r.todos[user]
		for _, todo := range todos {
			if todo.ID == ID {
				todo.Text = input.Text
				renamed = todo
				index = 1
				break
			}
		}
	}
	if index == -1 {
		return nil, fmt.Errorf("Todo ID %v Not Found", input.ID)
	}
	payload := model.RenameTodoPayload{
		Todo:             renamed,
		ClientMutationID: input.ClientMutationID,
	}
	return &payload, nil
}

func (r *queryResolver) User(ctx context.Context, id *string) (*model.User, error) {
	return r.users[*id], nil
}

func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	log.Printf("##### id %v", id)
	obj := relay.FromGlobalID(id)
	log.Printf("Node %#v", obj)
	if obj.Type == "User" {
		return r.users[obj.ID], nil
	} else if obj.Type == "Todo" {
		for _, todos := range r.todos {
			for _, todo := range todos {
				if todo.ID == obj.ID {
					log.Printf("%#v", todo)
					return todo, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("ID %v Not Found", id)
}

func (r *todoResolver) ID(ctx context.Context, obj *model.Todo) (string, error) {
	return relay.ToGlobalID("Todo", obj.ID), nil
}

func (r *userResolver) ID(ctx context.Context, obj *model.User) (string, error) {
	return relay.ToGlobalID("User", obj.ID), nil
}

func (r *userResolver) Todos(ctx context.Context, obj *model.User, status *model.Status, after *string, first *int, before *string, last *int) (*model.TodoConnection, error) {
	return resolveTodoConnection(r.todos[obj.UserID], status, after, first)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type userResolver struct{ *Resolver }