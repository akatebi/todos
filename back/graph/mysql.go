package graph

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// const (
// 	host     = "postgres-db-lb"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "testpassword"
// 	dbname   = "/data/pgdata"
// )

func InitDB() *sql.DB {
	// connection string
	// conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// open database
	log.Printf("### MYSQL ###")
	const conn = "root:password@tcp(mysql:3306)/"
	db, err := sql.Open("mysql", conn)
	if err != nil {
		panic(err)
	}
	log.Printf("### db ###, %v", db)

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS todos")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("USE todos")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected!")

	return db
}
