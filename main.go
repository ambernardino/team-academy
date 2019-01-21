package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Alunos struct {
	ID        int `gorm:"AUTO_INCREMENT"`
	FirstName string
	LastName  string
	DegreeID  int
	StartDate time.Time
}

func main() {
	db, err := gorm.Open("sqlite3", "clip_holy_grail.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	db.SingularTable(true)
	if !db.HasTable(Alunos{}) {
		err = db.CreateTable(Alunos{}).Error
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	newAluno := Alunos{FirstName: "John", LastName: "Doe", DegreeID: 1, StartDate: time.Now()}
	err = db.Save(&newAluno).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	var alunos []Alunos
	err = db.Find(&alunos).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(alunos)
}

/*func main() {
	database, err := sql.Open("sqlite3", "./clip_holy_grail.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS alunos (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT, curso_id NUMBER, start_date DATE)")
	if err != nil {
		fmt.Println(err)
		return
	}

	statement.Exec()
	statement, err = database.Prepare("INSERT INTO alunos (firstname, lastname, curso_id) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		return
	}

	statement.Exec("Francisco", "Peres", 50034)
	rows, err := database.Query("SELECT id, firstname, lastname, curso_id FROM alunos")
	if err != nil {
		fmt.Println(err)
		return
	}

	var id, cursoID int
	var firstname, lastname string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname, &cursoID)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname + " " + strconv.Itoa(cursoID))
	}
}*/
