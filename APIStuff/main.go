package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

}

func routing() {
	router := mux.NewRouter()
	APISubRouter := router.PathPrefix("/api/v1").Subrouter()
	APISubRouter.HandleFunc("/{username}:{password}@{ConnType}", login).Methods(http.MethodConnect)
	log.Fatal(http.ListenAndServe(":8080", router))
}
