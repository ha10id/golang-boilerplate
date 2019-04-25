package db

import (
	"context"
	"fmt"
	"log"

	"../config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Users ref collection
var Users *mongo.Collection

// Assets ref collection
var Assets *mongo.Collection

// Employees ref collection
var Employees *mongo.Collection

// News ref collection
var News *mongo.Collection

// Documents ref collection
var Documents *mongo.Collection

// Categories ref collection
var Departments *mongo.Collection

// Goverments ref collection
var Goverments *mongo.Collection

// DBInit Database
func Init() *mongo.Client {
	c := config.GetConfig()
	// log.Println(c.GetString("app.mongoEndpoint"))
	// Set client options
	clientOptions := options.Client().ApplyURI(c.GetString("app.mongoEndpoint"))
	dbName := c.GetString("app.mongoDB")
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
	Assets = client.Database(dbName).Collection("assets")
	Employees = client.Database(dbName).Collection("employees")
	Documents = client.Database(dbName).Collection("documents")
	Departments = client.Database(dbName).Collection("departments")
	Goverments = client.Database(dbName).Collection("goverments")
	News = client.Database(dbName).Collection("news")

	fmt.Println("Connected to MongoDB!")
	return client
}
