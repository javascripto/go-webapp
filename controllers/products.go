package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/javascripto/go-webapp/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func ListProducts(w http.ResponseWriter, r *http.Request) {
	products := models.GetAllProducts()
	templates.ExecuteTemplate(w, "Index", products)
}

func NewProduct(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "NewProduct", nil)
}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			log.Println("Error on converting price to float64:", err)
		}
		amount, err := strconv.Atoi(r.FormValue("amount"))
		if err != nil {
			log.Println("Error on converting amount to int:", err)
		}
		models.CreateNewProduct(name, description, price, amount)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println("Error on converting id to int:", err)
	}
	models.DeleteProduct(id)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
