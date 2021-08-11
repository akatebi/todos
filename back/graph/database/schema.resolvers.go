package database

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/akatebi/todos/graph/generated"
	"github.com/akatebi/todos/graph/model"
	"github.com/graphql-go/relay"
)

func (r *mutationResolver) AddTodo(ctx context.Context, input model.AddTodoInput) (*model.AddTodoPayload, error) {
	log.Printf("AddTodo %v", input)
	stmt, e := r.db.Prepare("INSERT INTO Todos(UserId, Text, Complete) VALUES(?,?,?)")
	Panic(e)
	UserId := relay.FromGlobalID(input.UserID).ID
	res, e := stmt.Exec(UserId, input.Text, false)
	Panic(e)
	id, e := res.LastInsertId()
	Panic(e)
	log.Printf("Insert id %v", id)
	user, _ := r.QueryUser(input.UserID)
	todo, _ := r.QueryTodo(strconv.Itoa(int(id)))
	payload := &model.AddTodoPayload{
		ClientMutationID: input.ClientMutationID,
		User:             user,
		TodoEdge: &model.TodoEdge{
			Cursor: *EncodeCursor(int(id)),
			Node:   todo,
		},
	}
	return payload, nil
}

func (r *mutationResolver) ChangeTodoStatus(ctx context.Context, input model.ChangeTodoStatusInput) (*model.ChangeTodoStatusPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) MarkAllTodos(ctx context.Context, input model.MarkAllTodosInput) (*model.MarkAllTodosPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveCompletedTodos(ctx context.Context, input model.RemoveCompletedTodosInput) (*model.RemoveCompletedTodosPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveTodo(ctx context.Context, input model.RemoveTodoInput) (*model.RemoveTodoPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RenameTodo(ctx context.Context, input model.RenameTodoInput) (*model.RenameTodoPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id *string) (*model.User, error) {
	log.Printf("### User %v", *id)
	rows, err := r.db.Query("SELECT ID, Email FROM Users WHERE email=? LIMIT 1", *id)
	user := &model.User{}
	var ID int
	for rows.Next() {
		rows.Scan(&ID, &user.Email)
	}
	rows.Close()
	Panic(err)
	r.db.QueryRow("SELECT COUNT(*) FROM Todos WHERE UserId=?", ID).Scan(&user.TotalCount)
	r.db.QueryRow("SELECT COUNT(*) FROM Todos WHERE UserId=? AND Complete=true", ID).Scan(&user.CompletedCount)
	user.ID = relay.ToGlobalID("User", strconv.Itoa(ID))
	return user, nil
}

func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	log.Printf("Node %v", id)
	obj := relay.FromGlobalID(id)
	if obj.Type == "User" {
		return r.QueryUser(obj.ID)
	} else if obj.Type == "Todo" {
		return r.QueryTodo(obj.ID)
	}
	return nil, fmt.Errorf("ID %v Not Found", id)
}

func (r *userResolver) Todos(ctx context.Context, obj *model.User, status *model.Status, after *string, first *int, before *string, last *int) (*model.TodoConnection, error) {
	log.Printf("Todos")
	UserId := relay.FromGlobalID(obj.ID).ID
	return r.resolveTodoConnection(UserId, status, after, first, before, last)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
