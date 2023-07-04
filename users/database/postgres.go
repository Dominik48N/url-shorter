package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var db *sql.DB

func ConnectToPostgres() {
	log.Println("Connecting to PostgreSQL...")
	connectionString := fmt.Sprintf(
		"postgresql://%s:%s@%s/%s",
		os.Getenv("POSTGRES_USERNAME"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_DATABASE"),
	)
	var err error
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connected to PostgreSQL!")
}

func InsertUser(username string, password string) error {
	_, err := db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, password)
	return err
}
