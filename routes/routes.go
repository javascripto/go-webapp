package routes

import (
	"net/http"

	"github.com/javascripto/go-webapp/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.ListProducts)
	http.HandleFunc("/new", controllers.NewProduct)
	http.HandleFunc("/insert", controllers.InsertProduct)
	http.HandleFunc("/delete", controllers.DeleteProduct)
	http.HandleFunc("/edit", controllers.EditProduct)
	http.HandleFunc("/update", controllers.UpdateProduct)
	http.HandleFunc("/favicon.ico", doNothing)
}

func doNothing(w http.ResponseWriter, r *http.Request) {}
