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
	Users_ids string,
	status *model.Status,
	after *string, first *int,
	before *string, last *int) (*model.TodoConnection, error) {

	log.Printf("Users_id %v status %v, after %v, first %v", Users_ids, status, after, *first)

	Users_id, err := strconv.Atoi(Users_ids)
	Panic(err)
	var rows *sql.Rows
	if *status == model.StatusAny {
		rows, err = r.db.Query("Select * FROM Todos WHERE Users_id = ? AND id > ? LIMIT ?", Users_id, decodeCursor(after), *first)
	} else {
		rows, err = r.db.Query("Select * FROM Todos WHERE Users_id = ? AND id > ? AND Complete = ? LIMIT ?", Users_id, decodeCursor(after), *status == model.StatusCompleted, *first)
	}
	Panic(err)
	log.Printf("Todos %v", rows)

	var edges []*model.TodoEdge
	count := 0
	var StartCursor, EndCursor *string
	for rows.Next() {
		var ID, Users_id int
		var Text string
		var Complete bool
		err = rows.Scan(&ID, &Users_id, &Text, &Complete)
		Panic(err)
		if count == 0 {
			StartCursor = encodeCursor(ID)
		} else {
			EndCursor = encodeCursor(ID)
		}
		count++
		log.Printf("#### ID %v %v %v %v", ID, Users_id, Text, Complete)
		Node := &model.Todo{
			ID:       relay.ToGlobalID("Todo", strconv.Itoa(ID)),
			Text:     Text,
			Complete: Complete,
		}
		log.Printf("#### Node %v", Node)
		edge := &model.TodoEdge{
			Cursor: *encodeCursor(ID),
			Node:   Node,
		}
		edges = append(edges, edge)
	}
	Panic(rows.Err())
	rows.Close()
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

func encodeCursor(id int) *string {
	cursor := string(base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(id))))
	return &cursor
}

func decodeCursor(cursor *string) int {
	if cursor == nil {
		return 0
	}
	sDec, _ := base64.StdEncoding.DecodeString(*cursor)
	id, _ := strconv.Atoi(string(sDec))
	return id
}
