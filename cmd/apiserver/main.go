package main

import (
	"github.com/nizepart/rest-go/internal/app/apiserver"
	"log"
)

func main() {
	if err := apiserver.Start(); err != nil {
		log.Fatal(err)
	}
}
