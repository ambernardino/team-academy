package main

import (
	"fmt"
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

	err = professor.CreateTableIfNotExists(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = student.CreateTableIfNotExists(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = subject.CreateTableIfNotExists(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = student_subject.CreateTableIfNotExists(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = professor.CreateProfessor(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	newStudent := student.Student{FirstName: "Pedro", LastName: "Oliveira", CursoID: 1, StartDate: time.Now()}
	err = student.CreateStudent(db, newStudent)
	if err != nil {
		fmt.Println(err)
		return
	}

	newSubject := subject.Subject{ID: 2, Name: "Eletrónica 1", Description: "Uma seca desgraçada"}
	err = subject.CreateSubject(db, newSubject)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = student_subject.AddStudentToSubject(db, 1, 1)
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

	s := student.Student{ID: 1, FirstName: "Ricardo", LastName: "Cenas", CursoID: 1, StartDate: time.Now()}
	err = student.UpdateStudent(db, s)
	if err != nil {
		fmt.Println(err)
		return
	}

	updatedSubject := subject.Subject{ID: 2, Name: "Eletrónica 2", Description: "Outra seca desgraçada"}
	err = subject.UpdateSubjectInfo(db, updatedSubject)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = professor.DeleteProfessor(db, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = student.DeleteStudent(db, 3)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = subject.DeleteSubject(db, 8)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = student_subject.RemoveStudentFromSubject(db, 1, 1)
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

	students, err := student.GetAllStudents(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(students)

	subjects, err := subject.GetAllSubjects(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(subjects)

	studentSubject, err := student_subject.GetSubjectsByStudentID(db, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(studentSubject)

	studentSubject, err = student_subject.GetSubjectsByStudentID(db, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(studentSubject)
}
