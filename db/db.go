package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectToDatabase() *sql.DB {
	connectionString := "user=docker dbname=go_store password=docker host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err.Error())
	}
	return db
}
