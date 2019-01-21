package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

/*func main() {
	database, err := sql.Open("sqlite3", "./clip_holy_grail.db")
	if err != nil {
 		fmt.Println(err)
		return
	}

	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS prof (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT, cursoIds TEXT, startDate DATE, cadeiraIds TEXT)")
	if err != nil {
		fmt.Println(err)
		return
	}

	statement.Exec()
	statement, err = database.Prepare("INSERT INTO prof (firstname, lastname, cursoIds, cadeiraIds) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		return
	}

	statement.Exec("Paulo", "Montezuma", "MIEEC", "IT")
	rows, err := database.Query("SELECT id, firstname, lastname, cursoIds, cadeiraIds FROM prof")
	if err != nil {
		fmt.Println(err)
		return
	}

	var id int
	var firstname, lastname string
	var cursoIds string
	var cadeiraIds string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname, &cursoIds, &cadeiraIds)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname + " " + cursoIds + " " + cadeiraIds)
	}
}*/

type Professor struct {
	ID         int `gorm:"AUTO_INCREMENT"`
	FirstName  string
	LastName   string
	CursoIds   string
	CadeiraIds string
	StartDate  time.Time
}

func main() {
	db, err := gorm.Open("sqlite3", "clip_holy_grail.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	db.SingularTable(true)
	if !db.HasTable(Professor{}) {
		err = db.CreateTable(Professor{}).Error
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	newProfessor := Professor{FirstName: "Paulo", LastName: "Montezuma", CursoIds: "MIEEC", CadeiraIds: "IT", StartDate: time.Now()}
	err = db.Save(&newProfessor).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	var professors []Professor
	err = db.Find(&professors).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(professors)
}
