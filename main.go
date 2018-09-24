package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mynotes/controller"
	"github.com/mynotes/middleware"
)

func main() {
	router := mux.NewRouter()
	adapter := middleware.Adapt(router, middleware.Cors, middleware.WithDB)

	router.Handle("/notes", controller.CreateNote()).Methods("POST")
	router.Handle("/notes", controller.ReadNote()).Methods("GET")
	router.Handle("/notes/{id}", controller.UpdateNote()).Methods("PUT")
	router.Handle("/notes/{id}", controller.FindNote()).Methods("GET")
	router.Handle("/notes/{id}", controller.DeleteNote()).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", adapter))
}
