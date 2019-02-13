package subject

import (
	"fmt"
	"team-academy/component"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func Test_CreateSubject(t *testing.T) {
	// Given
	subject := Subject{ID: 666, Name: "Test", Description: "Test"}
	db, err := StartDB()
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

	err = DeleteSubject(db, subject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	return
}

func Test_CreateDuplicateSubject(t *testing.T) {
	// Given
	db, err := StartDB()
	if err != nil {
		t.Error(err)
		return
	}

	subject := Subject{ID: 666, Name: "Test", Description: "Test"}
	err = CreateSubject(db, subject)
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	err = CreateSubject(db, subject)
	if err != component.ErrSubjectAlreadyExists {
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

	err = DeleteSubject(db, subject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	return
}

func Test_UpdateSubject(t *testing.T) {
	// Given
	db, err := StartDB()
	if err != nil {
		t.Error(err)
		return
	}

	newSubject := Subject{ID: 666, Name: "Test", Description: "Test"}
	err = CreateSubject(db, newSubject)
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	updatedSubject := Subject{ID: 666, Name: "Updated", Description: "Updated"}
	err = UpdateSubjectInfo(db, updatedSubject)
	if err != nil {
		t.Error(err)
		return
	}

	testSubject, err := GetSubjectByID(db, updatedSubject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	// Assert
	if updatedSubject != testSubject {
		t.Error("The subject wasn't updated ")
		return
	}

	err = DeleteSubject(db, newSubject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	return
}

func Test_DeleteSubject(t *testing.T) {
	// Given
	db, err := StartDB()
	if err != nil {
		t.Error(err)
		return
	}

	newSubject := Subject{ID: 666, Name: "Test", Description: "Test"}
	err = CreateSubject(db, newSubject)
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	err = DeleteSubject(db, newSubject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = GetSubjectByID(db, newSubject.ID)

	// Assert
	if err == nil {
		t.Error("Subject wasn't properly deleted.")
		return
	}

	return
}

func Test_GetSubjectByID(t *testing.T) {
	// Given
	newSubject := Subject{ID: 666, Name: "Test", Description: "Test"}
	db, err := StartDB()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Perform
	err = CreateSubject(db, newSubject)
	if err != nil {
		t.Error(err)
		return
	}

	testSubject, err := GetSubjectByID(db, newSubject.ID)

	// Assert
	if newSubject != testSubject {
		t.Error("The fetched subject is different from the original subject")
		return
	}

	err = DeleteSubject(db, testSubject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	return
}

func Test_GetSubjectByName(t *testing.T) {
	// Given
	newSubject := Subject{ID: 666, Name: "Test", Description: "Test"}
	db, err := StartDB()
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	err = CreateSubject(db, newSubject)
	if err != nil {
		t.Error(err)
		return
	}

	testSubject, err := GetSubjectByName(db, newSubject.Name)
	if err != nil {
		t.Error("Can't gather subject by name.")
		return
	}

	// Assert
	if newSubject != testSubject {
		t.Error("The fetched subject is different from the original subject")
		return
	}

	err = DeleteSubject(db, testSubject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	return
}

func StartDB() (db *gorm.DB, err error) {
	db, err = gorm.Open("sqlite3", "../clip_holy_grail.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	db.SingularTable(true)
	return
}
