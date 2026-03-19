package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func Connect() {
	var err error
	db, err = pgxpool.New(context.TODO(), os.Getenv("DSN"))

	if err != nil {
		log.Fatal("Unable to connect database!")
	}

	log.Println("Succesfully connected to database")
}
