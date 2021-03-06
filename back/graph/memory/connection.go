package memory

import (
	"encoding/base64"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/akatebi/todos/graph/model"
)

func resolveTodoConnection(todosAll []*model.Todo, status *model.Status, after *string, first *int) (*model.TodoConnection, error) {
	var todos []*model.Todo
	if *status != model.StatusAny {
		for _, todo := range todosAll {
			if todo.Complete == true {
				todos = append(todos, todo)
			}
		}
	} else {
		todos = todosAll
	}
	from, to, err := calcRange(todos, first, after)
	log.Printf("## resolveTodoConnection ## from %d, to %d", from, to)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	todoConnection := model.TodoConnection{
		PageInfo: &model.PageInfo{
			StartCursor:     encodeCursor(from),
			EndCursor:       encodeCursor(to - 1),
			HasNextPage:     to < len(todos),
			HasPreviousPage: from < to,
		},
		Edges: edges(todos, from, to),
	}
	return &todoConnection, nil
}

func calcRange(todos []*model.Todo, first *int, after *string) (int, int, error) {
	from := 0
	if after != nil {
		b, err := base64.StdEncoding.DecodeString(*after)
		if err != nil {
			return 0, 0, err
		}
		i, err := strconv.Atoi(strings.TrimPrefix(string(b), "cursor"))
		if err != nil {
			return 0, 0, err
		}
		from = i
	}
	to := len(todos)
	if first != nil {
		to = from + *first
		if to > len(todos) {
			to = len(todos)
		}
	}
	return from, to, nil
}

func edges(todos []*model.Todo, from int, to int) []*model.TodoEdge {
	edges := make([]*model.TodoEdge, to-from)
	for i := range edges {
		todo := *todos[from+i]
		edges[i] = &model.TodoEdge{
			Cursor: *encodeCursor(from + i),
			Node:   &todo,
		}
	}
	return edges
}

func encodeCursor(i int) *string {
	cursor := string(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("cursor%d", i+1))))
	return &cursor
}
