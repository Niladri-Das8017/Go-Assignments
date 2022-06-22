package controller

import (
	"net/http"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to CRUD API by Niladri Das</h1>"))
}
