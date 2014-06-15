package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello Golang Paris")
}

func main() {

	// creation de l'objet mux
	r := mux.NewRouter()

	r.HandleFunc("/", HelloHandler)

	// injection de mux dans l'objet http
	http.Handle("/", r) // pattern string, Handler func
	log.Println("Listening...")

	// écoute sur le port 3000
	http.ListenAndServe(":3000", nil)
}
