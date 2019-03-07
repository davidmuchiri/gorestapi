package routes

import (
	"github.com/dinobambino7/gorestapi/routehandlers"
	"github.com/gorilla/mux"
)

//Routes function holds all the endpoints to our rest api
func Routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/{id}", routehandlers.AddProduct).Methods("POST")
	router.HandleFunc("/api/getProducts", routehandlers.GetProducts).Methods("GET")
	router.HandleFunc("/api/getProduct/{id}", routehandlers.GetProduct).Methods("GET")
	router.HandleFunc("/api/updateProduct/{id}", routehandlers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/deleteProduct/{id}", routehandlers.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/api/deleteProducts", routehandlers.DeleteProducts).Methods("DELETE")

	return router
}
