package app

import (
	"net/http"
)

// affichage la page de login et connecte
func UserConnexionHandler(w http.ResponseWriter, r *http.Request) {

	// initialisation des variables de user
	var u User
	u.Email = r.PostFormValue("email")
	u.Password = r.PostFormValue("password")

	// initialisation de la page de connexion
	pu := new(PageUser)

	// vérifie que l'email et le password ne sont pas vide
	if u.Email == "" && u.Password == "" {
		pu.View()
	} else {
		/*
				session, _ := store.Get(r, "cocktails_cookie")
			    // Set some session values.
			    session.Values["id"] = 1
			    session.Values["prenom"] = "Foo"
			    // Save it.
			    session.Save(r, w)
		*/
		pu.Connexion()
	}

	//insersion dans l'interface Page
	var p Page
	p = pu
	renderHtml(w, p, r)
}

// affichage la page home et déconnecte
func UserDeconnexionHandler(w http.ResponseWriter, r *http.Request) {

	/*
			session, _ := store.Get(r, "cocktails_cookie")
		    // Set some session values.
		    session.Values["id"] = 0
		    // Save it.
		    session.Save(r, w)
	*/
	// initialisation de la page de connexion
	pu := new(PageUser)

	//insersion dans l'interface Page
	var p Page
	p = pu
	renderHtml(w, p, r)
}
