package main

import (
	"fmt"
	"team-academy/professor"
	"team-academy/student"
	"team-academy/student_subject"
	"team-academy/subject"

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

	updatedSubject := subject.Subject{ID: 2, Name: "Eletrónica 2", Description: "Outra seca desgraçada"}
	err = subject.UpdateSubjectInfo(db, updatedSubject)
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

	studentsSubjects, err := student_subject.GetSubjectsByStudentID(db, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(studentsSubjects)

	studentsSubjects, err = student_subject.GetSubjectsByStudentID(db, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(studentsSubjects)
}
