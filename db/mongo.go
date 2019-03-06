package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var uri = URI()

//Client we export client
var Client, err = mongo.NewClient(options.Client().ApplyURI(uri))

//ConnectDB connect to mongo db
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
