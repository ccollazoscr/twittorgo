package routers

import (
	"errors"
	"strings"

	"github.com/ccollazoscr/twittorgo/bd"
	"github.com/ccollazoscr/twittorgo/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*Email del usuairo que esta haciendo una petición*/
var Email string

/*IDUsuario identificador del usuario*/
var IDUsuario string

/*ProcesoToken procesa un token JWT*/
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
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("Token invalido")
	}

	return claims, false, string(""), err
}
