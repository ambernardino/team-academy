package professor

import (
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func Test_CreateProfessors(t *testing.T) {
	// Given
	prof := Professor{ID: 6, FirstName: "Paulo", LastName: "Montezuma", CursoIDs: "MIEEC", CadeiraIDS: "PM", StartDate: time.Now().UTC()}
	db, err := initializeDB()
	if err != nil {
		t.Error(err)
		return
	}
	// Perform
	err = CreateProfessor(db, prof)
	if err != nil {
		t.Error(err)
		return
	}
	prof2, err := GetProfessorByID(db, prof.ID)
	if prof == prof2 {
		return
	}
	t.Errorf("Expected: %v Got: %v", prof, prof2)
}

func Test_UpdateProfessorInfo(t *testing.T) {
	prof := Professor{ID: 6, FirstName: "Pinto", LastName: "Pimentão", CursoIDs: "MIEEC", CadeiraIDS: "PM", StartDate: time.Now().UTC()}
	profUpdated := Professor{ID: 6, FirstName: "Paulo", LastName: "Montezuma", CursoIDs: "MIEEC", CadeiraIDS: "PM", StartDate: time.Now().UTC()}
	db, err := initializeDB()
	if err != nil {
		t.Error(err)
		return
	}
	err = CreateProfessor(db, prof)
	if err != nil {
		t.Error(err)
		return
	}
	err = UpdateProfessorInfo(db, profUpdated)
	if err != nil {
		t.Error(err)
		return
	}
	GetProf, err := GetProfessorByID(db, profUpdated.ID)
	if err != nil {
		t.Error(err)
		return
	}
	if GetProf == profUpdated {
		return
	}
	t.Fatalf("Expected: %v Got: %v", profUpdated, GetProf)
}

func Test_RemoveProfessor(t *testing.T) {
	prof := Professor{ID: 6, FirstName: "Paulo", LastName: "Montezuma", CursoIDs: "MIEEC", CadeiraIDS: "PM", StartDate: time.Now().UTC()}
	db, err := initializeDB()
	if err != nil {
		t.Error(err)
		return
	}
	err = CreateProfessor(db, prof)
	if err != nil {
		t.Error(err)
		return
	}
	err = DeleteProfessor(db, prof.ID)
	if err != nil {
		t.Error(err)
		return
	}
	getprof, err := GetProfessorByID(db, prof.ID)
	if err == nil {
		t.Fatalf("Expected %v Got: %v", getprof, prof)
	}
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
