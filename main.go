package main

import (
	"log"

	"github.com/dinobambino7/gorestapi/db"
	"github.com/dinobambino7/gorestapi/routes"
	"github.com/dinobambino7/gorestapi/server"
)

const port = ":3000"

func main() {

	// gets the routes from the routes package
	router := routes.Routes()

	// passed the mux router to the server
	server := server.New(router)

	// connects to the database
	db.ConnectDB()

	log.Println("starting server at port", port)
	log.Fatal(server.ListenAndServe())
}
