package main

import (
	"net/http"
	"text/template"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Amount      int
}

const SERVER_PORT = ":8000"

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(SERVER_PORT, nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{Name: "T-shirt", Description: "Blue, really pretty t-shirt", Price: 29, Amount: 10},
		{Name: "Laptop", Description: "Very fast laptop", Price: 1999, Amount: 2},
		{Name: "Sneakers", Description: "Very cool sneakers", Price: 99, Amount: 5},
	}
	templates.ExecuteTemplate(w, "Index", products)
}
