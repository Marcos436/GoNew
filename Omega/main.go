package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"
	_ "github.com/lib/pq"
)

type Produtos struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var templates = template.Must(template.ParseGlob("templates/pages/*.html"))

func main() {
	db := ConectarDatabase()
	defer db.Close()

	http.HandleFunc("/", Home)
	http.ListenAndServe(":8000", nil)

}
func ConectarDatabase() *sql.DB {

	conexao := "user=postgres dbname=postgres password=9891 host=localhost sslmode=disable"

	db, err := sql.Open("postgres", conexao)

	if err != nil {
		log.Println("Erro na tentativa de conexão DataBase")
		panic(err.Error())
	}
	return db
}

func Home(w http.ResponseWriter, r *http.Request) {

	db := ConectarDatabase()

	selectProdutos, err := db.Query("select * from produtos")

	if err != nil {
		log.Println("Erro na Query de seleção de produtos")
		panic(err.Error())
	}

	p := Produtos{}
	produtos := []Produtos{}

	for selectProdutos.Next() {
		var id int
		var nome string
		var descricao string
		var preco float64
		var quantidade int

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			log.Println("Erro na atribuição de valor das variáveis ")
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco                    // informando que as colunas do database e igual as variaveis
		p.Quantidade = quantidade

		produtos = append(produtos, p)

	}

	templates.ExecuteTemplate(w, "Home", produtos)
     defer db.Close()
}
