package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

func connectToDatabase() *sql.DB {
	connectionString := "user=docker dbname=go_store password=docker host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Amount      int
}

const SERVER_PORT = ":8000"

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/favicon.ico", doNothing)
	fmt.Printf("App running on port %s\n", SERVER_PORT)
	http.ListenAndServe(SERVER_PORT, nil)
}

func doNothing(w http.ResponseWriter, r *http.Request) {}

func index(w http.ResponseWriter, r *http.Request) {
	db := connectToDatabase()
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
	templates.ExecuteTemplate(w, "Index", products)
}
