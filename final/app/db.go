package app

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// fonction permettant de se connecter à la base de donnée
func connectToDatabase() gorm.DB {
	db, _ := gorm.Open("mysql", "root:root@tcp(127.0.0.1:8889)/blackloop?charset=utf8&parseTime=True")
	db.SingularTable(true)
	return db
}

// connexion à la base de donnée
var db = connectToDatabase()
