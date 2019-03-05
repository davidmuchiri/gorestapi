package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dinobambino7/gorestapi/routes"
)

const port = ":3000"

func main() {
	router := routes.Routes()

	fmt.Println("server started at port....", port)
	log.Fatal(http.ListenAndServe(port, router))
}
