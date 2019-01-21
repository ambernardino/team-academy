package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Professor struct {
	ID         int `gorm:"AUTO_INCREMENT"`
	FirstName  string
	LastName   string
	CursoIDs   string
	CadeiraIDS string
	StartDate  time.Time
}

func main() {
	db, err := gorm.Open("sqlite3", "clip_holy_grail.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = CreateTableIfNotExists(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = CreateProfessor(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	professors, err := GetAllProfessors(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(professors)
}

func CreateTableIfNotExists(db *gorm.DB) (err error) {
	db.SingularTable(true)
	if !db.HasTable(Professor{}) {
		return db.CreateTable(Professor{}).Error
	}
	return
}

func CreateProfessor(db *gorm.DB) (err error) {
	newProfessor := Professor{FirstName: "Paulo", LastName: "Pinto", CursoIDs: "MIEEC", CadeiraIDS: "PM", StartDate: time.Now()}
	return db.Save(&newProfessor).Error
}

func GetAllProfessors(db *gorm.DB) (professors []Professor, err error) {
	err = db.Find(&professors).Error
	return
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
