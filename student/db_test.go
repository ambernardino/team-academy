package student

import (
	"fmt"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func Test_CreateStudent(t *testing.T) {
	testStudent := Student{ID: 1, FirstName: "Maria", LastName: "Joaquina", CursoID: 1, StartDate: time.Now().UTC()}
	db, err := initializeDB()
	if err != nil {
		t.Error("DB is not initialized")
		return
	}

	err = CreateStudent(db, testStudent)
	if err != nil {
		t.Error("Couldn't create a new student")
		return
	}

	testStudent2, err := GetStudentByID(db, testStudent.ID)
	if err != nil {
		t.Error("Couldn't get student by ID")
		return
	}

	if testStudent == testStudent2 {
		return
	}
	t.Errorf("Expected: %v Got: %v", testStudent, testStudent2)
	return

}

func Test_UpdateStudent(t *testing.T) {
	testStudent := Student{ID: 1, FirstName: "Eleutério", LastName: "Azemeís", CursoID: 1, StartDate: time.Now().UTC()}
	db, err := initializeDB()
	if err != nil {
		t.Error("DB is not initialized")
		return
	}

	err = CreateStudent(db, testStudent)
	if err != nil {
		t.Error("Couldn't create a new student")
		return
	}

	testStudent = Student{ID: 1, FirstName: "Eleutério", LastName: "Arnaldo", CursoID: 1, StartDate: time.Now().UTC()}
	err = UpdateStudent(db, testStudent)
	if err != nil {
		t.Error("Couldn't update a student")
		return
	}

	testStudent2, err := GetStudentByID(db, testStudent.ID)
	if err != nil {
		t.Error("Couldn't get student by ID")
		return
	}

	if testStudent == testStudent2 {
		return
	}
	t.Errorf("Expected: %v Got: %v", testStudent, testStudent2)
	return
}

func Test_DeleteStudent(t *testing.T) {
	testStudent := Student{ID: 1, FirstName: "Eleutério", LastName: "Azemeís", CursoID: 1, StartDate: time.Now().UTC()}
	db, err := initializeDB()
	if err != nil {
		t.Error("DB is not initialized")
		return
	}

	err = CreateStudent(db, testStudent)
	if err != nil {
		t.Error("Couldn't create a new student")
		return
	}

	err = DeleteStudent(db, testStudent.ID)
	if err != nil {
		t.Error("Couldn't delete a student")
		return
	}

	testStudent2, err := GetStudentByID(db, testStudent.ID)
	if err != nil {
		return
	}
	t.Errorf("Expected: %v , Got: %v", err, testStudent2)
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
