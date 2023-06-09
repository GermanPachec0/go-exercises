package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/germanpachec0/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("LocalHost:8080", r))
	fmt.Printf("Starting server at port 8080\n")
}
