package products

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dinobambino7/gorestapi/utils"

	"github.com/gorilla/mux"
)

//AddProduct gets product and adds product to db
func AddProduct(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	response := make(map[string]interface{})

	var product Product
	_ = json.NewDecoder(req.Body).Decode(&product)
	product.ProductID = params["id"]
	_, err := AddProductToDB(product)

	if err != nil {
		response = utils.Message(false, "could not add product to db", nil)
		log.Println("could not add product to db", err)
		utils.Response(res, response)
		return
	}

	response = utils.Message(true, "created a new product", params["id"])
	utils.Response(res, response)
}

//GetProduct gets product and adds product to db
func GetProduct(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	response := make(map[string]interface{})

	product, err := GetProductFromDB(id)

	if err != nil {
		response = utils.Message(false, "could not get product", nil)
		log.Println("could not get product from db", err)
		utils.Response(res, response)
		return
	}

	response = utils.Message(true, "we got a product", product)
	utils.Response(res, response)
}

//GetProducts gets product and adds product to db
func GetProducts(res http.ResponseWriter, req *http.Request) {
	response := make(map[string]interface{})

	products, err := GetProductsFromDB()

	if err != nil {
		response = utils.Message(false, "could not get products", "")
		log.Println("could not get products", err)
		utils.Response(res, response)
		return
	}
	if len(products) == 0 {
		response = utils.Message(true, "no products", products)
		utils.Response(res, response)
		return
	}

	response = utils.Message(true, "here are your products", products)
	utils.Response(res, response)
}

//UpdateProduct gets product and adds product to db
func UpdateProduct(res http.ResponseWriter, req *http.Request) {
	response := make(map[string]interface{})
	params := mux.Vars(req)
	var updates Product

	_ = json.NewDecoder(req.Body).Decode(&updates)
	result, err := UpdateProductDB(params["id"], updates)

	if err != nil {
		response = utils.Message(false, "could not update product", "")
		log.Println("could not update product", err)
		utils.Response(res, response)
		return
	}

	response = utils.Message(true, "product updated", result.MatchedCount)
	utils.Response(res, response)
}

//DeleteProduct gets product and adds product to db
func DeleteProduct(res http.ResponseWriter, req *http.Request) {
	response := make(map[string]interface{})

	params := mux.Vars(req)
	id := params["id"]

	count, err := DeleteProductFromDB(id)

	if err != nil {
		response = utils.Message(false, "could not delete product", "")
		log.Println("could not delete data", err)
		utils.Response(res, response)
		return
	}
	if count == 0 {
		response = utils.Message(true, "product not found", "")
		utils.Response(res, response)
		return
	}

	response = utils.Message(true, "product deleted", count)
	utils.Response(res, response)
}

//DeleteProducts gets product and adds product to db
func DeleteProducts(res http.ResponseWriter, req *http.Request) {
	response := make(map[string]interface{})
	count, err := DeleteProductsFromDB()

	if err != nil {
		response = utils.Message(false, "could not delete products", "")
		log.Println("could not delete products", err)
		utils.Response(res, response)
		return
	}
	if count == 0 {
		response = utils.Message(true, "no products", "")
		utils.Response(res, response)
		return
	}

	response = utils.Message(true, "products deleted", count)
	utils.Response(res, response)
}
