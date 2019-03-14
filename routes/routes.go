package routes

import (
	"github.com/dinobambino7/gorestapi/middleware"
	"github.com/dinobambino7/gorestapi/products"
	"github.com/dinobambino7/gorestapi/users"
	"github.com/gorilla/mux"
)

//Routes function holds all the endpoints to our rest api
func Routes() *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.JwtAuthentication)

	router.HandleFunc("/api/{id}", products.AddProduct).Methods("POST")
	router.HandleFunc("/api/getProducts", products.GetProducts).Methods("GET")
	router.HandleFunc("/api/getProduct/{id}", products.GetProduct).Methods("GET")
	router.HandleFunc("/api/updateProduct/{id}", products.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/deleteProduct/{id}", products.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/api/deleteProducts", products.DeleteProducts).Methods("DELETE")

	router.HandleFunc("/api/users/register", users.Register).Methods("POST")
	router.HandleFunc("/api/users/login", users.Authenticate).Methods("POST")

	return router
}
