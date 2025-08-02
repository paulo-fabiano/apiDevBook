package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(writer http.ResponseWriter, statusCode int, data interface{}) {

	writer.WriteHeader(statusCode)
	writer.Header().Set("Content-Type", "application/json")
	
	if data != nil {
		err := json.NewEncoder(writer).Encode(data)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func Erro(writer http.ResponseWriter, statusCode int, err error) {

	JSON(writer, statusCode, struct {
		Erro string	`json:"erro"`
	}{
		Erro: err.Error(),
	})

}