package main

import (
	"fmt"
	"team-academy/professor"
	"time"

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
	err = professor.CreateTableIfNotExist(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = professor.CreateProfessors(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	prof := professor.Professor{ID: 6, FirstName: "Sérgio", LastName: "Onofre", CursoIds: "MIEEC", CadeiraIds: "IT", StartDate: time.Now()}
	err = professor.UpdateProfessor(db, prof)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = professor.RemoveProfessor(db, 8)
	if err != nil {
		return
	}
	profs, err := professor.GetAllProfessors(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(profs)
}

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
