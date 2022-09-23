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

// HELPER FUNCS --------------------------------------------------------------
func conn() *mongo.Client {
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

func checkReadDB(email string, pwd string) {
	client := conn()
	collection := client.Database("Main").Collection("users")
	var result bson.M
	err := collection.FindOne(context.TODO(), bson.D{{"email", email}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return
		}
		panic(err)
	}
	fmt.Println(result)
}

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
