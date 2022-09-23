package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func login(res http.ResponseWriter, req *http.Request) {
	PathParams := mux.Vars(req) // Getting the parameters of the URL as a dictionary.
	res.Header().Set("Content-Type", "application/json")
	username, ok := PathParams["username"]
}
