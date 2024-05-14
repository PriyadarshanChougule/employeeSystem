package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbConnection *mongo.Database

type Collections struct {
	Employees *mongo.Collection
	Login     *mongo.Collection
}

var AllCollections Collections

func init() {
	fmt.Println("initializing db package...")
	connectDB()
}

func connectDB() {

	uri := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), uri)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to DB...")
	dbConnection := client.Database("G10CRUD")

	AllCollections = Collections{
		Employees: dbConnection.Collection("employees"),
		Login:     dbConnection.Collection("login"),
	}

}

func GetAllCollections() Collections {
	return AllCollections
}
