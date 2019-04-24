package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Users ref collection
var Users *mongo.Collection

// Employees ref collection
var Employees *mongo.Collection

// Documents ref collection
var Documents *mongo.Collection

// Categories ref collection
var Departments *mongo.Collection

// Goverments ref collection
var Goverments *mongo.Collection

// DBInit Database
func DBInit() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://10.127.130.213:27017")
	dbName := "citsk"
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	// export collections
	Users = client.Database(dbName).Collection("users")
	Employees = client.Database(dbName).Collection("employees")
	Documents = client.Database(dbName).Collection("documents")
	Departments = client.Database(dbName).Collection("departments")
	Goverments = client.Database(dbName).Collection("goverments")

	fmt.Println("Connected to MongoDB!")
	return client
}
