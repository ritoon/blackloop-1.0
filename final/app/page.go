package app

type Page interface {
	GetTemplate() string
	SetSessionData(u User)
}

type PageWeb struct {
	TitrePage    string
	Template     string
	SessIdUser   int64
	SessNameUser string
	SessIsLogged bool
}

func (p *PageWeb) GetTemplate() string {
	return p.Template
}

func (p *PageWeb) SetSessionData(u User) {
	if u.Id != 0 && u.Prenom != "" {
		p.SessIdUser = u.Id
		p.SessIsLogged = true
		p.SessNameUser = u.Prenom

	} else {
		p.SessIdUser = 0
		p.SessIsLogged = false
		p.SessNameUser = ""
	}
}
