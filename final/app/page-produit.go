package app

type PageProduit struct {
	Produit
	PageWeb
}

// fonction public
// permet d'afficher un produit
func (p *PageProduit) View() {
	p.Template = "produit"
	p.Produit = p.Produit.GetById()
	p.TitrePage = "Cocktail : " + p.Produit.Titre
}

// fonction public
// permet d'ajouter un produit
func (p *PageProduit) Add() {
	p.Template = "produit_formulaire"
	p.Produit = p.Produit.GetById()
	p.TitrePage = "Edition cocktail : " + p.Produit.Titre
}

// fonction public
// permet d'editer un produit
func (p *PageProduit) Edit() {
	p.Template = "produit_formulaire"
	p.Produit = p.Produit.GetById()
	p.TitrePage = "Edition cocktail : " + p.Produit.Titre
}

type PageProduitList struct {
	Produits []Produit
	PageWeb
}

// fonction public
// permet d'afficher une liste de produits
func (p *PageProduitList) View() {

	var produit Produit

	p.Template = "produit_liste"
	p.Produits = produit.GetList()
	p.TitrePage = "Liste des produits"
}
