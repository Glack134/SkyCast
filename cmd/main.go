package main

import (
	"log"

	"github.com/Glack134/ScyCast"
)

func main() {
	srv := new(ScyCast.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("running server %s", err.Error())
	}
}
