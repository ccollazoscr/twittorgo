package main

import (
	"log"

	"github.com/ccollazoscr/twittorgo/bd"
	"github.com/ccollazoscr/twittorgo/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
