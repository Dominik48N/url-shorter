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

func GetURLFromDatabase(path string) (string, error) {
	var url string
	stmt, err := db.Prepare("SELECT redirect_url FROM urls WHERE link = $1 LIMIT 1")
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	err = stmt.QueryRow(path).Scan(&url)
	if err != nil {
		return "", err
	}
	return url, nil
}
