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
  