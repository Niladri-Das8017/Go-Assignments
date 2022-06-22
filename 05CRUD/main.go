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
	database.InitDB()

	//Routing
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4040", r))
}
