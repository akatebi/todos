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

func (r *mutationResolver) AddUser(ctx context.Context, input model.AddUserInput) (*model.AddUserPayload, error) {
	log.Printf("##### AddUser #####")
	stmt, e := r.db.Prepare("INSERT INTO user(email) VALUES(?)")
	Panic(e)
	res, e := stmt.Exec(input.Email)
	Panic(e)
	id, e := res.LastInsertId()
	Panic(e)
	log.Printf("Insert User id %v", id)
	payload := &model.AddUserPayload{
		ID:               relay.ToGlobalID("User", strconv.FormatInt(id, 10)),
		ClientMutationID: input.ClientMutationID,
	}
	return payload, nil
}

func (r *mutationResolver) RemoveUser(ctx context.Context, input model.RemoveUserInput) (*model.RemoveUserPayload, error) {
	log.Printf("##### RemoveUser #####")
	id := relay.FromGlobalID(input.ID).ID
	Stmt, err := r.db.Prepare("DELETE FROM user WHERE id=?")
	res, err := Stmt.Exec(id)
	Panic(err)
	rowsAffected, err := res.RowsAffected()
	Panic(err)
	log.Printf("Rows Affected %v", rowsAffected)
	payload := &model.RemoveUserPayload{
		ClientMutationID: input.ClientMutationID,
	}
	return payload, nil
}

func (r *mutationResolver) AddTodo(ctx context.Context, input model.AddTodoInput) (*model.AddTodoPayload, error) {
	log.Printf("##### AddTodo %#v #####", input)
	stmt, e := r.db.Prepare("INSERT INTO todo(user_id, text, complete) VALUES(?,?,?)")
	Panic(e)
	user_id := relay.FromGlobalID(input.UserID).ID
	res, e := stmt.Exec(user_id, input.Text, false)
	Panic(e)
	id, e := res.LastInsertId()
	Panic(e)
	log.Printf("Insert id %v", id)
	todo := r.QueryTodo(strconv.FormatInt(id, 10))
	user := r.QueryUser(user_id)
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
	log.Printf("##### ChangeTodoStatus #####")
	obj := relay.FromGlobalID(input.ID)
	stmt, err := r.db.Prepare("update todo set Complete=? where id=?")
	Panic(err)
	res, err := stmt.Exec(input.Complete, obj.ID)
	Panic(err)
	a, err := res.RowsAffected()
	Panic(err)
	log.Printf("Rows affected %v", a)
	todo := r.QueryTodo(obj.ID)
	user_id := relay.FromGlobalID(input.UserID).ID
	user := r.QueryUser(user_id)
	payload := &model.ChangeTodoStatusPayload{
		ClientMutationID: input.ClientMutationID,
		User:             user,
		Todo:             todo,
	}
	return payload, nil
}

func (r *mutationResolver) ClearCompletedTodos(ctx context.Context, input model.ClearCompletedTodosInput) (*model.ClearCompletedTodosPayload, error) {
	log.Printf("##### RemoveCompletedTodos #####")
	id := relay.FromGlobalID(input.UserID).ID
	var deletedTodoIds []string
	rows, err := r.db.Query("SELECT id FROM todo WHERE user_id=? AND complete=true", id)
	Panic(err)
	for rows.Next() {
		var id int
		rows.Scan(&id)
		deletedTodoIds = append(deletedTodoIds, relay.ToGlobalID("Todo", strconv.Itoa(id)))
	}
	defer rows.Close()
	Stmt, err := r.db.Prepare("DELETE FROM todo WHERE user_id=? AND complete=true")
	res, err := Stmt.Exec(id)
	Panic(err)
	rowsAffected, err := res.RowsAffected()
	Panic(err)
	log.Printf("Rows Affected %v", rowsAffected)
	user_id := relay.FromGlobalID(input.UserID).ID
	user := r.QueryUser(user_id)
	payload := &model.ClearCompletedTodosPayload{
		ClientMutationID: input.ClientMutationID,
		DeletedTodoIds:   deletedTodoIds,
		User:             user,
	}
	return payload, nil
}

func (r *mutationResolver) MarkAllTodos(ctx context.Context, input model.MarkAllTodosInput) (*model.MarkAllTodosPayload, error) {
	log.Printf("##### MarkAllTodos #####")
	user_id := relay.FromGlobalID(input.UserID).ID
	changedTodos := r.QueryMarkAllTodos(user_id, input.Complete)
	user := r.QueryUser(user_id)
	payload := &model.MarkAllTodosPayload{
		ClientMutationID: input.ClientMutationID,
		User:             user,
		ChangedTodos:     changedTodos,
	}
	return payload, nil
}

func (r *mutationResolver) RemoveTodo(ctx context.Context, input model.RemoveTodoInput) (*model.RemoveTodoPayload, error) {
	log.Printf("##### RemoveTodo #####")
	id := relay.FromGlobalID(input.ID).ID
	Stmt, err := r.db.Prepare("DELETE FROM todo WHERE id=?")
	res, err := Stmt.Exec(id)
	Panic(err)
	rowsAffected, err := res.RowsAffected()
	Panic(err)
	log.Printf("Rows Affected %v", rowsAffected)
	user_id := relay.FromGlobalID(input.UserID).ID
	user := r.QueryUser(user_id)
	payload := &model.RemoveTodoPayload{
		ClientMutationID: input.ClientMutationID,
		DeletedTodoID:    input.ID,
		User:             user,
	}
	return payload, nil
}

func (r *mutationResolver) RenameTodo(ctx context.Context, input model.RenameTodoInput) (*model.RenameTodoPayload, error) {
	ID := relay.FromGlobalID(input.ID).ID
	log.Printf("##### RenameTodo #####")
	Stmt, err := r.db.Prepare("UPDATE todo SET text=? WHERE id=?")
	Panic(err)
	res, err := Stmt.Exec(input.Text, ID)
	Panic(err)
	rowsAffected, err := res.RowsAffected()
	Panic(err)
	log.Printf("Rows Affected %v", rowsAffected)
	todo := r.QueryTodo(ID)
	log.Printf("Todo %v", todo)
	payload := &model.RenameTodoPayload{
		ClientMutationID: input.ClientMutationID,
		Todo:             todo,
	}
	return payload, nil
}

func (r *queryResolver) User(ctx context.Context, email *string) (*model.User, error) {
	log.Printf("##### User %v #####", *email)
	var user_id int
	user := &model.User{
		ID:             "",
		Email:          *email,
		Todos:          &model.TodoConnection{},
		TotalCount:     0,
		CompletedCount: 0,
	}
	r.db.QueryRow("SELECT id FROM user WHERE email=? LIMIT 1", *email).Scan(&user_id)
	r.db.QueryRow("SELECT COUNT(*) FROM todo WHERE user_id=?", user_id).Scan(&user.TotalCount)
	r.db.QueryRow("SELECT COUNT(*) FROM todo WHERE user_id=? AND Complete=true", user_id).Scan(&user.CompletedCount)
	user.ID = relay.ToGlobalID("User", strconv.Itoa(user_id))
	log.Printf("user %v", user)
	return user, nil
}

func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	log.Printf("##### Node %v #####", id)
	obj := relay.FromGlobalID(id)
	if obj.Type == "User" {
		return r.QueryUser(obj.ID), nil
	} else if obj.Type == "Todo" {
		return r.QueryTodo(obj.ID), nil
	}
	return nil, fmt.Errorf("ID %v Not Found", id)
}

func (r *userResolver) Todos(ctx context.Context, obj *model.User, status *model.Status, after *string, first *int, before *string, last *int) (*model.TodoConnection, error) {
	log.Printf("##### Todos #####")
	user_id := relay.FromGlobalID(obj.ID).ID
	after_ := DecodeCursor(after)
	before_ := DecodeCursor(before)
	var first_, last_ int
	if first != nil {
		first_ = *first
	}
	if last != nil {
		last_ = *last
	}
	return r.resolveTodoConnection(user_id, status, after_, first_, before_, last_)
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
