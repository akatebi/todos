package myfirebase

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"strconv"

	"cloud.google.com/go/firestore"
	"github.com/akatebi/todos/graph/model"
	"github.com/graphql-go/relay"
	"google.golang.org/api/iterator"
)

func (r *userResolver) getQueryRef(ctx context.Context,
	user_id string,
	status *model.Status,
	after *string, first *int,
	before *string, last *int) *firestore.Query {
	var queryRef firestore.Query
	colRef := r.client.Collection("users").Doc(user_id).Collection("todos")

	if after != nil {
		after_, _ := strconv.ParseInt(*after, 10, 64)
		var first_ int = 9999999
		if first != nil {
			first_ = *first
		}
		log.Printf("##### user_id %v status %v, after %v, first %v", user_id, status, after_, first_)
		if *status == model.StatusCompleted {
			queryRef = colRef.Where("complete", "==", true).OrderBy("created", firestore.Asc).StartAfter(after_).Limit(first_)
		} else {
			queryRef = colRef.OrderBy("created", firestore.Asc).StartAfter(after_).Limit(first_)
		}
	} else if before != nil {
		before_, _ := strconv.ParseInt(*before, 10, 64)
		var last_ int = 9999999
		if last != nil {
			last_ = *last
		}
		log.Printf("##### user_id %v status %v, before %v, last %v", user_id, status, before_, last_)
		if *status == model.StatusCompleted {
			queryRef = colRef.Where("complete", "==", true).OrderBy("created", firestore.Asc).EndBefore(before_).Limit(last_)
		} else {
			queryRef = colRef.OrderBy("created", firestore.Desc).StartAfter(before_).Limit(last_)
		}
	} else {
		log.Printf("##### user_id %v status %v", user_id, status)
		if *status == model.StatusCompleted {
			queryRef = colRef.Where("complete", "==", true).OrderBy("created", firestore.Asc)
		} else {
			queryRef = colRef.OrderBy("created", firestore.Asc)
		}
	}
	return &queryRef
}

func (r *userResolver) resolveTodoConnection(
	ctx context.Context,
	user_id string,
	status *model.Status,
	after *string, first *int,
	before *string, last *int) (*model.TodoConnection, error) {

	queryRef := r.getQueryRef(ctx, user_id, status, after, first, before, last)

	iter := queryRef.Documents(ctx)

	var edges []*model.TodoEdge
	count := 0
	var StartCursor, EndCursor *string
	for {
		snapshot, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("##### ERROR %v", err)
		}
		doc := snapshot.Data()
		log.Printf("#### Doc: %#v", doc)
		ID := snapshot.Ref.ID
		Text, _ := doc["text"].(string)
		Complete, _ := doc["complete"].(bool)
		Created, _ := doc["created"].(int64)
		Createds := fmt.Sprintf("%v", Created)
		if count == 0 {
			StartCursor = &Createds
			EndCursor = &Createds
		} else {
			EndCursor = &Createds
		}
		count++
		log.Printf("#### ID %v %v %v %v", ID, user_id, Text, Complete)
		Node := &model.Todo{
			ID:       relay.ToGlobalID("Todo", ID),
			Text:     Text,
			Complete: Complete,
		}
		log.Printf("#### Node %v", Node)
		edge := &model.TodoEdge{
			Cursor: Createds,
			Node:   Node,
		}
		edges = append(edges, edge)
	}

	if after != nil {
		queryRef = r.getQueryRef(ctx, user_id, status, EndCursor, first, before, last)
	} else if before != nil {
		queryRef = r.getQueryRef(ctx, user_id, status, after, first, EndCursor, last)
	} else {
		queryRef = r.getQueryRef(ctx, user_id, status, EndCursor, first, before, last)
	}

	iter = queryRef.Documents(ctx)
	_, err := iter.Next()
	var HasPage bool = true
	if err == iterator.Done {
		HasPage = false
	}

	pageInfo := &model.PageInfo{
		StartCursor:     StartCursor,
		EndCursor:       EndCursor,
		HasNextPage:     HasPage,
		HasPreviousPage: HasPage,
	}
	todoConnection := &model.TodoConnection{
		PageInfo: pageInfo,
		Edges:    edges,
	}
	return todoConnection, nil
}

func EncodeCursor(id string) *string {
	cursor := string(base64.StdEncoding.EncodeToString([]byte(id)))
	return &cursor
}

func DecodeCursor(cursor *string) string {
	if cursor == nil {
		return ""
	}
	sDec, _ := base64.StdEncoding.DecodeString(*cursor)
	id := string(sDec)
	return id
}
