package bd

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword encrypta un password en la bd*/
func EncriptarPassword(pass string) (string, error) {
	//Cantidad de pasadas sobre el texto para encryptarlo
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)

	return string(bytes), err
}
