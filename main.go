package main

import (
	"fmt"
	"team-academy/student_subject"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := gorm.Open("sqlite3", "clip_holy_grail.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	db.SingularTable(true)
	err = student_subject.CreateTable(db)
	if err != nil {
		fmt.Println(err)
		return
	}
}
