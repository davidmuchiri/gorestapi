package routes

import (
	"github.com/dinobambino7/gorestapi/handlers"
	"github.com/gorilla/mux"
)

//Routes function holds all the endpoints to our rest api
func Routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/{id}", handlers.AddProduct).Methods("POST")
	router.HandleFunc("/api/getProducts", handlers.GetProducts).Methods("GET")
	router.HandleFunc("/api/getProduct/{id}", handlers.GetProduct).Methods("GET")
	router.HandleFunc("/api/deleteProduct/{id}", handlers.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/api/deleteProducts", handlers.DeleteProducts).Methods("DELETE")

	return router
}
