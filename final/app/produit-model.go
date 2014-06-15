package app

import (
	. "strconv"
	"time"
)

type Produit struct {
	Id          int64
	Titre       string
	Description string
	UrlImage    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// permet d'enregistrer les éléments du formulaire
func (p Produit) Save() {
	db.Save(&p)
}

// permet d'enregistrer les modifications des éléments du formulaire
func (p Produit) Update() {
	db.First(&p, &p.Id).Update(&p)
}

// Delete permet de supprimer un produit
func (p Produit) Delete() {
	db.Delete(&p)
}

// permet de récupérer un produit
// en fonction de son id
func (p Produit) GetById() Produit {
	db.Where("id = ?", Itoa(int(p.Id))).Find(&p)
	return p
}

// permet de récupérer toute la listes des produits
func (p Produit) GetList() []Produit {
	var produits []Produit
	db.Find(&produits)
	return produits
}
