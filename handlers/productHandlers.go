package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dinobambino7/gorestapi/lib"
	"github.com/gorilla/mux"
)

var products = lib.OurProducts()

func AddProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var product lib.Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	product.ProductID = params["id"]
	products = append(products, product)
	fmt.Println(products)
	json.NewEncoder(w).Encode("created new product")
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params)

	for _, item := range products {
		if item.ProductID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(false)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params["id"])
	for index, product := range products {
		if product.ProductID == params["id"] {
			products = append(products[:index], products[index+1:]...)
			continue
		}
	}
	fmt.Println(products)
	json.NewEncoder(w).Encode(products)
}

func DeleteProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products = nil
	json.NewEncoder(w).Encode("All products deleted")
	fmt.Println(products)
}
