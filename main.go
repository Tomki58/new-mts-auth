package main

import (
	"log"
	"mts/auth/httpserver"
)

func main() {
	serivce, err := httpserver.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := serivce.Server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
