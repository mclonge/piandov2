package handlers

import (
	"log"
	"net/http"
	"os" //Acceso sistema operativo

	"github.com/gorilla/mux" //Manejadores de path de mi API y sus funciones asi como respuestas
	"github.com/mclonge/piando/middlew"
	"github.com/mclonge/piando/routers"
	"github.com/rs/cors" //permisos a API para acceder desde culquier lugar
)

/*Manejadores Contiene todas las rutas y endpoint de mi API*/
func Manejadores() {
	router := mux.NewRouter()
	router.HandleFunc("/registro", middlew.CheckBD(routers.Registro)).Methods("POST")
	//mas endpoints

	PORT := os.Getenv("PORT") //Leer variables de entorno del SO
	if PORT == "" {
		PORT = "8082"
	}
	handler := cors.AllowAll().Handler(router) //Permisos a todos, en el manejador middleware para ver permisos
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
