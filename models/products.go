package models

import db "github.com/javascripto/go-webapp/db"

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Amount      int
}

func GetAllProducts() []Product {
	db := db.ConnectToDatabase()
	defer db.Close()

	productsQuery, err := db.Query("SELECT * FROM products")
	product := Product{}
	products := []Product{}

	if err != nil {
		panic(err.Error())
	}

	for productsQuery.Next() {
		err = productsQuery.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Amount)
		if err != nil {
			panic(err.Error())
		}
		products = append(products, product)
	}

	return products
}
