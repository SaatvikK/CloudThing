package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func login(res http.ResponseWriter, req *http.Request) {
	PathParams := mux.Vars(req) // Getting the parameters of the URL as a dictionary.
	res.Header().Set("Content-Type", "application/json")
	username, ok := PathParams["username"]
	if !ok {
		fmt.Println("USERNAME ERR")
		fmt.Println(ok)
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(`{ "result": false, "reason": "Username error.", "error": nil, "data": nill }`))
		return
	}

	password, ok := PathParams["pwd"]
	if !ok {
		fmt.Println("USERNAME ERR")
		fmt.Println(ok)
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(`{ "result": false, "reason": "Password error.", "error": nil, "data": nill }`))
	}

}
