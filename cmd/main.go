package main

import (
	"log"

	"github.com/pschlafley/trinityHR/api"
	"github.com/pschlafley/trinityHR/db"
	"github.com/pschlafley/trinityHR/types"
)

func main() {
	store, err := db.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	logger := types.InitializeLogger()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	// dropDBerr := store.DropTable()
	//
	// if dropDBerr != nil {
	// 	log.Fatal(dropDBerr)
	// }

	if connStr, err := store.Init(); err != nil {
		log.Fatalf("DB_Error: %v\n", err)
	} else {
		logger.Info(connStr)
	}

	server := api.NewAPIServer("localhost:3000", store)
	server.Run(logger)
}
