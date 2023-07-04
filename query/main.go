package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	// PostgreSQL connection
	log.Println("Connect to postgres...")
	connectionString := "postgresql://" + os.Getenv("POSTGRES_USERNAME") + ":" + os.Getenv("POSTGRES_PASSWORD") + "@" + os.Getenv("POSTGRES_HOST") + "/" + os.Getenv("POSTGRES_DATABASE")
	database, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalln(err)
	}
	defer database.Close()
	log.Println("Connected to postgres!")

}
