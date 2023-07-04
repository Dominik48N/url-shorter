package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	_ "github.com/lib/pq"
)

var fallbackUrl = strings.TrimSpace(os.Getenv("FALLBACK_URL"))

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
		path = path[1:] // The first is always a "/"!

		if !isValidURL(path) {
			handleUnknownURLs(w, r)
			return
		}

		// TODO: Redis caching

		var url string
		err := database.QueryRow("SELECT redirect_url FROM urls WHERE link = $1", path).Scan(&url)
		if err != nil {
			handleUnknownURLs(w, r)
			return
		}

		http.Redirect(w, r, url, http.StatusSeeOther)
	})
	http.ListenAndServe(":3000", nil)

}

func handleUnknownURLs(w http.ResponseWriter, r *http.Request) {
	if len(fallbackUrl) == 0 {
		http.Redirect(w, r, fallbackUrl, http.StatusSeeOther)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "This URL was not found.")
}

func isValidURL(url string) bool {
	if len(url) < 3 || len(url) > 12 {
		return false
	}

	match, _ := regexp.MatchString("^[A-Za-z]+$", url)
	return match
}
