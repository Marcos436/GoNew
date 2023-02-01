package main

import (
	"net/http"
	"text/template"
)


var templates = template.Must(template.ParseGlob("templates/pages/*.html"))


func main() {


    http.HandleFunc("/", Home)
	http.ListenAndServe(":8000",nil)

}

func Home(w http.ResponseWriter, r *http.Request){

	templates.ExecuteTemplate(w, "Home", nil)


}