package main

import (
	"go_app/router"
)

func main() {
	// Konfiguracja routera
	e := router.New()

	// Start serwera
	e.Logger.Fatal(e.Start(":8080"))
}
