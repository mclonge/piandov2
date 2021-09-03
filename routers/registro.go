package routers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mclonge/piando/bd"
	"github.com/mclonge/piando/models"
)

func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t) //Body es un objeto stream que solo se puede leer una ve. Una vez leido se destruye
	//Ya no puedo usar Body en mas sitios
	if err != nil {
		log.Fatal("Ha ocurrido un error en el parseo de json")
		http.Error(w, "Error en los datos recibidos: -- "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email no puede ser vacio: -- ", 400)
		return
	}

	if len(t.Pass) < 6 {
		http.Error(w, "La longitud de pass es errónea: -- ", 400)
		return
	}
	_, encontrado, _ := bd.CheckUserExist(t.Email)

	if encontrado {
		http.Error(w, "Ya existe el usuario!!!", 400)
		return
	}

	_, status, err := bd.AddUser(t)

	if err != nil {
		http.Error(w, "Ocurrió un error al hacer el registro: --"+err.Error(), 500)
		return
	}
	if status == false {
		http.Error(w, "No se ha conseguido hacer el registro: --", 500)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
