package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/Dominik48N/url-shorter/url-creator/database"
	"github.com/gorilla/mux"
)

const apiVersion = "v1"

func main() {
	database.ConnectToPostgres()

	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/v" + apiVersion).Subrouter()
	apiRouter.HandleFunc("/create", createHandler).Methods("POST")

	http.ListenAndServe(fmt.Sprintf(":%d", getHttpPort()), router)
}

func getHttpPort() int {
	port, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		return 3000
	}
	return port
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	var createRequest CreateRequest
	err := json.NewDecoder(r.Body).Decode(&createRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: Check valid target url (+ abuse/blacklist!)

	username := "cba" // TODO: Check if user is logged in (+ username get)
	id := "abc"       // TODO: Generate unused id

	database.CreateURL(id, createRequest.RedirectUrl, username)

	w.WriteHeader(http.StatusCreated)
}
