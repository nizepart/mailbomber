package main

import (
	"log"

	"github.com/nizepart/mailbomber/internal/app/apiserver"
)

func main() {
	if err := apiserver.Start(); err != nil {
		log.Fatal(err)
	}
}
