package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Dominik48N/url-shorter/stats-saver/database"
	"github.com/julienschmidt/httprouter"
)

const apiVersion = "v1"

func main() {
	database.ConnectToPostgres()

	router := httprouter.New()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", getHttpPort()), router))
}

func getHttpPort() int {
	port, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		return 3000
	}
	return port
}
