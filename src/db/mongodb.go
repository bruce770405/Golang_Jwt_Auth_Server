package db

import (
	"context"
	"exception"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/**
<p>
  怎麼保有connection instance
  實作connection pool.
*/

func Connection() (database *mongo.Database) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("http://localhost:27017"))
	exception.Fatal(err)

	err = client.Ping(context.Background(), nil)
	exception.Fatal(err)

	fmt.Println("Connected to MongoDB!")

	// handle for the trainers collection in the test database
	// collection := client.Database("test").Collection("trainers")

	database = client.Database("dev_db")
	return
}

func DisConnection() {

	// Close the connection
	//err = client.Disconnect(context.TODO())

	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Connection to MongoDB closed.")

}
