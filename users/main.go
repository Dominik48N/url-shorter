package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Dominik48N/url-shorter/users/database"
	"github.com/Dominik48N/url-shorter/users/hashing"
	"github.com/Dominik48N/url-shorter/users/user"
	"github.com/julienschmidt/httprouter"
)

const apiVersion = "v1"

func main() {
	database.ConnectToPostgres()

	router := httprouter.New()
	router.POST("/v"+apiVersion+"/register", registerHandler)
	router.POST("/v"+apiVersion+"/login", loginHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", getHttpPort()), router))
}

func getHttpPort() int {
	port, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		return 3000
	}
	return port
}

func registerHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user user.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usernameExists, err := database.IsUserNameExists(user.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if usernameExists {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	hashedPassword := hashing.HashPassword(user.Password)

	err = database.InsertUser(user.Username, hashedPassword)
	if err != nil {
		log.Fatalln(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func loginHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var userType user.User
	err := json.NewDecoder(r.Body).Decode(&userType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := database.GetPassword(userType.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !hashing.CheckPassword(userType.Password, hashedPassword) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := user.AuthUser(userType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(token))
}
