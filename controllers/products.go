package controllers

import (
	"net/http"
	"text/template"

	"github.com/javascripto/go-webapp/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func ListProducts(w http.ResponseWriter, r *http.Request) {
	products := models.GetAllProducts()
	templates.ExecuteTemplate(w, "Index", products)
}
