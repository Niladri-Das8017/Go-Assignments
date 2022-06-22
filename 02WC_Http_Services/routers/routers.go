package routers

import (
	"httpservices/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/wordcount", controller.WordCount).Methods("POST")

	return router

}
