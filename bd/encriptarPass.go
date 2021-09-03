package bd

import "golang.org/x/crypto/bcrypt"

/*EncriptarPass Funcion para encriptar la pass del user antes de registrar*/
func EncriptarPass(pass string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost) //Pasamos pass como byte

	return string(bytes), err
}
