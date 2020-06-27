package db

import (
	"database/sql"
	"log"

	"github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	// DB init
	pgUrl, err := pq.ParseURL("postgres://renishb@localhost:5432/golangjwtdb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	DB, err = sql.Open("postgres", pgUrl)
	if err != nil {
		log.Fatal(err)
	}
}
