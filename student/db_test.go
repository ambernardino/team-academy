package student

import (
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func Test_CreateStudent(t *testing.T) {
	// Given
	db, err := StartDB()
	if err != nil {
		t.Error(err)
		return
	}

	newStudent := Student{ID: 666, FirstName: "Test", LastName: "Test", CursoID: 666, StartDate: time.Now().UTC().Unix(), Email: "t.test_666@campus.fct.unl.pt"}
	err = CreateStudent(db, newStudent)
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	testStudent, err := GetStudentByID(db, newStudent.ID)
	if testStudent != newStudent {
		t.Errorf("Expected %v, got %v", newStudent, testStudent)
		return
	}

	err = DeleteStudent(db, newStudent.ID)
	if err != nil {
		t.Error(err)
		return
	}

	return
}

func Test_UpdateStudentInfo(t *testing.T) {
	// Given
	db, err := StartDB()
	if err != nil {
		t.Error(err)
		return
	}

	newStudent := Student{ID: 666, FirstName: "Test", LastName: "Test", CursoID: 666, StartDate: time.Now().UTC().Unix(), Email: "t.test_666@campus.fct.unl.pt"}
	err = CreateStudent(db, newStudent)
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	updatedStudent := Student{ID: 666, FirstName: "Updated", LastName: "Updated", CursoID: 666, StartDate: time.Now().UTC().Unix(), Email: "u.updated_666@campus.fct.unl.pt"}
	err = UpdateStudent(db, updatedStudent)
	if err != nil {
		t.Error(err)
		return
	}

	fetchedStudent, err := GetStudentByID(db, newStudent.ID)
	if err != nil {
		t.Error(err)
		return
	}

	// Assert
	if fetchedStudent != updatedStudent {
		t.Errorf("Expected %v, got %v", updatedStudent, fetchedStudent)
		return
	}

	err = DeleteStudent(db, newStudent.ID)
	if err != nil {
		t.Error(err)
		return
	}

	return
}

func Test_DeleteStudent(t *testing.T) {
	// Given
	db, err := StartDB()
	if err != nil {
		t.Error(err)
		return
	}

	newStudent := Student{ID: 666, FirstName: "Test", LastName: "Test", CursoID: 666, StartDate: time.Now().UTC().Unix(), Email: "t.test_666@campus.fct.unl.pt"}
	err = CreateStudent(db, newStudent)
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	err = DeleteStudent(db, newStudent.ID)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = GetStudentByID(db, newStudent.ID)

	// Assert
	if err == nil {
		t.Error("Student wasn't properly deleted.")
		return
	}

	return
}

func Test_GetStudentByID(t *testing.T) {
	// Given
	db, err := StartDB()
	if err != nil {
		t.Error(err)
		return
	}

	newStudent := Student{ID: 666, FirstName: "Test", LastName: "Test", CursoID: 666, StartDate: time.Now().UTC().Unix(), Email: "t.test_666@campus.fct.unl.pt"}
	err = CreateStudent(db, newStudent)
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	fetchedStudent, err := GetStudentByID(db, newStudent.ID)
	if err != nil {
		t.Error(err)
		return
	}

	// Assert
	if fetchedStudent != newStudent {
		t.Errorf("Expected %v, got %v", newStudent, fetchedStudent)
		return
	}

	err = DeleteStudent(db, newStudent.ID)
	if err != nil {
		t.Error(err)
		return
	}

	return
}

func Test_GetStudentByEmail(t *testing.T) {
	// Given
	db, err := StartDB()
	if err != nil {
		t.Error(err)
		return
	}

	newStudent := Student{ID: 666, FirstName: "Test", LastName: "Test", CursoID: 666, StartDate: time.Now().UTC().Unix(), Email: "t.test_666@campus.fct.unl.pt"}
	err = CreateStudent(db, newStudent)
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	fetchedStudent, err := GetStudentByEmail(db, newStudent.Email)
	if err != nil {
		t.Error(err)
		return
	}

	// Assert
	if fetchedStudent != newStudent {
		t.Errorf("Expected %v, got %v", newStudent, fetchedStudent)
		return
	}

	err = DeleteStudent(db, newStudent.ID)
	if err != nil {
		t.Error(err)
		return
	}

	return
}

func StartDB() (db *gorm.DB, err error) {
	db, err = gorm.Open("sqlite3", "../clip_holy_grail.db")
	if err != nil {
		return
	}

	db.SingularTable(true)
	return
}
