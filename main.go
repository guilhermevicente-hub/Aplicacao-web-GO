package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html")) // carrega todos os templates com .html

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	produtos := []Produto {
		{Nome: "Camiseta", Descricao: "Azul escura", Preco: 39, Quantidade: 5},
		{"Tenis", "Nike Dunk", 280, 4},
		{"Fone", "Sem fio", 59, 2},
		{"Notebook", "Dell Usado", 1800, 2},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
