package autenticacao

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/paulo-fabiano/apiDevBook/src/config"
)

// CriarToken gera um token com as permissões do usuário
func CriarToken(usuarioID uint64) (string, error) {

	permissoes := jwt.MapClaims{}

	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioID"] = usuarioID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)

	return token.SignedString([]byte(config.SecretKey)) // SECRET

}

// ValidarToken valida se o token passado na requisição é válido.
func ValidarToken(request *http.Request) error {

	tokenString := extrairToken(request)
	token, err := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Token inválido, faça login novamente")

}

// ExtrairUsuarioID retorna o ID do usuário que está salvo no Token
func ExtrairUsuarioID(request *http.Request) (uint64, error) {
	
	tokenString := extrairToken(request)
	token, err := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if err != nil {
		return 0, err
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioID, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["usuarioID"]), 10, 64)
		if err != nil {
			return 0, nil
		}

		return usuarioID, nil
	}

	return 0, errors.New("Token inválido")
 

}

func extrairToken(request *http.Request) string {

	token := request.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 { // Bearer vmkmfdkjnfdjn... 
		return strings.Split(token, " ")[1] // Pegamos só o Token
	}

	return ""

}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado")
	}

	return config.SecretKey, nil

}