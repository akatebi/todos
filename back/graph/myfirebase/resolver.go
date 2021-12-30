package myfirebase

//go:generate go run github.com/99designs/gqlgen

//  k exec -it mysql-69567fc988-r54x2 -c mysql -- bash

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"

	"github.com/akatebi/todos/graph/model"
	"github.com/graphql-go/relay"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client *firestore.Client
}

func (r *Resolver) QueryUser(ctx context.Context, UserID string) *model.User {
	log.Printf("### UserID %v", UserID)
	ID := relay.ToGlobalID("User", UserID)
	user := &model.User{
		ID:             ID,
		Email:          UserID,
		Todos:          &model.TodoConnection{},
		TotalCount:     0,
		CompletedCount: 0,
	}
	iter := r.client.Collection("users").Doc(UserID).Collection("todos").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		user.TotalCount += 1
		if doc.Data()["complete"] == true {
			user.CompletedCount += 1
		}
		log.Printf("## data: %v", doc.Data())
	}
	return user
}

func (r *Resolver) QueryTodo(ctx context.Context, ref *firestore.DocumentRef) *model.Todo {
	snapshot, err := ref.Get(ctx)
	log.Printf("### Err %v", err)
	doc := snapshot.Data()
	todo := &model.Todo{
		Text:     doc["text"].(string),
		Complete: doc["complete"].(bool),
		ID:       relay.ToGlobalID("Todo", ref.ID),
	}
	return todo
}

func (r *Resolver) QueryMarkAllTodos(ctx context.Context, user_id string, complete bool) []*model.Todo {
	log.Printf("QueryMarkAllTodos complete %v", complete)
	colRef := r.client.Collection("users").Doc(user_id).Collection("todos")
	iter := colRef.DocumentRefs(ctx)
	todos := []*model.Todo{}
	for {
		docRef, err := iter.Next()
		if err == iterator.Done {
			break
		}
		docRef.Update(ctx, []firestore.Update{
			{
				Path:  "complete",
				Value: complete,
			},
		})
		snapshut, _ := docRef.Get(ctx)
		text := snapshut.Data()["text"].(string)
		complete := snapshut.Data()["complete"].(bool)
		todo := &model.Todo{
			ID:       relay.ToGlobalID("Todo", docRef.ID),
			Complete: complete,
			Text:     text,
		}
		todos = append(todos, todo)
	}
	return todos
}

func (r *Resolver) Close() {
	r.client.Close()
}

func (r *Resolver) Open() {

	ctx := context.Background()
	sa := option.WithCredentialsFile("./serviceAccount.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	r.client = client

	r.initialize(ctx)
}

func DeleteCollection(ctx context.Context, client *firestore.Client,
	ref *firestore.CollectionRef, batchSize int) error {

	for {
		// Get a batch of documents
		iter := ref.Limit(batchSize).Documents(ctx)
		numDeleted := 0

		// Iterate through the documents, adding
		// a delete operation for each one to a
		// WriteBatch.
		batch := client.Batch()
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}

			batch.Delete(doc.Ref)
			numDeleted++
		}

		// If there are no documents to delete,
		// the process is over.
		if numDeleted == 0 {
			return nil
		}

		_, err := batch.Commit(ctx)
		if err != nil {
			return err
		}
	}
}

func Created() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func (r *Resolver) initialize(ctx context.Context) {
	emails := []string{"me@gmail.com"}
	for _, email := range emails {
		// DeleteCollection(ctx, r.client, r.client.Collection("users").Doc(email).Collection("todos"), 100)
		colRef := r.client.Collection("users").Doc(email).Collection("todos")
		iter := colRef.Documents(ctx)
		_, err := iter.Next()
		if err == iterator.Done {
			_, _, err := colRef.Add(ctx, map[string]interface{}{
				"text":     "Taste JavaScript",
				"complete": true,
				"created":  Created(),
			})
			if err != nil {
				// Handle any errors in an appropriate way, such as returning them.
				log.Printf("An error has occurred: %s", err)
			}
			_, _, err = colRef.Add(ctx, map[string]interface{}{
				"text":     "Buy a unicorn",
				"complete": false,
				"created":  Created(),
			})
			if err != nil {
				// Handle any errors in an appropriate way, such as returning them.
				log.Printf("An error has occurred: %s", err)
			}
			_, _, err = colRef.Add(ctx, map[string]interface{}{
				"text":     "Get a customer",
				"complete": false,
				"created":  Created(),
			})
			if err != nil {
				// Handle any errors in an appropriate way, such as returning them.
				log.Printf("An error has occurred: %s", err)
			}
		}
	}
}

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}
