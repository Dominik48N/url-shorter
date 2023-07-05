package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Dominik48N/url-shorter/url-creator/authorization"
	"github.com/Dominik48N/url-shorter/url-creator/database"
	"github.com/Dominik48N/url-shorter/url-creator/generator"
	"github.com/julienschmidt/httprouter"
)

const apiVersion = "v1"

func main() {
	database.ConnectToPostgres()

	router := httprouter.New()
	router.POST("/v"+apiVersion+"/create", authorization.AuthMiddleware(createHandler))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", getHttpPort()), router))
}

func getHttpPort() int {
	port, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		return 3000
	}
	return port
}

func createHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var createRequest CreateRequest
	err := json.NewDecoder(r.Body).Decode(&createRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: Check valid target url (+ abuse/blacklist!)

	username := r.Context().Value("username").(string)
	id, err := generator.GenerateRandomLink()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	database.CreateURL(id, createRequest.RedirectUrl, username)

	w.WriteHeader(http.StatusCreated)
}
