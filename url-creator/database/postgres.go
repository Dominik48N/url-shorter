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

func CreateURL(id, url, username string) error {
	_, err := db.Exec("INSERT INTO urls (link, redirect_url, owner) VALUES ($1, $2, $3)", id, url, username)
	return err
}

func CheckLinkExists(link string) (bool, error) {
	err := db.QueryRow("SELECT link FROM urls WHERE link = $1 LIMIT 1", link).Scan()
	if err == nil {
		return true, nil
	}

	if err == sql.ErrNoRows {
		return false, nil
	}

	return false, err
}
