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
	id_Users string,
	status *model.Status,
	after *string, first *int,
	before *string, last *int) (*model.TodoConnection, error) {

	afterInt := DecodeCursor(after)
	log.Printf("id_User %v status %v, after %v, first %v", id_Users, status, afterInt, *first)

	id_User, err := strconv.Atoi(id_Users)
	Panic(err)
	var rows *sql.Rows
	if *status == model.StatusAny {
		rows, err = r.db.Query("Select * FROM Todos WHERE id_User = ? AND id > ? LIMIT ?", id_User, afterInt, *first)
	} else {
		rows, err = r.db.Query("Select * FROM Todos WHERE id_User = ? AND id > ? AND Complete = ? LIMIT ?", id_User, afterInt, *status == model.StatusCompleted, *first)
	}
	Panic(err)
	log.Printf("Todos %v", rows)

	var edges []*model.TodoEdge
	count := 0
	var StartCursor, EndCursor *string
	for rows.Next() {
		var ID, id_User int
		var Text string
		var Complete bool
		err = rows.Scan(&ID, &id_User, &Text, &Complete)
		Panic(err)
		if count == 0 {
			StartCursor = EncodeCursor(ID)
		} else {
			EndCursor = EncodeCursor(ID)
		}
		count++
		log.Printf("#### ID %v %v %v %v", ID, id_User, Text, Complete)
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
