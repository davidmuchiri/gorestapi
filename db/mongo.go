package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var e = godotenv.Load("config/config.env")

var uri = os.Getenv("MLAB_Credentials")

//Client client is a variable that we use to connect to collections in the database
var Client, err = mongo.NewClient(options.Client().ApplyURI(uri))

//ConnectDB is a function that connects to the database
func ConnectDB() {
	if e != nil {
		fmt.Println(e)
	}

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = Client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("connected to mongodb on mlab")

}
