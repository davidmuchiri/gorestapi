package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var uri = URI()

//Client client is a variable that we use to connect to collections in the database
var Client, err = mongo.NewClient(options.Client().ApplyURI(uri))

//ConnectDB is a function that connects to the database
func ConnectDB() {
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = Client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("connected to mongodb on mlab")

}
