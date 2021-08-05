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
		"DROP DATABASE IF EXISTS app",
		"CREATE DATABASE IF NOT EXISTS app",
		"USE app",
		`CREATE TABLE Users (
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			UserID TEXT,
			TotalCount INT,
			CompletedCount INT)`,
	}

	for i, cmd := range cmds {
		log.Printf("%v : %v", i, cmd)
		_, err = db.Exec(cmd)
		if err != nil {
			panic(err)
		}
	}

	return db
}
