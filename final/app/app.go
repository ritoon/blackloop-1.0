package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
)

// gestion des cessions
var store = sessions.NewCookieStore([]byte("something-very-secret"))

// gestion des templates
var templates *template.Template

// initialisation du package
func init() {
	templates = template.Must(template.ParseGlob("./app/vues/*"))
}

// fonction permettant de rendu des pages en fonction du type HTML
func renderHtml(w http.ResponseWriter, p Page, r *http.Request) {

	// creation du header
	w.Header().Add("Accept-Charset", "utf-8")
	w.Header().Add("Content-Type", "text/html")

	// gestion de la session
	session, err := store.Get(r, "cocktails_cookie")

	if err == nil && session.Values["id"] != nil && session.Values["id"] != "0" {
		var u User
		u.Id = session.Values["id"].(int64)
		u.Prenom = session.Values["prenom"].(string)
		p.SetSessionData(u)

	} else {
		// permet de supprimer la session en cours
		session.Options.MaxAge = -1
		var u User
		p.SetSessionData(u)
	}

	// reponse
	templates.ExecuteTemplate(w, p.GetTemplate(), p)
}

// permet de renvoyer un objet générique en JSON
func renderJson(w http.ResponseWriter, msg interface{}) {

	// creation du header
	w.Header().Add("Accept-Charset", "utf-8")
	w.Header().Add("Content-Type", "application/json")

	b, err := json.Marshal(msg)
	if err != nil {
		// envoie un message d'erreur
		fmt.Fprint(w, "error")
	}

	// reponse
	fmt.Fprint(w, string(b))
}

// permet de renvoyer une réponse en TXT
func renderString(w http.ResponseWriter, msg string) {

	// creation du header
	w.Header().Add("Accept-Charset", "utf-8")
	w.Header().Add("Content-Type", "text/plain")

	// reponse
	fmt.Fprint(w, msg)
}
