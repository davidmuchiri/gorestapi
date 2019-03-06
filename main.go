package main

import (
	"log"

	"github.com/dinobambino7/gorestapi/db"
	"github.com/dinobambino7/gorestapi/routes"
	"github.com/dinobambino7/gorestapi/server"
)

const port = ":3000"

func main() {

	router := routes.Routes()
	server := server.New(router)
	db.ConnectDB()

	log.Println("starting server at port", port)
	log.Fatal(server.ListenAndServe())
}
