package app

import (
	"github.com/gorilla/mux"
	. "strconv"
)

type Page interface {
	SetSessionData(u User) (v bool)
}

type PageWeb struct {
	TitrePage     string
	SessIdUser    int64
	SessNameUser  string
	SessIsLogged  bool
	VarSessServer string
	VarSessCookie string
}

type PageProduitList struct {
	Produits []Produit
	PageWeb
}

type PageProduit struct {
	Produit
	PageWeb
}

// affichage d'un produit
func ProduitHandler(w http.ResponseWriter, r *http.Request) {

	// récupération de la variable id
	vars := mux.Vars(r)
	id := vars["id"]

	// initialisation de l'objet PageProduit
	pp := new(PageProduit)
	pp.Produit.Id, _ = ParseInt(id, 0, 64)
	pp.View()

	//insersion dans l'interface Page
	var p Page
	p = pp
	Render(w, p, r)
}

// fonction permettant de rendu des pages en fonction du type HTML ou TXT
func Render(w http.ResponseWriter, p Page, r *http.Request) {

	w.Header().Add("Accept-Charset", "utf-8")
	w.Header().Add("Content-Type", "text/html")

	session, err := store.Get(r, "cookie")
	templates.ExecuteTemplate(w, Templ, p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
