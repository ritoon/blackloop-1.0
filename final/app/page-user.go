package app

type PageUser struct {
	User
	PageWeb
}

// permet d'afficher la page connexion
func (pu *PageUser) View() {
	pu.Template = "connexion"
	pu.TitrePage = "Connexion"
}

// permet d'afficher la page connexion
func (pu *PageUser) Connexion() {
	pu.Template = "connexion"
	pu.TitrePage = "Connexion"
}

// permet de déconnecter
// et permet d'afficher la page d'accueil
func (pu *PageUser) Deconnexion() {
	pu.Template = "deconnexion"
}
