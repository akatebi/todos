package graph

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "postgres-db-lb"
	port     = 5432
	user     = "postgres"
	password = "testpassword"
	dbname   = "/data/pgdata"
)

func init() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	log.Printf("### db ###, %v", db)

	// close database
	defer db.Close()

	fmt.Println("Connected!")
}
