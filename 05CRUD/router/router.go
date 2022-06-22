package router

import (
	"github.com/gorilla/mux"
	"github.com/niladridas/crud/controller"
	"github.com/niladridas/crud/database"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", controller.ServeHome).Methods("GET")
	router.HandleFunc("/api/contact", database.InjectDatabase(controller.GetAllContacts)).Methods("GET")
	router.HandleFunc("/api/contact", database.InjectDatabase(controller.CreateContact)).Methods("POST")
	router.HandleFunc("/api/contact/{id}", database.InjectDatabase(controller.UpdateContact)).Methods("PUT")
	router.HandleFunc("/api/contact/{id}", database.InjectDatabase(controller.DeleteContact)).Methods("DELETE")
	router.HandleFunc("/api/contact", database.InjectDatabase(controller.DeleteAllContacts)).Methods("DELETE")

	return router
}
