package graph

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func InitDB() *sql.DB {

	db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/")
	if err != nil {
		panic(err)
	}
	log.Printf("### db ###, %v", db)

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("### Connected! ###")

	cmds := []string{
		"DROP DATABASE IF EXISTS Todos",
		"CREATE DATABASE IF NOT EXISTS Todos",
		"USE Todos",
		`CREATE TABLE Users (
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			UserID VARCHAR(32),
			TotalCount INT,
			CompletedCount INT
		) ENGINE=INNODB`,
		`CREATE TABLE Todos (
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			Users_id INT,
			Text VARCHAR(32),
			Completed INT,
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

	Db = db

	return db
}
