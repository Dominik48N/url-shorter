package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-redis/redis"
	"github.com/go-redis/redis/v8"
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

	// Redis Cluster connection
	urlCachingDuration := getUrlCachingDuration()
	redisCluster := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    getRedisNodesFromEnv(),
		Password: os.Getenv("REDIS_PASSWORD"),
	})

	// HTTP Server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		path = path[1:] // The first is always a "/"!

		if !isValidURL(path) {
			handleUnknownURLs(w, r)
			return
		}

		var url string
		ctx := context.Background()
		cachedUrl, err := redisCluster.Get(ctx, "url:"+path).Result()
		if err != nil {
			err = database.QueryRow("SELECT redirect_url FROM urls WHERE link = $1", path).Scan(&url)
			if err != nil {
				url = "not_found"
			}
		} else {
			url = cachedUrl
		}

		if url != "not_found" {
			http.Redirect(w, r, url, http.StatusSeeOther)
		} else {
			handleUnknownURLs(w, r)
		}

		err = redisCluster.Set(ctx, "url:"+path, url, urlCachingDuration).Err()
		if err != nil {
			log.Fatalln(err)
		}
	})
	http.ListenAndServe(":3000", nil)

}

func getUrlCachingDuration() time.Duration {
	urlCachingDuration, err := strconv.Atoi(os.Getenv("URL_CACHING_TIME"))
	if err != nil {
		return 180 * time.Second
	}
	return time.Duration(urlCachingDuration)
}

func getRedisNodesFromEnv() []string {
	nodes := os.Getenv("REDIS_HOSTS")
	return strings.Split(nodes, ",")
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
