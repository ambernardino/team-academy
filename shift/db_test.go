package shift

import (
	"fmt"
	"team-academy/classroom"
	"team-academy/department"
	"team-academy/subject"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func Test_CreateShift(t *testing.T) {
	// Given
	db, err := StartDB()
	if err != nil {
		t.Error(err)
		return
	}

	newSubject := subject.Subject{ID: 666, Name: "Test", Description: "Test"}
	err = subject.CreateSubject(db, newSubject)
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

	newClassroom := classroom.Classroom{ID: 666, Name: "Test", DepartmentID: newDepartment.ID}
	err = classroom.CreateClassroom(db, newClassroom)
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	newShift := Shift{ID: 666, Type: "T", ShiftNum: 666, SubjectID: newSubject.ID, ClassroomID: newClassroom.ID}
	err = CreateShift(db, newShift)
	if err != nil {
		t.Error("Can't create shift")
		return
	}

	// Assert
	fetchedShift, err := GetShiftByID(db, newShift.ID)
	if err != nil {
		t.Error(err)
		return
	}

	if fetchedShift != newShift {
		t.Errorf("Expected %v, got %v", newShift, fetchedShift)
		return
	}

	err = DeleteShift(db, newShift.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = classroom.DeleteClassroom(db, newClassroom.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = department.DeleteDepartment(db, newDepartment.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = subject.DeleteSubject(db, newSubject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	return
}

func Test_CreateShiftOnNonExistantClassroom(t *testing.T) {
	// Given
	db, err := StartDB()
	if err != nil {
		t.Error(err)
		return
	}

	newSubject := subject.Subject{ID: 666, Name: "Test", Description: "Test"}
	err = subject.CreateSubject(db, newSubject)
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	newShift := Shift{ID: 666, Type: "T", ShiftNum: 666, SubjectID: newSubject.ID, ClassroomID: 666}
	err = CreateShift(db, newShift)
	if err == nil {
		t.Error("Shift was created without a classroom")
		return
	}

	// Assert
	fetchedShift, err := GetShiftByID(db, newShift.ID)
	if err == nil {
		t.Errorf("Got %v", fetchedShift)
		return
	}

	err = DeleteShift(db, newShift.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = subject.DeleteSubject(db, newSubject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	return
}

func Test_CreateShiftOnNonExistantSubject(t *testing.T) {
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

	newClassroom := classroom.Classroom{ID: 666, Name: "Test", DepartmentID: newDepartment.ID}
	err = classroom.CreateClassroom(db, newClassroom)
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	newShift := Shift{ID: 666, Type: "T", ShiftNum: 666, SubjectID: 666, ClassroomID: newClassroom.ID}
	err = CreateShift(db, newShift)
	if err == nil {
		t.Error("Can't create shift")
		return
	}

	// Assert
	fetchedShift, err := GetShiftByID(db, newShift.ID)
	if err == nil {
		t.Error("Shift was created anyway")
		return
	}

	if fetchedShift == newShift {
		t.Errorf("Got %v", fetchedShift)
		return
	}

	err = DeleteShift(db, newShift.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = classroom.DeleteClassroom(db, newClassroom.ID)
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

func Test_UpdateShift(t *testing.T) {
	// Given
	db, err := StartDB()
	if err != nil {
		t.Error(err)
		return
	}

	newSubject := subject.Subject{ID: 666, Name: "Test", Description: "Test"}
	err = subject.CreateSubject(db, newSubject)
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

	newClassroom := classroom.Classroom{ID: 666, Name: "Test", DepartmentID: newDepartment.ID}
	err = classroom.CreateClassroom(db, newClassroom)
	if err != nil {
		t.Error(err)
		return
	}

	newShift := Shift{ID: 666, Type: "T", ShiftNum: 666, SubjectID: newSubject.ID, ClassroomID: newClassroom.ID}
	err = CreateShift(db, newShift)
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	updatedShift := Shift{ID: 666, Type: "P", ShiftNum: 666, SubjectID: newSubject.ID, ClassroomID: newClassroom.ID}
	err = UpdateShift(db, updatedShift)
	if err != nil {
		t.Error("Couldn't update shift")
		return
	}

	// Assert
	fetchedShift, err := GetShiftByID(db, updatedShift.ID)
	if err != nil {
		t.Error(err)
		return
	}

	if fetchedShift != updatedShift {
		t.Errorf("Expected %v, got %v", updatedShift, fetchedShift)
		return
	}

	err = DeleteShift(db, newShift.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = classroom.DeleteClassroom(db, newClassroom.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = department.DeleteDepartment(db, newDepartment.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = subject.DeleteSubject(db, newSubject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	return
}

func Test_DeleteShift(t *testing.T) {
	// Given
	db, err := StartDB()
	if err != nil {
		t.Error(err)
		return
	}

	newSubject := subject.Subject{ID: 666, Name: "Test", Description: "Test"}
	err = subject.CreateSubject(db, newSubject)
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

	newClassroom := classroom.Classroom{ID: 666, Name: "Test", DepartmentID: newDepartment.ID}
	err = classroom.CreateClassroom(db, newClassroom)
	if err != nil {
		t.Error(err)
		return
	}

	newShift := Shift{ID: 666, Type: "T", ShiftNum: 666, SubjectID: newSubject.ID, ClassroomID: newClassroom.ID}
	err = CreateShift(db, newShift)
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	err = DeleteShift(db, newShift.ID)

	// Assert
	if err != nil {
		t.Error("Can't delete shift")
	}

	_, err = GetShiftByID(db, newShift.ID)
	if err == nil {
		t.Error("Shift wasn't properly deleted")
		return
	}

	err = classroom.DeleteClassroom(db, newClassroom.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = department.DeleteDepartment(db, newDepartment.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = subject.DeleteSubject(db, newSubject.ID)
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
