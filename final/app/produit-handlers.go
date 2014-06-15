package app

import (
	"github.com/gorilla/mux"
	"net/http"
	. "strconv"
)

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
	renderHtml(w, p, r)
}

// permet d'ajouter un produit
func ProduitAjouterHandler(w http.ResponseWriter, r *http.Request) {

	// récupération des éléments du formulaire
	titre := r.PostFormValue("titre")
	description := r.PostFormValue("description")
	urlImage := r.PostFormValue("url_image")

	if titre != "" && description != "" && urlImage != "" {
		var p Produit
		p.Titre = titre
		p.Description = description
		p.UrlImage = urlImage
		p.Save()
		http.Redirect(w, r, "/", http.StatusFound)
	}

	// initialisation de l'objet PageProduit
	pp := new(PageProduit)
	pp.Add()

	//insersion dans l'interface Page
	var p Page
	p = pp
	renderHtml(w, p, r)
}

// permet d'editer un produit
func ProduitEditerHandler(w http.ResponseWriter, r *http.Request) {
	// récupération de la variable id
	vars := mux.Vars(r)
	id := vars["id"]

	// récupération des éléments du formulaire
	titre := r.PostFormValue("titre")
	description := r.PostFormValue("description")
	urlImage := r.PostFormValue("url_image")

	if titre != "" && description != "" && urlImage != "" {
		var p Produit
		p.Id, _ = ParseInt(id, 0, 64)
		p.Titre = titre
		p.Description = description
		p.UrlImage = urlImage
		p.Save()
		http.Redirect(w, r, "/produit/"+id+"/", http.StatusFound)
	}

	// initialisation de l'objet PageProduit
	pp := new(PageProduit)
	pp.Produit.Id, _ = ParseInt(id, 0, 64)
	pp.Edit()

	//insersion dans l'interface Page
	var p Page
	p = pp
	renderHtml(w, p, r)
}

// permet de supprimer un produit
func ProduitSupprimerHandler(w http.ResponseWriter, r *http.Request) {
	// récupération de la variable id
	vars := mux.Vars(r)
	id := vars["id"]

	// initialisation de l'objet PageProduit
	pp := new(PageProduit)
	pp.Produit.Id, _ = ParseInt(id, 0, 64)
	pp.Produit.Delete()

	// redirection vers l'accueil
	http.Redirect(w, r, "/", http.StatusFound)
}

// affichage d'une liste de produits
func ProduitListeHandler(w http.ResponseWriter, r *http.Request) {

	// initialisation de l'objet PageProduit
	ppl := new(PageProduitList)
	ppl.View()

	//insersion dans l'interface Page
	var p Page
	p = ppl
	renderHtml(w, p, r)
}
