package seguranca

import "golang.org/x/crypto/bcrypt"

// Hash recebe uma string e retorna um Hash
func Hash(senha string) ([]byte, error) {

	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)

}

// VerificarSenha compara um Hash e uma senha enviada pelo usuário
func VerificarSenha(senhaHash string, senha string) error {

	return bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(senha))

}