package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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

	// HTTP Server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		path = path[1:]

		if strings.Contains(path, "/") || strings.Contains(path, ".") {
			fmt.Fprintf(w, "Cannot find url!") // TODO: Maybe redirection to other website
			return
		}

		// TODO: Redis caching
		var url string
		err := database.QueryRow("SELECT redirect_url FROM urls WHERE link = $1", path).Scan(&url)
		if err != nil {
			fmt.Fprintf(w, "Cannot find url!") // TODO: Maybe redirection to other website
			return
		}

		http.Redirect(w, r, url, http.StatusSeeOther)
	})
	http.ListenAndServe(":3000", nil)

}
