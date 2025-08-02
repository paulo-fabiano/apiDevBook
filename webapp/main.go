package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/paulo-fabiano/apiDevBook/src/router"
	"github.com/paulo-fabiano/apiDevBook/src/router/rotas"
	"github.com/paulo-fabiano/apiDevBook/utils"
)

func main() {

	utils.CarregarTemplates()
	router := router.Gerar()
	rotas.Configurar(router) // Carregando as Rotas
	
	fmt.Println("Rodando WEBAPP")
	log.Fatal(http.ListenAndServe(":3000", router)) 

}