// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Node interface {
	IsNode()
}

type AddTodoInput struct {
	Text             string  `json:"text"`
	UserID           string  `json:"userId"`
	ClientMutationID *string `json:"clientMutationId"`
}

type AddTodoPayload struct {
	TodoEdge         *TodoEdge `json:"todoEdge"`
	User             *User     `json:"user"`
	ClientMutationID *string   `json:"clientMutationId"`
}

type AddUserInput struct {
	Email            string  `json:"email"`
	ClientMutationID *string `json:"clientMutationId"`
}

type AddUserPayload struct {
	ID               string  `json:"id"`
	ClientMutationID *string `json:"clientMutationId"`
}

type ChangeTodoStatusInput struct {
	Complete         bool    `json:"complete"`
	ID               string  `json:"id"`
	UserID           string  `json:"userId"`
	ClientMutationID *string `json:"clientMutationId"`
}

type ChangeTodoStatusPayload struct {
	Todo             *Todo   `json:"todo"`
	User             *User   `json:"user"`
	ClientMutationID *string `json:"clientMutationId"`
}

type ClearCompletedTodosInput struct {
	UserID           string  `json:"userId"`
	ClientMutationID *string `json:"clientMutationId"`
}

type ClearCompletedTodosPayload struct {
	DeletedTodoIds   []string `json:"deletedTodoIds"`
	User             *User    `json:"user"`
	ClientMutationID *string  `json:"clientMutationId"`
}

type MarkAllTodosInput struct {
	Complete         bool    `json:"complete"`
	UserID           string  `json:"userId"`
	ClientMutationID *string `json:"clientMutationId"`
}

type MarkAllTodosPayload struct {
	ChangedTodos     []*Todo `json:"changedTodos"`
	User             *User   `json:"user"`
	ClientMutationID *string `json:"clientMutationId"`
}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	EndCursor       *string `json:"endCursor"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor"`
}

type RemoveTodoInput struct {
	ID               string  `json:"id"`
	UserID           string  `json:"userId"`
	ClientMutationID *string `json:"clientMutationId"`
}

type RemoveTodoPayload struct {
	DeletedTodoID    string  `json:"deletedTodoId"`
	User             *User   `json:"user"`
	ClientMutationID *string `json:"clientMutationId"`
}

type RemoveUserInput struct {
	Email            string  `json:"email"`
	ClientMutationID *string `json:"clientMutationId"`
}

type RemoveUserPayload struct {
	ClientMutationID *string `json:"clientMutationId"`
}

type RenameTodoInput struct {
	ID               string  `json:"id"`
	Text             string  `json:"text"`
	ClientMutationID *string `json:"clientMutationId"`
}

type RenameTodoPayload struct {
	Todo             *Todo   `json:"todo"`
	ClientMutationID *string `json:"clientMutationId"`
}

type Todo struct {
	ID       string `json:"id"`
	Text     string `json:"text"`
	Complete bool   `json:"complete"`
}

func (Todo) IsNode() {}

type TodoConnection struct {
	PageInfo *PageInfo   `json:"pageInfo"`
	Edges    []*TodoEdge `json:"edges"`
}

type TodoEdge struct {
	Node   *Todo  `json:"node"`
	Cursor string `json:"cursor"`
}

type User struct {
	ID             string          `json:"id"`
	Email          string          `json:"email"`
	Todos          *TodoConnection `json:"todos"`
	TotalCount     int             `json:"totalCount"`
	CompletedCount int             `json:"completedCount"`
}

func (User) IsNode() {}

type Status string

const (
	StatusAny       Status = "ANY"
	StatusCompleted Status = "COMPLETED"
)

var AllStatus = []Status{
	StatusAny,
	StatusCompleted,
}

func (e Status) IsValid() bool {
	switch e {
	case StatusAny, StatusCompleted:
		return true
	}
	return false
}

func (e Status) String() string {
	return string(e)
}

func (e *Status) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Status(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}

func (e Status) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
