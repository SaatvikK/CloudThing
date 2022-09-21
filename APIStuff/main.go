package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

}

func routing() {
	router := mux.NewRouter()
	APISubRouter := router.PathPrefix("/api/v1").Subrouter()
	APISubRouter.HandleFunc("/{username}:{password}@{ConnType}").Methods(http.MethodConnect)

}
