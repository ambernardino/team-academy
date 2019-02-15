package professor_shift

import (
	"fmt"
	"team-academy/classroom"
	"team-academy/department"
	"team-academy/professor"
	"team-academy/shift"
	"team-academy/subject"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func Test_AddProfessorToShift(t *testing.T) {
	db, err := StartDB()
	if err != nil {
		t.Error(err)
		return
	}

	newProfessor := professor.Professor{ID: 666, FirstName: "Test", LastName: "Test", CursoID: 666, StartDate: time.Now().UTC().Unix(), Email: "t.test_666@fct.unl.pt"}
	err = professor.CreateProfessor(db, newProfessor)
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

	newClassroom := classroom.Classroom{ID: 666, Name: "Test", DepartmentID: 666}
	err = classroom.CreateClassroom(db, newClassroom)
	if err != nil {
		t.Error(err)
		return
	}

	newShift := shift.Shift{ID: 666, Type: "T", ShiftNum: 666, SubjectID: 666}
	err = shift.CreateShift(db, newShift)
	if err != nil {
		t.Error(err)
		return
	}

	err = AddProfessorToShift(db, newProfessor.ID, newShift.ID)
	if err != nil {
		t.Error("Couldn't add professor to shift")
		return
	}

	fetchedProfessorShift, err := GetProfessorShift(db, newProfessor.ID, newShift.ID)
	if err != nil {
		t.Error(err)
		return
	}

	if fetchedProfessorShift.ProfessorID != newProfessor.ID || fetchedProfessorShift.ShiftID != newShift.ID {
		t.Errorf("Expected %v %v, got %v %v", newProfessor.ID, newShift.ID, fetchedProfessorShift.ProfessorID, fetchedProfessorShift.ShiftID)
		return
	}

	err = RemoveProfessorFromShift(db, newProfessor.ID, newShift.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = shift.DeleteShift(db, newShift.ID)
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

	err = professor.DeleteProfessor(db, newProfessor.ID)
	if err != nil {
		t.Error(err)
		return
	}

	return
}

func Test_AddNonExistantProfessorToShift(t *testing.T) {
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

	newClassroom := classroom.Classroom{ID: 666, Name: "Test", DepartmentID: 666}
	err = classroom.CreateClassroom(db, newClassroom)
	if err != nil {
		t.Error(err)
		return
	}

	newShift := shift.Shift{ID: 666, Type: "T", ShiftNum: 666, SubjectID: 666, ClassroomID: 666}
	err = shift.CreateShift(db, newShift)
	if err != nil {
		t.Error(err)
		return
	}

	err = AddProfessorToShift(db, 666, newShift.ID)
	if err == nil {
		t.Error("Non existant professor added to shift")
		return
	}

	fetchedProfessorShift, err := GetProfessorShift(db, 666, newShift.ID)
	if err == nil {
		t.Error("Fetched non existant professor in shift")
		return
	}

	if fetchedProfessorShift.ProfessorID == 666 && fetchedProfessorShift.ShiftID == newShift.ID {
		t.Errorf("Expected %v %v, got %v %v", 666, newShift.ID, fetchedProfessorShift.ProfessorID, fetchedProfessorShift.ShiftID)
		return
	}

	err = RemoveProfessorFromShift(db, 666, newShift.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = shift.DeleteShift(db, newShift.ID)
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

func Test_AddProfessorToNonExistantShift(t *testing.T) {
	db, err := StartDB()
	if err != nil {
		t.Error(err)
		return
	}

	newProfessor := professor.Professor{ID: 666, FirstName: "Test", LastName: "Test", CursoID: 666, StartDate: time.Now().UTC().Unix(), Email: "t.test_666@fct.unl.pt"}
	err = professor.CreateProfessor(db, newProfessor)
	if err != nil {
		t.Error(err)
		return
	}

	err = AddProfessorToShift(db, newProfessor.ID, 666)
	if err == nil {
		t.Error("Professor added to non existant shift")
		return
	}

	fetchedProfessorShift, err := GetProfessorShift(db, newProfessor.ID, 666)
	if err == nil {
		t.Error("Fetched professor from non existant shift")
		return
	}

	if fetchedProfessorShift.ProfessorID == newProfessor.ID && fetchedProfessorShift.ShiftID == 666 {
		t.Errorf("Expected %v %v, got %v %v", newProfessor.ID, 666, fetchedProfessorShift.ProfessorID, fetchedProfessorShift.ShiftID)
		return
	}

	err = RemoveProfessorFromShift(db, newProfessor.ID, 666)
	if err != nil {
		t.Error(err)
		return
	}

	err = professor.DeleteProfessor(db, newProfessor.ID)
	if err != nil {
		t.Error(err)
		return
	}

	return
}

func Test_AddRepeatedProfessorToShift(t *testing.T) {
	db, err := StartDB()
	if err != nil {
		t.Error(err)
		return
	}

	newProfessor := professor.Professor{ID: 666, FirstName: "Test", LastName: "Test", CursoID: 666, StartDate: time.Now().UTC().Unix(), Email: "t.test_666@fct.unl.pt"}
	err = professor.CreateProfessor(db, newProfessor)
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

	newClassroom := classroom.Classroom{ID: 666, Name: "Test", DepartmentID: 666}
	err = classroom.CreateClassroom(db, newClassroom)
	if err != nil {
		t.Error(err)
		return
	}

	newShift := shift.Shift{ID: 666, Type: "T", ShiftNum: 666, SubjectID: 666}
	err = shift.CreateShift(db, newShift)
	if err != nil {
		t.Error(err)
		return
	}

	err = AddProfessorToShift(db, newProfessor.ID, newShift.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = AddProfessorToShift(db, newProfessor.ID, newShift.ID)
	if err == nil {
		t.Error("Added repeated professor to shift")
		return
	}

	fetchedProfessorShift, err := GetProfessorShift(db, newProfessor.ID, newShift.ID)
	if err != nil {
		t.Error(err)
		return
	}

	if fetchedProfessorShift.ProfessorID != newProfessor.ID || fetchedProfessorShift.ShiftID != newShift.ID {
		t.Errorf("Expected %v %v, got %v %v", newProfessor.ID, newShift.ID, fetchedProfessorShift.ProfessorID, fetchedProfessorShift.ShiftID)
		return
	}

	err = RemoveProfessorFromShift(db, newProfessor.ID, newShift.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = shift.DeleteShift(db, newShift.ID)
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

	err = professor.DeleteProfessor(db, newProfessor.ID)
	if err != nil {
		t.Error(err)
		return
	}

	return
}

func Test_RemoveProfessorFromShift(t *testing.T) {
	db, err := StartDB()
	if err != nil {
		t.Error(err)
		return
	}

	newProfessor := professor.Professor{ID: 666, FirstName: "Test", LastName: "Test", CursoID: 666, StartDate: time.Now().UTC().Unix(), Email: "t.test_666@fct.unl.pt"}
	err = professor.CreateProfessor(db, newProfessor)
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

	newClassroom := classroom.Classroom{ID: 666, Name: "Test", DepartmentID: 666}
	err = classroom.CreateClassroom(db, newClassroom)
	if err != nil {
		t.Error(err)
		return
	}

	newShift := shift.Shift{ID: 666, Type: "T", ShiftNum: 666, SubjectID: 666}
	err = shift.CreateShift(db, newShift)
	if err != nil {
		t.Error(err)
		return
	}

	err = AddProfessorToShift(db, newProfessor.ID, newShift.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = RemoveProfessorFromShift(db, newProfessor.ID, newShift.ID)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = GetProfessorShift(db, newProfessor.ID, newShift.ID)
	if err == nil {
		t.Error("Fetched professor removed from shift")
		return
	}

	err = shift.DeleteShift(db, newShift.ID)
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

	err = professor.DeleteProfessor(db, newProfessor.ID)
	if err != nil {
		t.Error(err)
		return
	}

	return
}

func Test_GetProfessorShift(t *testing.T) {
	db, err := StartDB()
	if err != nil {
		t.Error(err)
		return
	}

	newProfessor := professor.Professor{ID: 666, FirstName: "Test", LastName: "Test", CursoID: 666, StartDate: time.Now().UTC().Unix(), Email: "t.test_666@fct.unl.pt"}
	err = professor.CreateProfessor(db, newProfessor)
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

	newClassroom := classroom.Classroom{ID: 666, Name: "Test", DepartmentID: 666}
	err = classroom.CreateClassroom(db, newClassroom)
	if err != nil {
		t.Error(err)
		return
	}

	newShift := shift.Shift{ID: 666, Type: "T", ShiftNum: 666, SubjectID: 666}
	err = shift.CreateShift(db, newShift)
	if err != nil {
		t.Error(err)
		return
	}

	err = AddProfessorToShift(db, newProfessor.ID, newShift.ID)
	if err != nil {
		t.Error(err)
		return
	}

	fetchedProfessorShift, err := GetProfessorShift(db, newProfessor.ID, newShift.ID)
	if err != nil {
		t.Error(err)
		return
	}

	if fetchedProfessorShift.ProfessorID != newProfessor.ID || fetchedProfessorShift.ShiftID != newShift.ID {
		t.Errorf("Expected %v %v, got %v %v", newProfessor.ID, newShift.ID, fetchedProfessorShift.ProfessorID, fetchedProfessorShift.ShiftID)
		return
	}

	err = RemoveProfessorFromShift(db, newProfessor.ID, newShift.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = shift.DeleteShift(db, newShift.ID)
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

	err = professor.DeleteProfessor(db, newProfessor.ID)
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
