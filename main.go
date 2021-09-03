package main

import (
	"log"

	"github.com/mclonge/piando/bd"
	"github.com/mclonge/piando/handlers"
)

func main() {
	log.Println("BRUMBRUM")

	if bd.CheckConnection() == 0 {
		log.Println("Sin conexión a BD")
		return
	}
	handlers.Manejadores()
}
