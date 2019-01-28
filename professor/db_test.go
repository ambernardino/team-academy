package professor

import (
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func Test_CreateProfessor(t *testing.T) {
	// Given
	professor := Professor{ID: 54, FirstName: "Paulo", LastName: "Pinto", CursoIDs: "MIEEC", CadeiraIDS: "PM", StartDate: time.Now().UTC()}
	db, err := initializeDB()
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	err = CreateProfessor(db, professor)
	if err != nil {
		t.Error(err)
		return
	}
	testProfessor, err := GetProfessorByID(db, professor.ID)
	if err != nil {
		t.Error(err)
		return
	}

	// Assert
	if professor != testProfessor {
		t.Errorf("Expected: %v Got: %v.", professor, testProfessor)
		return
	}
	return
}

func Test_UpdateProfessorInfo(t *testing.T) {
	// Given
	professor := Professor{ID: 1, FirstName: "Paulo", LastName: "Pinto", CursoIDs: "MIEEC", CadeiraIDS: "PM", StartDate: time.Now().UTC()}
	db, err := initializeDB()
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	err = CreateProfessor(db, professor)
	if err != nil {
		t.Error(err)
		return
	}
	professor = Professor{ID: 1, FirstName: "Mário", LastName: "Ventim", CursoIDs: "MIEEC", CadeiraIDS: "ET", StartDate: time.Now().UTC()}
	err = UpdateProfessorInfo(db, professor)
	if err != nil {
		t.Error(err)
		return
	}
	testProfessor, err := GetProfessorByID(db, professor.ID)
	if err != nil {
		t.Error(err)
		return
	}

	// Assert
	if professor != testProfessor {
		t.Errorf("Expected: %v Got: %v.", professor, testProfessor)
		return
	}
	return
}

func Test_DeleteProfessor(t *testing.T) {
	// Given
	professor := Professor{ID: 1, FirstName: "Paulo", LastName: "Pinto", CursoIDs: "MIEEC", CadeiraIDS: "PM", StartDate: time.Now().UTC()}
	db, err := initializeDB()
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	err = CreateProfessor(db, professor)
	if err != nil {
		t.Error(err)
		return
	}
	err = DeleteProfessor(db, professor.ID)
	if err != nil {
		t.Error(err)
		return
	}

	// Assert
	testProfessor, err := GetProfessorByID(db, professor.ID)
	if err != nil {
		return
	}
	t.Errorf("Expected: %v Got: %v.", err, testProfessor)
	return
}

func initializeDB() (DB *gorm.DB, err error) {
	DB, err = gorm.Open("sqlite3", "../clip_holy_grail.db")
	if err != nil {
		return
	}
	DB.SingularTable(true)
	return
}
