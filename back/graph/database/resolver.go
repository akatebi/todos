package database

//  k exec -it mysql-69567fc988-r54x2 -c mysql -- bash
import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	db *sql.DB
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
	ErrorCheck(err)

	fmt.Println("### Connected! ###")

	cmds := []string{
		"DROP DATABASE IF EXISTS Todos",
		"CREATE DATABASE IF NOT EXISTS Todos",
		"USE Todos",
		`CREATE TABLE Users (
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			UserID VARCHAR(32),
			TotalCount INT UNSIGNED DEFAULT 0,
			CompletedCount INT UNSIGNED DEFAULT 0
		) ENGINE=INNODB`,
		`CREATE TABLE Todos (
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			Users_id INT,
			Text VARCHAR(32),
			Completed BOOLEAN DEFAULT false,
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

	UserID := "me@gmail.com"
	stmt, e := db.Prepare("INSERT INTO Users(UserID) VALUES(?)")
	ErrorCheck(e)
	res, e := stmt.Exec(UserID)
	ErrorCheck(e)
	Users_id, e := res.LastInsertId()
	ErrorCheck(e)
	log.Printf("Insert Users_id %v", Users_id)

	stmt, e = db.Prepare("INSERT INTO Todos(Users_id, Text, Completed) VALUES(?,?,?)")
	ErrorCheck(e)
	res, e = stmt.Exec(Users_id, "Taste JavaScript", true)
	ErrorCheck(e)
	id, e := res.LastInsertId()
	log.Printf("Insert id %v", id)
	res, e = stmt.Exec(Users_id, "Buy a unicorn", false)
	ErrorCheck(e)
	id, e = res.LastInsertId()
	ErrorCheck(e)
	log.Printf("Insert id %v", id)

	r.db = db
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}
