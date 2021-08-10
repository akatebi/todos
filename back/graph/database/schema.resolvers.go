package database

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
	panic(fmt.Errorf("not implemented"))
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
	rows, err := r.db.Query("Select * FROM Users WHERE UserID=? LIMIT 1", *id)
	Panic(err)
	log.Printf("User %v", rows)
	var UserID string
	var ID, TotalCount, CompletedCount int
	for rows.Next() {
		err = rows.Scan(&ID, &UserID, &CompletedCount, &TotalCount)
		Panic(err)
		fmt.Println(ID, UserID, CompletedCount, TotalCount)
	}
	Panic(rows.Err())
	rows.Close()
	return &model.User{
		ID:             relay.ToGlobalID("User", ToString(ID)),
		UserID:         UserID,
		TotalCount:     TotalCount,
		CompletedCount: CompletedCount,
	}, nil
}

func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	log.Printf("Node")
	node := relay.FromGlobalID(id)
	if node.Type == "User" {
		rows, err := r.db.Query("SELECT ID, UserID FROM Users WHERE id=? LIMIT 1", node.ID)
		user := &model.User{}
		for rows.Next() {
			rows.Scan(&user.ID, &user.UserID)
		}
		rows.Close()
		Panic(err)
		r.db.QueryRow("SELECT COUNT(*) FROM Todos WHERE Users_id=?", &user.ID).Scan(&user.TotalCount)
		r.db.QueryRow("SELECT COUNT(*) FROM Todos WHERE Users_id=? AND Complete=true", &user.ID).Scan(&user.CompletedCount)
		return user, nil
	} else if node.Type == "Todo" {
		todo := &model.Todo{}
		return todo, nil
	}
	return nil, fmt.Errorf("ID %v Not Found", id)
}

func (r *userResolver) Todos(ctx context.Context, obj *model.User, status *model.Status, after *string, first *int, before *string, last *int) (*model.TodoConnection, error) {
	log.Printf("Todos")
	Users_id := relay.FromGlobalID(obj.ID).ID
	return r.resolveTodoConnection(Users_id, status, after, first, before, last)
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
