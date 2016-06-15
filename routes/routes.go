package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//RegisterRoutes registers basic routes for server
func RegisterRoutes(router *mux.Router) {
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
}
