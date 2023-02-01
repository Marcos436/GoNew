package config

import (
	"database/sql"
	"log"
)


func DataBase() *sql.DB {
  
	Conexao := "user=postgres ddname=Usuario password=9891 host=localhost sslmode=disable"

	DB, err :=sql.Open("postgres", Conexao)

	if err != nil{
		log.Println("Erro na conexao com Databse",Conexao)
		panic(err.Error())
	}
	return DB
}