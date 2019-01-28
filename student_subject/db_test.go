package student_subject

import (
	"fmt"
	"team-academy/student"
	"team-academy/subject"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func Test_AddStudentToSubject(t *testing.T) {
	testStudent := student.Student{ID: 7, FirstName: "Eleutério", LastName: "Azemeís", CursoID: 1, StartDate: time.Now().UTC()}
	testSubject := subject.Subject{ID: 9, Name: "Análise Matemática 4", Description: "Easy"}

	db, err := initializeDB()
	if err != nil {
		t.Error("DB is not initialized")
		return
	}

	err = student.CreateStudent(db, testStudent)
	if err != nil {
		t.Error("Couldn't create a new student")
		return
	}

	err = subject.CreateSubject(db, testSubject)
	if err != nil {
		t.Error("Couldn't create a new subject")
		return
	}

	err = AddStudentToSubject(db, testStudent.ID, testSubject.ID)
	if err != nil {
		t.Error("Couldn't add the student to the subject")
		return
	}

	err = RemoveStudentFromSubject(db, testStudent.ID, testSubject.ID)
	if err != nil {
		t.Error("Couldn't remove the student from the subject")
		return
	}
	return
}

func Test_AddStudentToNonExistantSubject(t *testing.T) {
	testStudent := student.Student{ID: 51, FirstName: "Amílcar", LastName: "Alho", CursoID: 1, StartDate: time.Now().UTC()}
	testSubject := subject.Subject{ID: 21, Name: "Introdução às Telecomunicações", Description: "Stupid"}

	db, err := initializeDB()
	if err != nil {
		t.Error("DB is not initialized")
		return
	}

	err = student.CreateStudent(db, testStudent)
	if err != nil {
		t.Error("Couldn't create a new student")
		return
	}

	err = AddStudentToSubject(db, testStudent.ID, testSubject.ID)
	if err != nil {
		return
	}
	t.Errorf("Expected err != nil, Got %v", err)
	return
}

func Test_AddNonExistantStudentToSubject(t *testing.T) {
	testStudent := student.Student{ID: 7, FirstName: "Cristiano", LastName: "Ronaldo", CursoID: 1, StartDate: time.Now().UTC()}
	testSubject := subject.Subject{ID: 24, Name: "Sistemas de Telecomunicações", Description: "Easy"}

	db, err := initializeDB()
	if err != nil {
		t.Error("DB is not initialized")
		return
	}

	err = subject.CreateSubject(db, testSubject)
	if err != nil {
		t.Error("Couldn't create a new subject")
		return
	}

	err = AddStudentToSubject(db, testStudent.ID, testSubject.ID)
	if err != nil {
		return
	}
	t.Errorf("Expected err != nil, Got %v", err)
	return
}

func Test_AddRegistedStudentToSubject(t *testing.T) {
	testStudent := student.Student{ID: 57, FirstName: "Maria", LastName: "Manel", CursoID: 1, StartDate: time.Now().UTC()}
	testSubject := subject.Subject{ID: 14, Name: "Geometria", Description: "Easy"}

	db, err := initializeDB()
	if err != nil {
		t.Error("DB is not initialized")
		return
	}

	err = student.CreateStudent(db, testStudent)
	if err != nil {
		t.Error("Couldn't create a new student")
		return
	}

	err = subject.CreateSubject(db, testSubject)
	if err != nil {
		t.Error("Couldn't create a new subject")
		return
	}

	err = AddStudentToSubject(db, testStudent.ID, testSubject.ID)
	if err != nil {
		t.Error("Couldn't add the student to the subject")
		return
	}

	err = AddStudentToSubject(db, testStudent.ID, testSubject.ID)
	if err == nil {
		t.Error("Could add the student again to the subject")
		return
	}

	err = RemoveStudentFromSubject(db, testStudent.ID, testSubject.ID)
	if err != nil {
		t.Error("Couldn't remove the student from the subject")
		return
	}
	return
}

func Test_GetSubjectAndInfoByStudentID(t *testing.T) {
	testStudent := student.Student{ID: 18, FirstName: "Zézinho", LastName: "Manel", CursoID: 4, StartDate: time.Now().UTC()}
	testSubject := subject.Subject{ID: 20, Name: "Cálculo Numérico", Description: "Easy"}

	db, err := initializeDB()
	if err != nil {
		t.Error("DB is not initialized")
		return
	}

	err = student.CreateStudent(db, testStudent)
	if err != nil {
		t.Error("Couldn't create a new student")
		return
	}

	err = subject.CreateSubject(db, testSubject)
	if err != nil {
		t.Error("Couldn't create a new subject")
		return
	}

	err = AddStudentToSubject(db, testStudent.ID, testSubject.ID)
	if err != nil {
		t.Error("Couldn't add the student to the subject")
		return
	}

	_, err = GetSubjectAndInfoByStudentID(db, testStudent.ID)
	if err != nil {
		t.Error("Couldn't get subject and info of given student")
		return
	}

	err = RemoveStudentFromSubject(db, testStudent.ID, testSubject.ID)
	if err != nil {
		t.Error("Couldn't remove the student from the subject")
		return
	}
	return
}

func initializeDB() (DB *gorm.DB, err error) {
	DB, err = gorm.Open("sqlite3", "../clip_holy_grail.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	DB.SingularTable(true)
	return
}
