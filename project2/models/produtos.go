package models

import "github.com/Marcos436/GoNew"

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
		var quantidade int64
		var descrição string
		var nome string
		var preço float64

		err := selectAllProducts.Scan(&id, &quantidade, &nome, &preço, &descrição)

		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Quantidade = quantidade
		p.Descrição = descrição
		p.Preço = preço

		produto = append(produto, p)

	}
	defer db.Close()
	return produto

}
