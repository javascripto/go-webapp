package models

import (
	"database/sql"

	db "github.com/javascripto/go-webapp/db"
)

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

func CreateNewProduct(name string, description string, price float64, amount int) (sql.Result, error) {
	db := db.ConnectToDatabase()
	defer db.Close()
	statement, err := db.Prepare(`
		INSERT INTO products(name, description, price, amount)
		VALUES($1, $2, $3, $4)
	`)
	if err != nil {
		panic(err.Error())
	}
	result, err := statement.Exec(name, description, price, amount)
	if err != nil {
		panic(err.Error())
	}
	return result, nil
}

func DeleteProduct(id int) (sql.Result, error) {
	db := db.ConnectToDatabase()
	defer db.Close()
	statement, err := db.Prepare(`DELETE FROM products WHERE id=$1`)
	if err != nil {
		panic(err.Error())
	}
	result, err := statement.Exec(id)
	if err != nil {
		panic(err.Error())
	}
	return result, nil
}
