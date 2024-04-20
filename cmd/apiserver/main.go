package main

import (
	"log"

	"github.com/nizepart/rest-go/internal/app/apiserver"
)

func main() {
	if err := apiserver.Start(); err != nil {
		log.Fatal(err)
	}
}
