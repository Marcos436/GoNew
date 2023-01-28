package main

import (
	"html/template"
	"net/http"
)

var templ = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	productAll := models.BuscarTodosProdutos()

	templ.ExecuteTemplate(w, "Index", productAll)

}
