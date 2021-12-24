package main

import (
	"fmt"
	"net/http"
	"text/template"

	models "github.com/javascripto/go-webapp/models"
)

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
	products := models.GetAllProducts()
	templates.ExecuteTemplate(w, "Index", products)
}
