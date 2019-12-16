package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "1234"
	dbname   = "restaurantdb"
)

func main() {
	// session, err := store.Get(r, "session-name")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	dbconn, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer dbconn.Close()

	if err := dbconn.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Database connection is established...")
}
