package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	// DB init
	pgUrl, err := pq.ParseURL(os.Getenv("DBURL"))
	if err != nil {
		log.Fatal(err)
	}

	DB, err = sql.Open("postgres", pgUrl)
	if err != nil {
		log.Fatal(err)
	}
}
