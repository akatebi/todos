package database

//go:generate go run github.com/99designs/gqlgen

//  k exec -it mysql-69567fc988-r54x2 -c mysql -- bash

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/akatebi/todos/graph/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/graphql-go/relay"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	db *sql.DB
}

func (r *Resolver) QueryUser(id string) *model.User {
	user := &model.User{}
	r.db.QueryRow("SELECT email FROM user WHERE id=? LIMIT 1", id).Scan(&user.Email)
	r.db.QueryRow("SELECT COUNT(*) FROM todo WHERE user_id=?", id).Scan(&user.TotalCount)
	r.db.QueryRow("SELECT COUNT(*) FROM todo WHERE user_id=? AND complete=true", id).Scan(&user.CompletedCount)
	user.ID = relay.ToGlobalID("User", id)
	return user
}

func (r *Resolver) QueryTodo(id string) *model.Todo {
	todo := &model.Todo{}
	r.db.QueryRow("SELECT Text, Complete FROM todo WHERE id=? LIMIT 1", id).Scan(&todo.Text, &todo.Complete)
	todo.ID = relay.ToGlobalID("Todo", id)
	return todo
}

func (r *Resolver) QueryMarkAllTodos(user_id string, complete bool) []*model.Todo {
	log.Printf("QueryMarkAllTodos complete %v", complete)
	rows, err := r.db.Query("SELECT id, text FROM todo WHERE complete=? AND user_id=?", !complete, user_id)
	Panic(err)
	todos := []*model.Todo{}
	for rows.Next() {
		todo := &model.Todo{Complete: complete}
		var id int
		err = rows.Scan(&id, &todo.Text)
		Panic(err)
		todo.ID = relay.ToGlobalID("Todo", strconv.Itoa(id))
		todos = append(todos, todo)
	}
	rows.Close()
	Stmt, err := r.db.Prepare("UPDATE todo SET complete=? WHERE user_id=?")
	Panic(err)
	res, err := Stmt.Exec(complete, user_id)
	Panic(err)
	rowsAffected, err := res.RowsAffected()
	Panic(err)
	log.Printf("Rows Affected %v", rowsAffected)
	return todos
}

func (r *Resolver) Close() {
	r.db.Close()
}

func (r *Resolver) Open() {

	db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/")
	if err != nil {
		panic(err)
	}
	log.Printf("### db ###, %v", db)

	err = db.Ping()
	Panic(err)

	fmt.Println("### Connected! ###")

	cmds := []string{
		"DROP DATABASE IF EXISTS todo",
		"CREATE DATABASE IF NOT EXISTS todo",
		"USE todo",
		`CREATE TABLE user (
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			email VARCHAR(32) NOT NULL UNIQUE,
			totalCount INT DEFAULT 0,
			completedCount INT DEFAULT 0
		) ENGINE=INNODB`,
		`CREATE TABLE todo (
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			user_id INT,
			text VARCHAR(32),
			complete BOOLEAN DEFAULT false,
			FOREIGN KEY (user_id)
				REFERENCES user(id)
				ON DELETE CASCADE
		) ENGINE=INNODB`,
	}

	for i, cmd := range cmds {
		_, err = db.Exec(cmd)
		if err != nil {
			log.Printf("%v : %v", i, cmd)
			panic(err)
		}
	}

	emails := []string{"me@gmail.com"}
	for i, email := range emails {
		stmt, e := db.Prepare("INSERT INTO user(email) VALUES(?)")
		Panic(e)
		res, e := stmt.Exec(email)
		Panic(e)
		user_id, e := res.LastInsertId()
		Panic(e)
		log.Printf("Insert user_id %v", user_id)
		if i == 0 {
			stmt, e = db.Prepare("INSERT INTO todo(user_id, text, complete) VALUES(?,?,?)")
			Panic(e)

			res, e = stmt.Exec(user_id, "Taste JavaScript", true)
			Panic(e)
			id, e := res.LastInsertId()
			Panic(e)
			log.Printf("Insert id %v", id)

			res, e = stmt.Exec(user_id, "Buy a unicorn", false)
			Panic(e)
			id, e = res.LastInsertId()
			Panic(e)
			log.Printf("Insert id %v", id)

			res, e = stmt.Exec(user_id, "Get a customer", false)
			Panic(e)
			id, e = res.LastInsertId()
			Panic(e)
			log.Printf("Insert id %v", id)
		}
	}

	r.db = db
}

func Panic(err error) {
	if err != nil {
		panic(err.Error())
	}
}
