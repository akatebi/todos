package database

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

func (r *Resolver) QueryUser(id string) (*model.User, error) {
	var ID int
	user := &model.User{}
	r.db.QueryRow("SELECT ID, Email FROM Users WHERE id=? LIMIT 1", id).Scan(&ID, &user.Email)
	r.db.QueryRow("SELECT COUNT(*) FROM Todos WHERE UserId=?", &user.ID).Scan(&user.TotalCount)
	r.db.QueryRow("SELECT COUNT(*) FROM Todos WHERE UserId=? AND Complete=true", &user.ID).Scan(&user.CompletedCount)
	user.ID = relay.ToGlobalID("User", strconv.Itoa(ID))
	return user, nil
}

func (r *Resolver) QueryTodo(id string) (*model.Todo, error) {
	var ID int
	todo := &model.Todo{}
	r.db.QueryRow("SELECT ID, Text, Complete FROM Todos WHERE id=? LIMIT 1", id).Scan(&ID, &todo.Text, &todo.Complete)
	todo.ID = relay.ToGlobalID("Todo", strconv.Itoa(ID))
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
			UserId INT,
			Text VARCHAR(32),
			Complete BOOLEAN DEFAULT false,
			FOREIGN KEY (UserId)
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
	UserId, e := res.LastInsertId()
	Panic(e)
	log.Printf("Insert UserId %v", UserId)

	stmt, e = db.Prepare("INSERT INTO Todos(UserId, Text, Complete) VALUES(?,?,?)")
	Panic(e)

	res, e = stmt.Exec(UserId, "Taste JavaScript", true)
	Panic(e)
	id, e := res.LastInsertId()
	Panic(e)
	log.Printf("Insert id %v", id)

	res, e = stmt.Exec(UserId, "Buy a unicorn", false)
	Panic(e)
	id, e = res.LastInsertId()
	Panic(e)
	log.Printf("Insert id %v", id)

	res, e = stmt.Exec(UserId, "Get a customer", false)
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
