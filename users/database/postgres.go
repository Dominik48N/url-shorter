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
	stmt, err := db.Prepare("INSERT INTO users (username, password) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, password)
	if err != nil {
		return err
	}
	return nil
}

func GetPassword(username string) (string, error) {
	var password string
	stmt, err := db.Prepare("SELECT password FROM users WHERE username = $1 LIMIT 1")
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	err = stmt.QueryRow(username).Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}

func IsUserNameExists(username string) (bool, error) {
	var u string
	stmt, err := db.Prepare("SELECT username FROM users WHERE username = $1 LIMIT 1")
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(username).Scan(&u)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
