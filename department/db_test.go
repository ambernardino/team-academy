package department

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func Test_CreateDepartment(t *testing.T) {
	// Given
	db, err := StartDB()
	if err != nil {
		t.Error(err)
		return
	}

	newDepartment := Department{ID: 666, Name: "Test"}

	// Perform
	err = CreateDepartment(db, newDepartment)
	if err != nil {
		t.Error(err)
		return
	}

	fetchedDepartment, err := GetDepartmentByID(db, newDepartment.ID)
	if err != nil {
		t.Error(err)
		return
	}

	// Assert
	if fetchedDepartment != newDepartment {
		t.Errorf("Expected %v, got %v", newDepartment, fetchedDepartment)
		return
	}

	err = DeleteDepartment(db, newDepartment.ID)
	if err != nil {
		t.Error(err)
		return
	}

	return
}

func Test_DeleteDepartment(t *testing.T) {
	// Given
	db, err := StartDB()
	if err != nil {
		t.Error(err)
		return
	}

	newDepartment := Department{ID: 666, Name: "Test"}

	// Perform
	err = CreateDepartment(db, newDepartment)
	if err != nil {
		t.Error(err)
		return
	}

	err = DeleteDepartment(db, newDepartment.ID)
	if err != nil {
		t.Error("Can't delete department ", newDepartment.ID)
		return
	}

	_, err = GetDepartmentByID(db, newDepartment.ID)

	// Assert
	if err == nil {
		t.Error(err)
		return
	}

	return
}

func Test_GetDepartmentByID(t *testing.T) {
	// Given
	db, err := StartDB()
	if err != nil {
		t.Error(err)
		return
	}

	newDepartment := Department{ID: 666, Name: "Test"}

	// Perform
	err = CreateDepartment(db, newDepartment)
	if err != nil {
		t.Error(err)
		return
	}

	fetchedDepartment, err := GetDepartmentByID(db, newDepartment.ID)

	// Assert
	if err != nil {
		t.Error(err)
		return
	}

	if fetchedDepartment != newDepartment {
		t.Errorf("Expected %v, got %v", newDepartment, fetchedDepartment)
		return
	}

	err = DeleteDepartment(db, newDepartment.ID)
	if err != nil {
		t.Error(err)
		return
	}

	return
}

func StartDB() (DB *gorm.DB, err error) {
	DB, err = gorm.Open("sqlite3", "../clip_holy_grail.db")
	if err != nil {
		return
	}
	DB.SingularTable(true)
	return
}
