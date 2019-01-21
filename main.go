package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Student struct {
	ID        int `gorm:"AUTO_INCREMENT"`
	FirstName string
	LastName  string
	CursoID   int
	StartDate time.Time
}

func CreateTableIfNotExists(db *gorm.DB) (err error) {
	if !db.HasTable(Student{}) {
		return db.CreateTable(Student{}).Error
	}

	return
}

func CreateStudent(db *gorm.DB) (err error) {
	newStudent := Student{FirstName: "Pedro", LastName: "Oliveira", CursoID: 1, StartDate: time.Now()}
	return db.Save(&newStudent).Error
}

func GetAllStudents(db *gorm.DB) (students []Student, err error) {
	err = db.Find(&students).Error
	return
}

func main() {
	db, err := gorm.Open("sqlite3", "clip_holy_grail.db")
	if err != nil {
		return
	}
	db.SingularTable(true)

	err = CreateTableIfNotExists(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = CreateStudent(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	students, err := GetAllStudents(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(students)
}

/*func main() {
	/*database, err := sql.Open("sqlite3", "./clip_holy_grail.db")
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

	statement.Exec("Pedro", "Oliveira", "50544")
	rows, err := database.Query("SELECT id, firstname, lastname, curso_id FROM alunos")
	if err != nil {
		fmt.Println(err)
		return
	}

	var id, curso_id int
	var firstname, lastname string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname, &curso_id)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname + " " + strconv.Itoa(curso_id))
	}
}*/
