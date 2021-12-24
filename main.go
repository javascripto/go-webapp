package main

import (
	"fmt"
	"net/http"

	routes "github.com/javascripto/go-webapp/routes"
)

const SERVER_PORT = ":8000"

func main() {
	routes.LoadRoutes()
	fmt.Printf("App running on port %s\n", SERVER_PORT)
	http.ListenAndServe(SERVER_PORT, nil)
}
