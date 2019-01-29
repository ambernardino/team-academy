package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type App struct {
	Db *gorm.DB
}

var Application App

func StartDatabase() (err error) {
	db, err := gorm.Open("sqlite3", "clip_holy_grail.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	Application.Db = db
	db.SingularTable(true)
	return
}
