package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ccollazoscr/twittorgo/bd"
	"github.com/ccollazoscr/twittorgo/models"
)

/*ModificarPerfil modifica el perfil del usuario*/
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ingreso ModificarPerfil")
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}

	var status bool
	fmt.Printf("IDUsuario: %v\n", IDUsuario)
	status, err = bd.ModificoRegistro(t, IDUsuario)
	fmt.Printf("estatus: %v\n", status)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro. Reintente nuevamente "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado modificar el registro del usuario "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
