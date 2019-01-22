package student

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"testing"
	"time"
)

func Test_CreateStudent(t *testing.T) {
	db, err := initializeDB()
	if err != nil {
		t.Error(err)
		return
	}

	newStudent := Student{ID: 4, FirstName: "Pedro", LastName: "Oliveira", CursoID: 1, StartDate: time.Now().UTC()}
	err = CreateStudent(db, newStudent)
	if err != nil {
		t.Error(err)
		return
	}

	foundStudent, err := GetStudentByID(db, newStudent.ID)
	if err != nil {
		t.Error(err)
		return
	}

	if foundStudent != newStudent {
		t.Errorf("Wanted: %+v, got: %+v", newStudent, foundStudent)
		return
	}
}

func initializeDB() (DB *gorm.DB, err error) {
	DB, err = gorm.Open("sqlite3", "../clip_holy_grail.db")
	if err != nil {
		return
	}

	DB.SingularTable(true)
	return
}
