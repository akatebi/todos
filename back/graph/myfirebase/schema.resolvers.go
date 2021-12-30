package myfirebase

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/99designs/gqlgen/graphql"
	"github.com/akatebi/todos/graph/generated"
	"github.com/akatebi/todos/graph/model"
	"github.com/graphql-go/relay"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"google.golang.org/api/iterator"
)

func (r *mutationResolver) AddUser(ctx context.Context, input model.AddUserInput) (*model.AddUserPayload, error) {
	log.Printf("##### AddUser #####")
	documentRef, writeResult, err := r.client.Collection("users").Add(ctx, input.Email)
	if err != nil {
		graphql.AddError(ctx, err)
		return nil, err
	}
	log.Printf("documentRef id %v", documentRef)
	log.Printf("writeResult id %v", writeResult)
	payload := &model.AddUserPayload{
		ID:               relay.ToGlobalID("User", documentRef.ID),
		ClientMutationID: input.ClientMutationID,
	}
	return payload, nil
}

func (r *mutationResolver) RemoveUser(ctx context.Context, input model.RemoveUserInput) (*model.RemoveUserPayload, error) {
	log.Printf("##### RemoveUser #####")
	_, err := r.client.Collection("users").Doc(input.Email).Delete(ctx)
	if err != nil {
		graphql.AddError(ctx, err)
		return nil, err
	}
	ref := r.client.Collection("users").Doc(input.Email).Collection("todos")
	DeleteCollection(ctx, r.client, ref, 100)
	payload := &model.RemoveUserPayload{
		ClientMutationID: input.ClientMutationID,
	}
	return payload, nil
}

func (r *mutationResolver) AddTodo(ctx context.Context, input model.AddTodoInput) (*model.AddTodoPayload, error) {
	log.Printf("##### AddTodo %#v #####", input)
	user_id := relay.FromGlobalID(input.UserID).ID
	todos := r.client.Collection("users").Doc(user_id).Collection("todos")
	documentRef, _, err := todos.Add(ctx, map[string]interface{}{
		"text":     input.Text,
		"complete": false,
		"created":  Created(),
	})
	if err != nil {
		graphql.AddError(ctx, err)
		return nil, err
	}
	todo := r.QueryTodo(ctx, documentRef)

	user := r.QueryUser(ctx, user_id)
	payload := &model.AddTodoPayload{
		ClientMutationID: input.ClientMutationID,
		User:             user,
		TodoEdge: &model.TodoEdge{
			Cursor: fmt.Sprintf("%v", Created()),
			Node:   todo,
		},
	}
	return payload, nil
}

func (r *mutationResolver) ChangeTodoStatus(ctx context.Context, input model.ChangeTodoStatusInput) (*model.ChangeTodoStatusPayload, error) {
	log.Printf("##### ChangeTodoStatus #####")
	if relay.FromGlobalID(input.ID) == nil {
		err := &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: "Bad id",
		}
		return nil, err
	}
	if relay.FromGlobalID(input.UserID) == nil {
		err := &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: "Bad id",
		}
		return nil, err
	}
	user_id := relay.FromGlobalID(input.UserID).ID
	id := relay.FromGlobalID(input.ID).ID
	docRef := r.client.Collection("users").Doc(user_id).Collection("todos").Doc(id)
	_, err := docRef.Update(ctx, []firestore.Update{
		{
			Path:  "complete",
			Value: input.Complete,
		},
	})
	if err != nil {
		graphql.AddError(ctx, err)
		return nil, err
	}
	todo := r.QueryTodo(ctx, docRef)

	user := r.QueryUser(ctx, user_id)
	payload := &model.ChangeTodoStatusPayload{
		ClientMutationID: input.ClientMutationID,
		User:             user,
		Todo:             todo,
	}
	return payload, nil
}

func (r *mutationResolver) ClearCompletedTodos(ctx context.Context, input model.ClearCompletedTodosInput) (*model.ClearCompletedTodosPayload, error) {
	log.Printf("##### RemoveCompletedTodos #####")
	user_id := relay.FromGlobalID(input.UserID).ID
	var deletedTodoIds []string
	iter := r.client.Collection("users").Doc(user_id).Collection("todos").Where("complete", "==", true).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		complete := doc.Data()["complete"].(bool)
		if complete {
			doc.Ref.Delete(ctx)
			deletedTodoIds = append(deletedTodoIds, relay.ToGlobalID("Todo", doc.Ref.ID))
		}
	}
	user := r.QueryUser(ctx, user_id)
	payload := &model.ClearCompletedTodosPayload{
		ClientMutationID: input.ClientMutationID,
		DeletedTodoIds:   deletedTodoIds,
		User:             user,
	}
	return payload, nil
}

func (r *mutationResolver) MarkAllTodos(ctx context.Context, input model.MarkAllTodosInput) (*model.MarkAllTodosPayload, error) {
	log.Printf("##### MarkAllTodos ##### %v", input.Complete)
	if relay.FromGlobalID(input.UserID) == nil {
		err := &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: "Bad id",
		}
		return nil, err
	}
	user_id := relay.FromGlobalID(input.UserID).ID
	changedTodos := r.QueryMarkAllTodos(ctx, user_id, input.Complete)
	user := r.QueryUser(ctx, user_id)
	payload := &model.MarkAllTodosPayload{
		ClientMutationID: input.ClientMutationID,
		User:             user,
		ChangedTodos:     changedTodos,
	}
	return payload, nil
}

func (r *mutationResolver) RemoveTodo(ctx context.Context, input model.RemoveTodoInput) (*model.RemoveTodoPayload, error) {
	log.Printf("##### RemoveTodo ##### %#v", input)
	if relay.FromGlobalID(input.ID) == nil {
		err := &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: "Bad id",
		}
		return nil, err
	}
	id := relay.FromGlobalID(input.ID).ID
	user_id := relay.FromGlobalID(input.UserID).ID
	_, err := r.client.Collection("users").Doc(user_id).Collection("todos").Doc(id).Delete(ctx)
	Panic(err)
	if err != nil {
		graphql.AddError(ctx, err)
		return nil, err
	}
	if relay.FromGlobalID(input.UserID) == nil {
		err := &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: "Bad id",
		}
		return nil, err
	}
	user := r.QueryUser(ctx, user_id)
	payload := &model.RemoveTodoPayload{
		ClientMutationID: input.ClientMutationID,
		DeletedTodoID:    input.ID,
		User:             user,
	}
	return payload, nil
}

func (r *mutationResolver) RenameTodo(ctx context.Context, input model.RenameTodoInput) (*model.RenameTodoPayload, error) {
	log.Printf("##### RenameTodo #####")
	if relay.FromGlobalID(input.ID) == nil {
		err := &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: "Bad id",
		}
		return nil, err
	}
	user_id := relay.FromGlobalID(input.UserID).ID
	id := relay.FromGlobalID(input.ID).ID
	docRef := r.client.Collection("users").Doc(user_id).Collection("todos").Doc(id)
	_, err := docRef.Update(ctx, []firestore.Update{
		{
			Path:  "text",
			Value: input.Text,
		},
	})
	if err != nil {
		graphql.AddError(ctx, err)
		return nil, err
	}
	todo := r.QueryTodo(ctx, docRef)
	log.Printf("Todo %v", todo)
	payload := &model.RenameTodoPayload{
		ClientMutationID: input.ClientMutationID,
		Todo:             todo,
	}
	return payload, nil
}

func (r *queryResolver) User(ctx context.Context, email *string) (*model.User, error) {
	log.Printf("##### User %v #####", *email)
	user := r.QueryUser(ctx, *email)
	return user, nil
}

func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	log.Printf("##### Node %v #####", id)
	obj := relay.FromGlobalID(id)
	if obj == nil {
		err := &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: "Bad id",
		}
		return nil, err
	}
	if obj.Type == "User" {
		return r.QueryUser(ctx, obj.ID), nil
	}
	return nil, fmt.Errorf("ID %v Not Found", id)
}

func (r *userResolver) Todos(ctx context.Context, obj *model.User, status *model.Status, after *string, first *int, before *string, last *int) (*model.TodoConnection, error) {
	log.Printf("##### Todos Connection ##### obj %#v", obj)
	if relay.FromGlobalID(obj.ID) == nil {
		err := &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: "Bad id",
		}
		return nil, err
	}
	user_id := relay.FromGlobalID(obj.ID).ID
	return r.resolveTodoConnection(ctx, user_id, status, after, first, before, last)
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
