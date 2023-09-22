package main

import (
	"log"

	"github.com/pschlafley/trinityHR/api"
	"github.com/pschlafley/trinityHR/db"
)

func main() {
	store, err := db.NewPostgresStore()

	if err != nil {
		log.Fatal(err)
	}

	// dropDBerr := store.DropTable()

	// if dropDBerr != nil {
	// 	log.Fatal(dropDBerr)
	// }

	if connStr, err := store.Init(); err != nil {
		log.Fatalf("DB_Error: %v\n", err)
	} else {
		log.Printf("%v\n", connStr)
	}

	server := api.NewAPIServer("localhost:3000", store)
	server.Run()
}
