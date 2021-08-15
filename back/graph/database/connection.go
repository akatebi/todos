package database

import (
	"database/sql"
	"encoding/base64"
	"log"
	"strconv"

	"github.com/akatebi/todos/graph/model"
	"github.com/graphql-go/relay"
)

func (r *userResolver) resolveTodoConnection(
	user_ids string,
	status *model.Status,
	after int, first int,
	before int, last int) (*model.TodoConnection, error) {

	log.Printf("user_id %v status %v, after %v, first %v", user_ids, status, after, first)

	user_id, err := strconv.Atoi(user_ids)
	Panic(err)
	var rows *sql.Rows
	if *status == model.StatusAny {
		rows, err = r.db.Query("Select * FROM todo WHERE user_id = ? AND id > ? LIMIT ?", user_id, after, first)
	} else {
		rows, err = r.db.Query("Select * FROM todo WHERE user_id = ? AND id > ? AND complete = ? LIMIT ?", user_id, after, *status == model.StatusCompleted, first)
	}
	// Panic(err)
	log.Printf("Todos %v", err)

	var edges []*model.TodoEdge
	count := 0
	var StartCursor, EndCursor *string
	for rows.Next() {
		var ID, user_id int
		var Text string
		var Complete bool
		err = rows.Scan(&ID, &user_id, &Text, &Complete)
		Panic(err)
		if count == 0 {
			StartCursor = EncodeCursor(ID)
		} else {
			EndCursor = EncodeCursor(ID)
		}
		count++
		log.Printf("#### ID %v %v %v %v", ID, user_id, Text, Complete)
		Node := &model.Todo{
			ID:       relay.ToGlobalID("Todo", strconv.Itoa(ID)),
			Text:     Text,
			Complete: Complete,
		}
		log.Printf("#### Node %v", Node)
		edge := &model.TodoEdge{
			Cursor: *EncodeCursor(ID),
			Node:   Node,
		}
		edges = append(edges, edge)
	}
	Panic(rows.Err())
	defer rows.Close()
	pageInfo := &model.PageInfo{
		StartCursor:     StartCursor,
		EndCursor:       EndCursor,
		HasNextPage:     true,
		HasPreviousPage: true,
	}
	todoConnection := &model.TodoConnection{
		PageInfo: pageInfo,
		Edges:    edges,
	}
	return todoConnection, nil
}

func EncodeCursor(id int) *string {
	cursor := string(base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(id))))
	return &cursor
}

func DecodeCursor(cursor *string) int {
	if cursor == nil {
		return 0
	}
	sDec, _ := base64.StdEncoding.DecodeString(*cursor)
	id, _ := strconv.Atoi(string(sDec))
	return id
}
