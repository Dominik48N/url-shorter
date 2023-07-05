package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/Dominik48N/url-shorter/query/caching"
	"github.com/Dominik48N/url-shorter/query/database"
	_ "github.com/lib/pq"
)

const healthCheck = "health_check"
const notFound = "not_found"

var fallbackURL = strings.TrimSpace(os.Getenv("FALLBACK_URL"))

func main() {
	database.ConnectToPostgres()
	caching.ConnectToRedis()

	http.HandleFunc("/", handleURLRedirect)
	http.ListenAndServe(fmt.Sprintf(":%d", getHttpPort()), nil)
}

func getHttpPort() int {
	port, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		return 3000
	}
	return port
}

func handleURLRedirect(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:] // The first character is always a "/"

	if path == healthCheck {
		w.WriteHeader(http.StatusOK)
		return
	}

	if !isValidURL(path) {
		handleUnknownURLs(w, r)
		return
	}

	url, err := caching.GetURLFromCache(path)
	if err != nil {
		url, err = database.GetURLFromDatabase(path)
		if err != nil {
			url = notFound
		}
	}

	if url != notFound {
		http.Redirect(w, r, url, http.StatusSeeOther)
	} else {
		handleUnknownURLs(w, r)
	}

	err = caching.CacheURL(path, url)
	if err != nil {
		log.Println(err)
	}
}

func handleUnknownURLs(w http.ResponseWriter, r *http.Request) {
	if fallbackURL != "" {
		http.Redirect(w, r, fallbackURL, http.StatusSeeOther)
		return
	}

	http.Error(w, "URL not found!", http.StatusNotFound)
}

func isValidURL(url string) bool {
	if len(url) < 3 || len(url) > 12 {
		return false
	}

	match, _ := regexp.MatchString("^[A-Za-z]+$", url)
	return match
}
