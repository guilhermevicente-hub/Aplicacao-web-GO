package main

import (
	"net/http"
	"GoWeb/routes"

	_ "github.com/lib/pq"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}



// Precisei rodar (go mod init) no terminal para funcionar o (go get github.com/lib/pq)