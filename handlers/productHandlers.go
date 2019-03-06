package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dinobambino7/gorestapi/lib"
	"github.com/gorilla/mux"
)

var products = lib.OurProducts()

//AddProduct is a function that gets a new product from a http request and appends it to the products slice
func AddProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	params := mux.Vars(r)

	var product lib.Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	product.ProductID = params["id"]
	lib.AddProductToDB(product)
	json.NewEncoder(w).Encode("created new product")
}

//GetProduct is a function that returns a single product or false if product is not in the product list
func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Println(params)

	product, err := lib.GetProductFromDB(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(false)
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(product)

}

//GetProducts is a function that returns all products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var productss = lib.GetDataFromDB()

	if len(productss) == 0 {
		json.NewEncoder(w).Encode(false)
	}
	json.NewEncoder(w).Encode(productss)
}

//DeleteProduct deletes a single product from the products list
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Println(params["id"])

	res, err := lib.DeleteProductFromDB(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(false)
		log.Fatal(err)
		return
	}

	if res == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("could not find product")
		log.Println("could not find product")
		return
	}

	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(res)

	// for index, product := range products {
	// 	if product.ProductID == params["id"] {
	// 		products = append(products[:index], products[index+1:]...)
	// 		continue
	// 	}
	// }
	// fmt.Println(products)
	// json.NewEncoder(w).Encode(products)
}

//DeleteProducts is a function that deletes all products from the products list
func DeleteProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	products = nil
	json.NewEncoder(w).Encode("All products deleted")
	fmt.Println(products)
}
