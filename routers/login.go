package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ccollazoscr/twittorgo/bd"
	"github.com/ccollazoscr/twittorgo/jwt"
	"github.com/ccollazoscr/twittorgo/models"
)

/*Login realiza el login*/
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Usuario y/o contraseña invalidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una constraseña de al menos 6 caracteres", 400)
		return
	}

	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if existe == false {
		http.Error(w, "Usuario y/o contraseña invalidos "+err.Error(), 400)
		return
	}

	jwtkey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrión error al intentar generar el token "+err.Error(), 400)
		return
	}

	res := models.RespuestaLogin{
		Token: jwtkey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtkey,
		Expires: expirationTime,
	})
}
