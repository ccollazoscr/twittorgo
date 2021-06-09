package routers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ccollazoscr/twittorgo/bd"
	"github.com/ccollazoscr/twittorgo/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*Email del usuairo que esta haciendo una petición*/
var Email string

/*IDUsuario identificador del usuario*/
var IDUsuario string

/*ProcesoToken proceso token para extraer sus valores */
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MastersDelDesarrollo_grupodeFacebook")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, ID := bd.ChequeoYaExisteUsuario(claims.Email)

		if encontrado == true {
			Email = claims.Email
			IDUsuario = ID

			fmt.Println(ID)
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token Inválido")
	}
	return claims, false, string(""), err
}
