package lib

//Product is a structure of a product
type Product struct {
	ProductID          string  `json:"id" bson:"_id"`
	ProductName        string  `json:"productname" bson:"productname"`
	ProductDescription string  `json:"productdescription" bson:"productdescription"`
	ProductImg         string  `json:"productimg" bson:"productimg"`
	ProductPrice       float64 `json:"productprice" bson:"productprice"`
}

// Products this is a slice that we are exporting

//OurProducts returns a slice of Products
func OurProducts() []Product {
	var products []Product

	// products = append(products, Product{ProductID: "abc", ProductName: "samsung a8", ProductDescription: "this is a black phone", ProductImg: "http://cloudinary.com/something.jpg", ProductPrice: 39448})

	// products = append(products, Product{ProductID: "343fefr3", ProductName: "iphone 6s", ProductDescription: "this is an iphone", ProductImg: "http://cloudinary.com/something.jpg", ProductPrice: 67000})

	// products = append(products, Product{ProductID: "4r4r44", ProductName: "huawei mate 8", ProductDescription: "this is a matte phone", ProductImg: "http://cloudinary.com/something.jpg", ProductPrice: 70000})

	// products = append(products, Product{ProductID: "vrg444454", ProductName: "pixel", ProductDescription: "this is a pixel phone", ProductImg: "http://cloudinary.com/something.jpg", ProductPrice: 80000})

	// products = append(products, Product{ProductID: "12r33f67k7", ProductName: "samsung a7", ProductDescription: "this is a samsung a7 2016 phone", ProductImg: "http://cloudinary.com/something.jpg", ProductPrice: 47448})

	return products
}
