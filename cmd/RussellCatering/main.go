package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/dixonwille/RussellCatering/adapters"
	"github.com/dixonwille/RussellCatering/env"

	"github.com/dixonwille/RussellCatering/routes"

	"github.com/gorilla/mux"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()
	router := mux.NewRouter().StrictSlash(true)
	routes.RegisterRoutes(router)

	//Used to specify adapters at root level
	middleRouter := http.NewServeMux()
	//Apply global adapters here
	middleRouter.Handle("/", adapters.Adapt(router,
		adapters.Logging(env.Logger),
	))
	//Apply API only adapters
	middleRouter.Handle("/api/", adapters.Adapt(router,
		adapters.Logging(env.Logger),
		adapters.Header("Content-Type", "application/json"),
	))
	port := ":" + strconv.Itoa(env.Port)
	env.Logger.Println("Listening on port " + port + " for http")
	if err := http.ListenAndServe(port, middleRouter); err != nil {
		panic(err.Error())
	}
}
