package middlew

import (
	"net/http"

	"github.com/ccollazoscr/twittorgo/routers"
)

/*ValidoJWT validar JWT que nos viene en cada petición*/
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el token "+err.Error(), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	}
}
