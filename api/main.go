package main

import (
	"log"
	"net/http"

	"github.com/paulo-fabiano/apiDevBook/src/config"
	"github.com/paulo-fabiano/apiDevBook/src/router"
)

func main() {

	config.Carregar()

	log.Println("### Rodando API ###", )
	r := router.GerarRouter()
	log.Fatal(http.ListenAndServe(":5000", r))

}