package main

import (
	"log"
)

func main() {
	store, err := NewPostgresStore()

	if err != nil {
		log.Fatal(err)
	}

	if connStr, err := store.Init(); err != nil {
		log.Fatalf("DB_Error: %v\n", err)
	} else {
		log.Printf("%v\n", connStr)
	}

	server := NewAPIServer(":3000", store)
	server.Run()
}