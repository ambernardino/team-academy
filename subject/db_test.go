package subject

import (
	"team-academy/component"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func Test_CreateSubject(t *testing.T) {
	// Given
	subject := Subject{ID: 1, Name: "PIIC", Description: "StartDB"}
	db, err := initializeDB()
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	err = CreateSubject(db, subject)
	if err != nil {
		t.Error(err)
		return
	}
	testSubject, err := GetSubjectByID(db, subject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	// Assert
	if subject != testSubject {
		t.Errorf("Expected: %v Got: %v.", subject, testSubject)
		return
	}

	err = DeleteSubject(db, 1)
	if err != nil {
		t.Error(err)
		return
	}

	return
}

func Test_CreateDuplicateSubject(t *testing.T) {
	// Given
	subject := Subject{ID: 58, Name: "PIIC", Description: "StartDB"}
	db, err := initializeDB()
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	err = CreateSubject(db, subject)
	if err != nil {
		t.Error(err)
		return
	}

	err = CreateSubject(db, subject)
	if err != component.ErrSomethingAlreadyExists {
		t.Error(err)
		return
	}

	err = nil
	testSubject, err := GetSubjectByID(db, subject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	// Assert
	if subject != testSubject {
		t.Errorf("Expected: %v Got: %v.", subject, testSubject)
		return
	}

	err = DeleteSubject(db, 58)
	if err != nil {
		t.Error(err)
		return
	}
	return
}

func Test_DeleteSubject(t *testing.T) {

}

func initializeDB() (DB *gorm.DB, err error) {
	DB, err = gorm.Open("sqlite3", "../clip_holy_grail.db")
	if err != nil {
		return
	}
	DB.SingularTable(true)
	return
}
