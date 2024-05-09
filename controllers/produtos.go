package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"GoWeb/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html")) // carrega todos os templates com .html

func Index(w http.ResponseWriter, r *http.Request) {

	todosProdutos := models.BuscaProdutos()

	temp.ExecuteTemplate(w, "Index", todosProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro na conversão do proço: ", err)
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro na conversão da quantidade: ", err)
		}

		models.NovoProduto(nome, descricao, precoConvertido, quantidadeConvertida)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	models.DeleteProd(idProduto)
	http.Redirect(w, r, "/", 301)
}
