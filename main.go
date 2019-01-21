package main

import (
	"fmt"
	"team-academy/student"

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
	err = student.CreateTable(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = student.CreateStudent(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	students, err := student.GetStudents(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(students)

	//student.UpdateStudent(db, student.Student{ID: 1, FirstName: "Teste", LastName: "Teste", DegreeID: 20, StartDate: time.Now()})
	err = student.DeleteStudent(db, 3)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()
	fmt.Println("------------------------")
	fmt.Println()

	students, err = student.GetStudents(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(students)
}
