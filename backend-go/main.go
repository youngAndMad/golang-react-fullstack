package main

import (
	"log"
)

func main() {

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	server := newAPIServer(":8080", store)
	server.Run()
	//go run storago.go main.go api.go types.go
}
