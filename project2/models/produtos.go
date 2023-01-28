package models

import (
	"github.com/Marcos436/GoNew/db"
)

type Produto struct {
	Id         int64
	Nome       string
	Descrição  string
	Preço      float64
	Quantidade int64
}

func BuscarTodosProdutos() []Produto {

	db := db.ConectDatabase()

	selectAllProducts, err := db.Query("Select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produto := []Produto{}

	for selectAllProducts.Next() {

		var id int
		var nome string
		var descrição string
		var preço float64
		var quantidade int64

		err := selectAllProducts.Scan(&id, &nome, &descrição, &preço, &quantidade)

		if err != nil {
			panic(err.Error())
		}
		p.Id = int64(id)
		p.Nome = nome
		p.Descrição = descrição
		p.Preço = preço
		p.Quantidade = quantidade

		produto = append(produto, p)

	}
	defer db.Close()
	return produto

}
