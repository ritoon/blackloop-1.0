package app

import (
	"time"
)

type Produit struct {
	Id          int64
	Titre       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// permet d'enregistrer les éléments du formulaire
func (p Produit) Save() {
	db.Save(&p)
}

// permet d'enregistrer les modifications des éléments du formulaire
func (f Forum) Update() {
	db.First(&f, &f.Id).Update(&f)
}

// permet de récupérer toute la listes des produits
func (p Produit) getList() []Produit {
	var produits []Produit
	db.Find(&produits)
	return produits
}

// Delete permet de supprimer un produit
func (p Produit) Delete() {
	db.Delete(&p)
}
