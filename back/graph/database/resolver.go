package database

//  k exec -it mysql-69567fc988-r54x2 -c mysql -- bash
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/akatebi/todos/graph/model"
	_ "github.com/go-sql-driver/mysql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	db *sql.DB
}

func (r *Resolver) QueryUser(id string) (*model.User, error) {
	rows, err := r.db.Query("SELECT ID, UserID FROM Users WHERE id=? LIMIT 1", id)
	user := &model.User{}
	for rows.Next() {
		rows.Scan(&user.ID, &user.Email)
	}
	rows.Close()
	Panic(err)
	r.db.QueryRow("SELECT COUNT(*) FROM Todos WHERE Users_id=?", &user.ID).Scan(&user.TotalCount)
	r.db.QueryRow("SELECT COUNT(*) FROM Todos WHERE Users_id=? AND Complete=true", &user.ID).Scan(&user.CompletedCount)
	return user, nil
}

func (r *Resolver) QueryTodo(id string) (*model.Todo, error) {
	rows, err := r.db.Query("SELECT ID, Text, Complete FROM Todos WHERE id=? LIMIT 1", id)
	todo := &model.Todo{}
	for rows.Next() {
		rows.Scan(&todo.ID, &todo.Text, &todo.Complete)
	}
	rows.Close()
	Panic(err)
	return todo, nil
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
		"DROP DATABASE IF EXISTS Todos",
		"CREATE DATABASE IF NOT EXISTS Todos",
		"USE Todos",
		`CREATE TABLE Users (
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			Email VARCHAR(32),
			TotalCount INT DEFAULT 0,
			CompletedCount INT DEFAULT 0
		) ENGINE=INNODB`,
		`CREATE TABLE Todos (
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			Users_id INT,
			Text VARCHAR(32),
			Complete BOOLEAN DEFAULT false,
			FOREIGN KEY (Users_id)
				REFERENCES Users(id)
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

	Email := "me@gmail.com"
	stmt, e := db.Prepare("INSERT INTO Users(Email) VALUES(?)")
	Panic(e)
	res, e := stmt.Exec(Email)
	Panic(e)
	Users_id, e := res.LastInsertId()
	Panic(e)
	log.Printf("Insert Users_id %v", Users_id)

	stmt, e = db.Prepare("INSERT INTO Todos(Users_id, Text, Complete) VALUES(?,?,?)")
	Panic(e)

	res, e = stmt.Exec(Users_id, "Taste JavaScript", true)
	Panic(e)
	id, e := res.LastInsertId()
	Panic(e)
	log.Printf("Insert id %v", id)

	res, e = stmt.Exec(Users_id, "Buy a unicorn", false)
	Panic(e)
	id, e = res.LastInsertId()
	Panic(e)
	log.Printf("Insert id %v", id)

	res, e = stmt.Exec(Users_id, "Get a customer", false)
	Panic(e)
	id, e = res.LastInsertId()
	Panic(e)
	log.Printf("Insert id %v", id)

	r.db = db
}

func Panic(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func ToString(id int) string {
	json, err := json.Marshal(id)
	if err != nil {
		panic(err)
	}
	return string(json)
}
