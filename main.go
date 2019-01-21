package main

import (
	"fmt"
	"team-academy/student"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := gorm.Open("sqlite3", "clip_holy_grail.db")
	if err != nil {
		return
	}
	db.SingularTable(true)

	err = student.CreateTableIfNotExists(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = student.CreateStudent(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	s1 := student.Student{ID: 1, FirstName: "Ricardo", LastName: "Cenas", CursoID: 1, StartDate: time.Now()}
	err = student.UpdateStudent(db, s1)
	if err != nil {
		fmt.Println(err)
		return
	}

	s2 := student.Student{ID: 1}
	err = student.DeleteStudent(db, s2)
	if err != nil {
		fmt.Println(err)
		return
	}

	students, err := student.GetAllStudents(db)
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
