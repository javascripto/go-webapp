# Go Store

## Starting a new html template

- Create a new directory named as `templates` and create a new `index.html` file inside it.
- You can use the `{{ define "Index" }}` tag to define a template and `{{ end }}` tag to end the template.
- Add the following code to the `index.html` file.

  ```html
  {{ define "Index" }}
  <!DOCTYPE html>
  <html>
    <head></head>
  <body>
    <h1>Go Store</h1>
  </body>
  </html>
  {{end}}
  ```
- With the next terminal code you can generate a similar simpler file
  ```sh
  md templates
  touch templates/index.html
  echo '{{define "Index"}}\n<h1>Go Store</h1>\n{{end}}' > templates/index.html
  ```
- Now you are ready to start your go server code.
- Create the `main.go` file to start typing your go application. `echo 'package main' > main.go`
- Inside that file you can type the following code to load the template and create a http server that will render it.

  ```go
  package main

  import (
  	"net/http"
  	"text/template"
  )

  const SERVER_PORT = ":8000"

  var templates = template.Must(template.ParseGlob("templates/*.html"))

  func main() {
  	http.HandleFunc("/", index)
  	http.ListenAndServe(SERVER_PORT, nil)
  }

  func index(w http.ResponseWriter, r *http.Request) {
  	templates.ExecuteTemplate(w, "Index", nil)
  }
  ```
- Now you can run the application with the following command: `go run main.go`

## Binding data to the template

- You can bind data to the template using the `{{ . }}` tag.
- Create a new struct type to store Product data.
- Create a new slice of Product structs to bind to the template.
- Add a range tag to the template to iterate over the slice of Product structs.
- Code:
  ```go
  type Product struct {
  	Name        string
  	Description string
  	Price       float64
  	Amount      int
  }
  // request handler
  func index(w http.ResponseWriter, r *http.Request) {
  	products := []Product{
  		{Name: "T-shirt", Description: "Blue, really pretty t-shirt", Price: 29, Amount: 10},
  		{Name: "Laptop", Description: "Very fast laptop", Price: 1999, Amount: 2},
  		{Name: "Sneakers", Description: "Very cool sneakers", Price: 99, Amount: 5},
  	}
  	templates.ExecuteTemplate(w, "Index", products)
  }
  ```
  ```html
  <!-- Products loop and binding -->
  {{range .}}
    <tr>
      <td>{{.Name}}</td>
      <td>{{.Description}}</td>
      <td>{{.Price}}</td>
      <td>{{.Amount}}</td>
    </tr>
  {{end}}
  ```

## Connect application to postgres docker image

- Run a new docker container for postgres and start it. 
```
docker run --name go_store_postgres \
  -e POSTGRES_USER=docker \
  -e POSTGRES_PASSWORD=docker \
  -p 5432:5432 \
  -d postgres

docker start go_store_postgres

docker stop go_store_postgres
```
- Create a new database named as `go_store` and a new table named as `products` with the fields `id`, `name`, `description`, `price` and `amount`.
- Insert some records to the table.
- Install the library that will connect the application to the postgres database. `go get github.com/lib/pq`
- Now import the package pq to create a new connection to the database. `import ( _ "github.com/lib/pq" )`
- This is the code that returns a new connection to the database.
  ```go
  func connectToDatabase() *sql.DB {
  	connectionString := "user=docker dbname=go_store password=docker host=localhost sslmode=disable"
  	db, err := sql.Open("postgres", connectionString)
  	if err != nil {
  		panic(err.Error())
  	}
  	return db
  }
  ```
- Considering the `products` table has a columns named as `id`, the struct also have a field to represent that column.
- With the connection we can make queries to the database fetching for products.

  ```go
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
  ```