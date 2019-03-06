package lib

import (
	"context"
	"fmt"
	"log"

	"github.com/dinobambino7/gorestapi/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client = db.Client
var collection = client.Database("gotest").Collection("products")

var products = OurProducts()
var ctx = context.Background()

//AddProductsToDB is a function that adds data to the mongodb database
func AddProductsToDB() {

	for _, product := range products {
		res, err := collection.InsertOne(ctx, bson.M{
			"_id":                product.ProductID,
			"productname":        product.ProductName,
			"productdescription": product.ProductDescription,
			"productimg":         product.ProductImg,
			"productprice":       product.ProductPrice,
		})

		if err != nil {
			log.Println(err)
		}

		id := res.InsertedID
		fmt.Println(id)
	}
}

//AddProductToDB adds a single product to the database
func AddProductToDB(product Product) {
	res, err := collection.InsertOne(ctx, bson.M{
		"_id":                product.ProductID,
		"productname":        product.ProductName,
		"productdescription": product.ProductDescription,
		"productimg":         product.ProductImg,
		"productprice":       product.ProductPrice,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.InsertedID)
}

//GetDataFromDB gets products from the database
func GetDataFromDB() []Product {

	findOptions := options.Find()
	findOptions.SetLimit(2)

	products := OurProducts()

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var el Product
		err := cur.Decode(&el)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, el)

	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())

	return products
}

//GetProductFromDB is a function
func GetProductFromDB(id string) (Product, error) {
	filter := bson.M{"_id": id}
	var el Product
	err := collection.FindOne(context.TODO(), filter).Decode(&el)

	var msg error

	if err != nil {
		msg = err
	}
	return el, msg
}

//DeleteProductFromDB is a function
func DeleteProductFromDB(id string) (int64, error) {

	deleteRes, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})

	var msg error

	if err != nil {
		msg = err

	}

	return deleteRes.DeletedCount, msg
}
