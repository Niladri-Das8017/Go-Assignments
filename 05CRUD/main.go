package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/niladridas/crud/database"
	"github.com/niladridas/crud/router"
)

func main() {

	fmt.Println("\nWelcome to CRUD")

	//Creatiung Database Connection
	err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	//Routing
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4040", r))
}
