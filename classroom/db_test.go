package classroom

import (
	"fmt"
	"team-academy/department"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func Test_CreateClassroom(t *testing.T) {
	// Given
	db, err := StartDB()
	if err != nil {
		t.Error(err)
		return
	}

	newDepartment := department.Department{ID: 666, Name: "Test"}
	err = department.CreateDepartment(db, newDepartment)
	if err != nil {
		t.Error(err)
		return
	}

	newClassroom := Classroom{ID: 666, Name: "Test", DepartmentID: newDepartment.ID}

	// Perform
	err = CreateClassroom(db, newClassroom)
	if err != nil {
		t.Error(err)
		return
	}

	fetchedClassroom, err := GetClassroomByID(db, newClassroom.ID)

	// Assert
	if err != nil {
		t.Error(err)
		return
	}

	if fetchedClassroom != newClassroom {
		t.Errorf("Expected %v, got %v", newClassroom, fetchedClassroom)
		return
	}

	err = DeleteClassroom(db, newClassroom.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = department.DeleteDepartment(db, newDepartment.ID)
	if err != nil {
		t.Error(err)
		return
	}

	return
}

func Test_CreateClassroomOnNonExistantDepartment(t *testing.T) {
	// Given
	db, err := StartDB()
	if err != nil {
		t.Error(err)
		return
	}

	newClassroom := Classroom{ID: 666, Name: "Test", DepartmentID: 666}

	// Perform
	err = CreateClassroom(db, newClassroom)
	if err == nil {
		t.Error(err)
		return
	}

	fetchedClassroom, err := GetClassroomByID(db, newClassroom.ID)

	// Assert
	if err == nil {
		t.Error(err)
		return
	}

	if fetchedClassroom == newClassroom {
		t.Errorf("Expected %v, got %v", newClassroom, fetchedClassroom)
		return
	}

	err = DeleteClassroom(db, newClassroom.ID)
	if err != nil {
		t.Error(err)
		return
	}

	return
}

func Test_DeleteClassroom(t *testing.T) {
	// Given
	db, err := StartDB()
	if err != nil {
		t.Error(err)
		return
	}

	newDepartment := department.Department{ID: 666, Name: "Test"}
	err = department.CreateDepartment(db, newDepartment)
	if err != nil {
		t.Error(err)
		return
	}

	newClassroom := Classroom{ID: 666, Name: "Test", DepartmentID: newDepartment.ID}
	err = CreateClassroom(db, newClassroom)
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	err = DeleteClassroom(db, newClassroom.ID)

	// Assert
	if err != nil {
		t.Error(err)
		return
	}

	err = department.DeleteDepartment(db, newDepartment.ID)
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
