// Import packages
package main

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Standard struct template for returning information
type ReturnInfo struct {
	Result bool
	Reason interface{}
	Data   interface{}
	Error  interface{}
}

// HELPER FUNCS --------------------------------------------------------------
func conn() *mongo.Client { // Connecting to the MongoDB database
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println(reflect.TypeOf(client))
	return client
}

func checkLoginDetails(UserID string, pwd string) ReturnInfo { // Access the user's account and see if the login details are correct.
	client := conn()
	collection := client.Database("Main").Collection("users")
	var result bson.M
	err := collection.FindOne(context.TODO(), bson.D{{Key: "UserID", Value: UserID}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return ReturnInfo{false, "UserID incorrect.", nil, nil}
		}
		return ReturnInfo{false, "PANIC - unexpected error.", nil, err}
	}
	if result["password"] == pwd {
		return ReturnInfo{true, "Login details correct.", result, nil}
	} else {
		return ReturnInfo{false, "Password incorrect.", result, nil}
	}
}

// HANDLER FUNCS --------------------------------------------------------------
func postLogin(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.RemoteAddr)
	PathParams := mux.Vars(req) // Getting the parameters of the URL as a dictionary.
	res.Header().Set("Content-Type", "application/json")
	UserID, ok := PathParams["UserID"]
	fmt.Println(UserID)
	if !ok {
		fmt.Println("UserID ERR")
		fmt.Println(ok)
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintln(`{ "result": false, "reason": "UserID error.", "error": nil, "data": nil }`)))
		return
	}

	password, ok := PathParams["password"]
	if !ok {
		fmt.Println("PASSWORD ERR")
		fmt.Println(ok)
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintln(`{ "result": false, "reason": "Password error.", "error": nil, "data": nil }`)))
	}

	RR := checkLoginDetails(UserID, password)
	if RR.Result { // Login success
		SessionID := newSessionID(UserID, req.RemoteAddr)
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(fmt.Sprintln(`{ "result": true, "reason": "Congrats.", "data": { SessionID:`, SessionID, ` }, "error": nil }`)))
	} else { // Login failure
		res.WriteHeader(http.StatusUnauthorized)
		res.Write([]byte(fmt.Sprintln(`{ "result": false, "reason":`, RR.Reason, `, "data":`, RR.Data, `, "error":`, RR.Error, ` }`)))
	}
}

func postNewWorkspace(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.RemoteAddr)
	PathParams := mux.Vars(req) // Getting the parameters of the URL as a dictionary.
	res.Header().Set("Content-Type", "application/json")
	WorkspaceName, ok := PathParams["WorkspaceName"]
	if !ok {
		fmt.Println("Workspace ERR")
		fmt.Println(ok)
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintln(`{ "result": false, "reason": "Error in providing workspace name.", "error":`, ok, `, "data": nil }`)))
		return
	}
}
