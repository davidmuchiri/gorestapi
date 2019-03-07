package lib

// Product is a structure of the product
type Product struct {
	ProductID          string  `json:"id" bson:"_id"`
	ProductName        string  `json:"productname" bson:"productname"`
	ProductDescription string  `json:"productdescription" bson:"productdescription"`
	ProductImg         string  `json:"productimg" bson:"productimg"`
	ProductPrice       float64 `json:"productprice" bson:"productprice"`
}


//OurProducts returns a slice of Products
func OurProducts() []Product {
	var products []Product
	return products
}
