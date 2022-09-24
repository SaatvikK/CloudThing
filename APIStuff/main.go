// IMPORTING PACKAGES
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	routing()
}

func routing() {
	router := mux.NewRouter()
	APISubRouter := router.PathPrefix("/api/v1").Subrouter() // www.domain.com/api/v1/
	// This is how the client logins into the cloud. ConnType is intended to be either "discord" (if the client is)
	// a discord bot), "website" (if the client is the website), etc.
	APISubRouter.HandleFunc("/{UserID}:{password}@{ConnType}", postLogin).Methods(http.MethodPost)
	APISubRouter.HandleFunc("/{SessionID}/{WorkspaceName}/", postNewWorkspace).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8080", router)) // Start the server and listen on port 8080 (HTTPS).
}
