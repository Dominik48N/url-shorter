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
	stmt, err := db.Prepare("INSERT INTO urls (link, redirect_url, owner) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, url, username)
	if err != nil {
		return err
	}
	return nil
}

func CheckLinkExists(link string) (bool, error) {
	var l string
	stmt, err := db.Prepare("SELECT link FROM urls WHERE link = $1 LIMIT 1")
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(link).Scan(&l)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
