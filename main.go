package main

import (
	"fmt"
	"team-academy/grade"
	"team-academy/professor"
	"team-academy/student"
	"team-academy/student_subject"
	"team-academy/subject"
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
	err = populateDatabase(db)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func populateDatabase(db *gorm.DB) (err error) {
	err = professor.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	err = student.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	err = subject.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	err = student_subject.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	err = grade.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	newSubject := subject.Subject{ID: 1, Name: "Cadeira 1", Description: "Nothing"}
	err = subject.CreateSubject(db, newSubject)
	if err != nil {
		return
	}

	newSubject = subject.Subject{ID: 2, Name: "Cadeira 2", Description: "Nothing"}
	err = subject.CreateSubject(db, newSubject)
	if err != nil {
		return
	}

	newSubject = subject.Subject{ID: 3, Name: "Cadeira 3", Description: "Nothing"}
	err = subject.CreateSubject(db, newSubject)
	if err != nil {
		return
	}

	newProfessor := professor.Professor{ID: 1, FirstName: "Prof 1", LastName: "Prof 1", CursoIDs: "Curso 1", CadeiraIDS: "Cadeira 1", StartDate: time.Now().UTC()}
	err = professor.CreateProfessor(db, newProfessor)
	if err != nil {
		return
	}

	newProfessor = professor.Professor{ID: 2, FirstName: "Prof 2", LastName: "Prof 2", CursoIDs: "Curso 2", CadeiraIDS: "Cadeira 2", StartDate: time.Now().UTC()}
	err = professor.CreateProfessor(db, newProfessor)
	if err != nil {
		return
	}

	newProfessor = professor.Professor{ID: 3, FirstName: "Prof 3", LastName: "Prof 3", CursoIDs: "Curso 3", CadeiraIDS: "Cadeira 3", StartDate: time.Now().UTC()}
	err = professor.CreateProfessor(db, newProfessor)
	if err != nil {
		return
	}

	newStudent := student.Student{ID: 1, FirstName: "Student 1", LastName: "Student 1", CursoID: 1, StartDate: time.Now().UTC()}
	err = student.CreateStudent(db, newStudent)
	if err != nil {
		return
	}

	newStudent = student.Student{ID: 2, FirstName: "Student 2", LastName: "Student 2", CursoID: 2, StartDate: time.Now().UTC()}
	err = student.CreateStudent(db, newStudent)
	if err != nil {
		return
	}

	newStudent = student.Student{ID: 3, FirstName: "Student 3", LastName: "Student 3", CursoID: 3, StartDate: time.Now().UTC()}
	err = student.CreateStudent(db, newStudent)
	if err != nil {
		return
	}

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			err = student_subject.AddStudentToSubject(db, i, j)
			if err != nil {
				return
			}
		}
	}

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			newGrade := grade.Grade{SubjectID: i, StudentID: j, Rank: "Failed"}
			err = grade.GiveGrade(db, newGrade)
			if err != nil {
				return
			}
		}
	}

	return
}
