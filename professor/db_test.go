package professor

import (
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func Test_CreateProfessor(t *testing.T) {
	// Given
	db, err := StartDB()
	if err != nil {
		t.Error(err)
		return
	}

	newProfessor := Professor{ID: 666, FirstName: "Test", LastName: "Test", CursoID: 666, StartDate: time.Now().UTC().Unix(), Email: "t.test_666@fct.unl.pt"}
	err = CreateProfessor(db, newProfessor)
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	testProfessor, err := GetProfessorByID(db, newProfessor.ID)
	if testProfessor != newProfessor {
		t.Errorf("Expected %v, got %v", newProfessor, testProfessor)
		return
	}

	err = DeleteProfessor(db, newProfessor.ID)
	if err != nil {
		t.Error(err)
		return
	}

	return
}

func Test_UpdateProfessorInfo(t *testing.T) {
	// Given
	db, err := StartDB()
	if err != nil {
		t.Error(err)
		return
	}

	newProfessor := Professor{ID: 666, FirstName: "Test", LastName: "Test", CursoID: 666, StartDate: time.Now().UTC().Unix(), Email: "t.test_666@fct.unl.pt"}
	err = CreateProfessor(db, newProfessor)
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	updatedProfessor := Professor{ID: 666, FirstName: "Updated", LastName: "Updated", CursoID: 666, StartDate: time.Now().UTC().Unix(), Email: "u.updated_666@fct.unl.pt"}
	err = UpdateProfessorInfo(db, updatedProfessor)
	if err != nil {
		t.Error(err)
		return
	}

	fetchedProfessor, err := GetProfessorByID(db, newProfessor.ID)
	if err != nil {
		t.Error(err)
		return
	}

	// Assert
	if fetchedProfessor != updatedProfessor {
		t.Errorf("Expected %v, got %v", updatedProfessor, fetchedProfessor)
		return
	}

	err = DeleteProfessor(db, newProfessor.ID)
	if err != nil {
		t.Error(err)
		return
	}

	return
}

func Test_DeleteProfessor(t *testing.T) {
	// Given
	db, err := StartDB()
	if err != nil {
		t.Error(err)
		return
	}

	newProfessor := Professor{ID: 666, FirstName: "Test", LastName: "Test", CursoID: 666, StartDate: time.Now().UTC().Unix(), Email: "t.test_666@fct.unl.pt"}
	err = CreateProfessor(db, newProfessor)
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	err = DeleteProfessor(db, newProfessor.ID)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = GetProfessorByID(db, newProfessor.ID)

	// Assert
	if err == nil {
		t.Error("Professor wasn't properly deleted.")
		return
	}

	return
}

func Test_GetProfessorByID(t *testing.T) {
	// Given
	db, err := StartDB()
	if err != nil {
		t.Error(err)
		return
	}

	newProfessor := Professor{ID: 666, FirstName: "Test", LastName: "Test", CursoID: 666, StartDate: time.Now().UTC().Unix(), Email: "t.test_666@fct.unl.pt"}
	err = CreateProfessor(db, newProfessor)
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	fetchedProfessor, err := GetProfessorByID(db, newProfessor.ID)
	if err != nil {
		t.Error(err)
		return
	}

	// Assert
	if fetchedProfessor != newProfessor {
		t.Errorf("Expected %v, got %v", newProfessor, fetchedProfessor)
		return
	}

	err = DeleteProfessor(db, newProfessor.ID)
	if err != nil {
		t.Error(err)
		return
	}

	return
}

func Test_GetProfessorByEmail(t *testing.T) {
	// Given
	db, err := StartDB()
	if err != nil {
		t.Error(err)
		return
	}

	newProfessor := Professor{ID: 666, FirstName: "Test", LastName: "Test", CursoID: 666, StartDate: time.Now().UTC().Unix(), Email: "t.test_666@fct.unl.pt"}
	err = CreateProfessor(db, newProfessor)
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	fetchedProfessor, err := GetProfessorByEmail(db, newProfessor.Email)
	if err != nil {
		t.Error(err)
		return
	}

	// Assert
	if fetchedProfessor != newProfessor {
		t.Errorf("Expected %v, got %v", newProfessor, fetchedProfessor)
		return
	}

	err = DeleteProfessor(db, newProfessor.ID)
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
