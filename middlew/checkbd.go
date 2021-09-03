package middlew

import (
	"net/http"

	"github.com/mclonge/piando/bd"
)

/*CheckBD Verifico la conexcio antes de hacer nada*/
func CheckBD(next http.HandlerFunc) http.HandlerFunc {
	//un middleware recibe y responde el mismo tipo de dato
	//lo que recibo lo envio

	return func(w http.ResponseWriter, r *http.Request) {

		if bd.CheckConnection() == 0 {
			http.Error(w, "Conexión perdida con BD", 500)
			return
		}

		next.ServeHTTP(w, r)
	}

}
