package models

import (
	"GoWeb/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaProdutos() []Produto {

	db := db.ConectaBanco()

	selectProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}

func NovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaBanco()

	insertDados, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade)values($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insertDados.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeleteProd(id string){
	db := db.ConectaBanco()

	deletarProduto, err := db.Prepare("delete from produtos where id=$1")

	if err != nil {
		panic(err.Error())
	}

	deletarProduto.Exec(id)
	defer db.Close()
}