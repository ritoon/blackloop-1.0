package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"time"
)

// gestion des templates
var templates *template.Template

// initialisation du package
func init() {
	templates = template.Must(template.ParseGlob("./vues/*"))
}

func main() {

	// creation de l'objet mux
	r := mux.NewRouter()

	// listes des routes et Handlers
	r.HandleFunc("/", ProduitListeHandler)

	//gestion des fichiers statiques (JS, CSS, Images)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	// injection de mux dans l'objet http
	http.Handle("/", r) // pattern string, Handler func
	log.Println("Listening...")

	// écoute sur le port 3000
	http.ListenAndServe(":3000", nil)
}

// création de l'objet produit
type Produit struct {
	Id          int64
	Titre       string
	Description string
	UrlImage    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// création de l'objet page list produit
type PageProduitList struct {
	TitrePage string
	Produits  []Produit
}

// handler pour la lsite des produits
func ProduitListeHandler(w http.ResponseWriter, r *http.Request) {

	var ppl PageProduitList

	ppl.TitrePage = "Ma liste de produits"
	ppl.Produits = make([]Produit, 4)

	var p1 Produit
	p1.Titre = "Mojito"
	p1.UrlImage = "/img/1.jpg"
	ppl.Produits[0] = p1

	var p2 Produit
	p2.Titre = "Piña Colada"
	p2.UrlImage = "/img/2.jpg"
	ppl.Produits[1] = p2

	var p3 Produit
	p3.Titre = "Margarita"
	p3.UrlImage = "/img/3.jpg"
	ppl.Produits[2] = p3

	var p4 Produit
	p4.Titre = "Dirty Martini"
	p4.UrlImage = "/img/4.jpg"
	ppl.Produits[3] = p4

	// creation du header
	w.Header().Add("Accept-Charset", "utf-8")
	w.Header().Add("Content-Type", "text/html")

	// reponse
	templates.ExecuteTemplate(w, "produit_liste", ppl)
}
