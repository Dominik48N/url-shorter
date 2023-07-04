package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Dominik48N/url-shorter/users/database"
	"github.com/Dominik48N/url-shorter/users/user"
	"github.com/gorilla/mux"
)

func main() {
	database.ConnectToPostgres()

	router := mux.NewRouter()
	router.HandleFunc("/register", registerHandler).Methods("POST")

	http.ListenAndServe(fmt.Sprintf(":%d", getHttpPort()), router)
}

func getHttpPort() int {
	port, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		return 3000
	}
	return port
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var user user.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: Hash Password

	err = database.InsertUser(user.Username, user.Password)
	if err != nil {
		log.Fatalln(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
