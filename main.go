package main

import (
	"fmt"
	"log"

	"github.com/pschlafley/trinityHR/DbTypes"
)

func main() {

	auth := DbTypes.AuthenticationToken{}

	token, err := auth.NewAuthenticationToken()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(token)

	// store, err := db.NewPostgresStore()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// dropDBerr := store.DropTable()

	// if dropDBerr != nil {
	// 	log.Fatal(dropDBerr)
	// // }

	// if connStr, err := store.Init(); err != nil {
	// 	log.Fatalf("DB_Error: %v\n", err)
	// } else {
	// 	log.Printf("%v\n", connStr)
	// }

	// server := api.NewAPIServer(":3000", store)
	// server.Run()
}
