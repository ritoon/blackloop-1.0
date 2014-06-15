package main

import (
	"github.com/gorilla/mux"
	. "github.com/ritoon/blackloop/final/app"
	"log"
	"net/http"
)

func main() {

	// creation de l'objet mux
	r := mux.NewRouter()

	// listes des routes et Handlers
	r.HandleFunc("/", ProduitListeHandler)
	r.HandleFunc("/produit/{id:[0-9]+}/", ProduitHandler)
	// CRUD
	r.HandleFunc("/produit/ajouter/", ProduitAjouterHandler)
	r.HandleFunc("/produit/editer/{id:[0-9]+}/", ProduitEditerHandler)
	r.HandleFunc("/produit/supprimer/{id:[0-9]+}/", ProduitSupprimerHandler)

	// gestion de la connexion
	r.HandleFunc("/connexion/", UserConnexionHandler)
	r.HandleFunc("/deconnexion/", UserDeconnexionHandler)

	//gestion des fichiers statiques (JS, CSS, Images)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	// injection de mux dans l'objet http
	http.Handle("/", r) // pattern string, Handler func
	log.Println("Listening...")

	// écoute sur le port 3000
	http.ListenAndServe(":3000", nil)
}
