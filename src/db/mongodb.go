package db

import (
	"context"
	"exception"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var database *mongo.Database

/**
<p>
  實作connection pool.
*/
func Connection() {
	log.Println("Connecting database ...")
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.Background(), clientOptions)
	exception.Fatal(err)

	err = client.Ping(context.Background(), nil)
	exception.Fatal(err)

	log.Println("Connected to MongoDB!")

	// handle for the trainers collection in the test database
	// collection := client.Database("test").Collection("trainers")

	database = client.Database("dev_db")
	//return
}

/**

 */
func GetInstance() *mongo.Database {
	fmt.Println("GetInstance ... ", database)
	return database
}

func DisConnection() {

	if database == nil {
		return
	}

	// Close the connection
	err := database.Client().Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Connection to MongoDB closed.")

}
