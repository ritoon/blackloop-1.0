package app

import (
	. "strconv"
	"time"
)

type User struct {
	Id        int64
	Nom       string
	Prenom    string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// permet d'enregistrer les éléments du formulaire
func (u User) Save() {
	db.Save(&u)
}

// permet d'enregistrer les modifications des éléments du formulaire
func (u User) Update() {
	db.First(&u, &u.Id).Update(&u)
}

// Delete permet de supprimer un utilisateur
func (u User) Delete() {
	db.Delete(&u)
}

// permet de récupérer un utilisateur
// en fonction de son id
func (u User) GetById() User {
	db.Where("id = ?", Itoa(int(u.Id))).Find(&u)
	return u
}

// permet de récupérer toute la listes des produits
func (u User) GetList() []User {
	var users []User
	db.Find(&users)
	return users
}
