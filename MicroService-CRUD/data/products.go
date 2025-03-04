package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct{
    ID          int     `json:"id"`
	NAME        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Sku         string  `json:"sku"`
	CreatedOn   string  `json:"createdOn"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}
type Products []*Product 

func(thisProduct *Product) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(thisProduct)
}
func(productlist Products) ToJson(w io.Writer)error{
	e := json.NewEncoder(w)
	return e.Encode(productlist)
}
//service methods for get all product from datastore 

func GetProducts()Products{
	return productList
}
var productList = Products{
	&Product{ID:          1,
		NAME:        "LATTE",
		Description: "Frothy Milky Coffee",
		Price:       2.45,
		Sku:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		NAME:        "ESSPRESSO",
		Description: "Strong coffe without Milk",
		Price:       1.99,
		Sku:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},

}