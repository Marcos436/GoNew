package main

import (
	"net/http"
	"text/template"
	_ "github.com/lib/pq"
)

type Produtos struct{

	Nome string
	Descricao string
	Preco float64
	Quantidade int

}

var templates = template.Must(template.ParseGlob("templates/pages/*.html"))


func main() {


    http.HandleFunc("/", Home)
	http.ListenAndServe(":8000",nil)

}

func Home(w http.ResponseWriter, r *http.Request){

    produtos :=[]Produtos{
		{Nome:"Camisetas",
	     Descricao: "preta playBoy",
		 Preco: 199.90,
		 Quantidade: 500},

		{Nome:"short",
	     Descricao: "azul playBoy",
		 Preco: 59.90,
		 Quantidade: 50},

		 {Nome:"Boné",
	     Descricao: "cinza ligth",
		 Preco: 59.90,
		 Quantidade: 500},
	}

	templates.ExecuteTemplate(w, "Home", produtos)


}