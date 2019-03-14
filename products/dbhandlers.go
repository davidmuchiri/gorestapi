package products

import (
	"context"
	"log"

	"github.com/dinobambino7/gorestapi/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Product struct
type Product struct {
	ProductID          string  `json:"id" bson:"_id"`
	ProductName        string  `json:"productname" bson:"productname"`
	ProductDescription string  `json:"productdescription" bson:"productdescription"`
	ProductImg         string  `json:"productimg" bson:"productimg"`
	ProductPrice       float64 `json:"productprice" bson:"productprice"`
}

var client = db.Client
var collection = client.Database("gotest").Collection("products")
var ctx = context.Background()

//AddProductsToDB adds a number of products to the database
func AddProductsToDB(products []Product) ([]interface{}, error) {

	ids := []interface{}{}
	var msg error

	for _, product := range products {
		res, err := collection.InsertOne(ctx, bson.M{
			"_id":                product.ProductID,
			"productname":        product.ProductName,
			"productdescription": product.ProductDescription,
			"prodctimg":          product.ProductImg,
			"productprice":       product.ProductPrice,
		})
		ids = append(ids, res.InsertedID)
		msg = err
	}

	return ids, msg
}

//AddProductToDB adds a single product to the database
func AddProductToDB(product Product) (interface{}, error) {

	res, err := collection.InsertOne(ctx, bson.M{
		"_id":                product.ProductID,
		"productname":        product.ProductName,
		"productdescription": product.ProductDescription,
		"productimg":         product.ProductImg,
		"productprice":       product.ProductPrice,
	})

	return res.InsertedID, err
}

//GetProductsFromDB gets all products from the database
func GetProductsFromDB() ([]Product, error) {
	findOptions := options.Find()
	var products []Product

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)

	for cur.Next(context.TODO()) {
		var product Product
		err := cur.Decode(&product)
		if err != nil {
			log.Fatal(err)
		}

		products = append(products, product)
	}

	return products, err
}

//GetProductFromDB gets a single product from the database
func GetProductFromDB(id string) (Product, error) {
	filter := bson.M{"_id": id}
	var product Product
	err := collection.FindOne(context.TODO(), filter).Decode(&product)

	return product, err
}

//UpdateProductDB updates a single product
func UpdateProductDB(id string, updates Product) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id": id}

	update := bson.D{
		{"$set", bson.M{
			"productname":        updates.ProductName,
			"productdescription": updates.ProductDescription,
			"productimg":         updates.ProductImg,
			"productprice":       updates.ProductPrice,
		}},
	}
	res, err := collection.UpdateOne(context.TODO(), filter, update)
	return res, err

}

//DeleteProductsFromDB deletes all products from the database
func DeleteProductsFromDB() (int64, error) {
	res, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	return res.DeletedCount, err
}

//DeleteProductFromDB  deletes a single product from the database
func DeleteProductFromDB(id string) (int64, error) {
	res, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	return res.DeletedCount, err
}
