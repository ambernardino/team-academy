package component

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Application struct {
	DB *gorm.DB
}

var App Application

func StartDB() (err error) {
	db, err := gorm.Open("sqlite3", "clip_holy_grail.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	db.SingularTable(true)
	//err = populateDatabase(db)
	App = Application{DB: db}

	return
}
