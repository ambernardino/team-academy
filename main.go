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

	err = professor.CreateTableIfNotExists(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = professor.CreateProfessor(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	newProfessor := professor.Professor{ID: 10, FirstName: "Mário", LastName: "Ventim", CursoIDs: "MIEEC", CadeiraIDS: "ET", StartDate: time.Now()}
	err = professor.UpdateProfessorInfo(db, newProfessor)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = professor.DeleteProfessor(db, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	professors, err := professor.GetAllProfessors(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(professors)
}

/*func main() {
	database, err := sql.Open("sqlite3", "./clip_holy_grail.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS prof (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT," +
	 								   " cursoIDs TEXT, cadeiraIDs TEXT, startDate DATE)")
	if err != nil {
		fmt.Println(err)
		return
	}

	statement.Exec()
	statement, err = database.Prepare("INSERT INTO prof (firstname, lastname, cursoIDs, cadeiraIDs) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		return
	}

	statement.Exec("Paulo", "Pinto", "MIEEC", "PM")
	rows, err := database.Query("SELECT id, firstname, lastname, cursoIDs, cadeiraIDs FROM prof")
	if err != nil {
		fmt.Println(err)
		return
	}

	var id int
	var firstname, lastname, cursoIDs, cadeiraIDs string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname, &cursoIDs, &cadeiraIDs)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname + " " + cursoIDs + " " + cadeiraIDs)
	}
}*/
