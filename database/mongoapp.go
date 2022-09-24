package main

import (
	"context"
	"fmt"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

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
	err := collection.FindOne(context.TODO(), bson.D{{Key: "email", Value: email}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return
		}
		panic(err)
	}
	fmt.Println(reflect.TypeOf(result))
}

func main() {
	checkReadDB("saatvikrk@gmail.com", "HI")
}
