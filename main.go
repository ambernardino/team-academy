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
	db.SingularTable(true)

	err = student.CreateTableIfNotExists(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	s := student.Student{ID: 1, FirstName: "Ricardo", LastName: "Cenas", CursoID: 1, StartDate: time.Now()}
	err = student.CreateStudent(db, s)
	if err != nil {
		fmt.Println(err)
		return
	}
	s = student.Student{ID: 1, FirstName: "Paulo"}
	err = student.UpdateStudent(db, s)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = student.DeleteStudent(db, 3)
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

	//------------------------------------------------------------------------------------
	err = student_subject.CreateTableIfNotExists(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = student_subject.AddStudentToSubject(db, 1, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = student_subject.RemoveStudentFromSubject(db, 1, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	studentSubject, err := student_subject.GetSubjectsByStudentID(db, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	studentSubject, err = student_subject.GetStudentsBySubjectID(db, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(studentSubject)
	//-------------------------------------------------------------------------------------------------------------
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
	//err = professor.RemoveProfessor(db, 1)
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

	/*-------------------------------------------------SUBJECT TESTS----------------------------------------------------*/
	err = subject.CreateTableIfNotExist(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	subj := subject.Subject{Name: "IT", Description: "Amazing"}
	err = subject.CreateSubject(db, subj)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = subject.RemoveSubject(db, 8)
	if err != nil {
		fmt.Println(err)
		return
	}
	subje := subject.Subject{ID: 2, Name: "AED", Description: "yay"}
	err = subject.UpdateSubject(db, subje)
	if err != nil {
		fmt.Println(err)
		return
	}
	subs, err := subject.GetAllSubjects(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(subs)
}
