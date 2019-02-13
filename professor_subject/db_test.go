package professor_subject

import (
	"fmt"
	"team-academy/professor"
	"team-academy/subject"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func Test_AddProfessorToSubject(t *testing.T) {
	// Given
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

	// Perform
	err = AddProfessorToSubject(db, newProfessor.ID, newSubject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	fetchedProfessorSubject, err := GetProfessorSubject(db, newProfessor.ID, newSubject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	// Assert
	if fetchedProfessorSubject.ProfessorID != newProfessor.ID || fetchedProfessorSubject.SubjectID != newSubject.ID {
		t.Errorf("Expected %v %v, got %v %v", newProfessor.ID, newSubject.ID, fetchedProfessorSubject.ProfessorID, fetchedProfessorSubject.SubjectID)
		return
	}

	err = RemoveProfessorFromSubject(db, newProfessor.ID, newSubject.ID)
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

func Test_AddProfessorToNonExistantSubject(t *testing.T) {
	// Given
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

	// Perform
	err = AddProfessorToSubject(db, newProfessor.ID, 666)
	if err == nil {
		t.Error(err)
		return
	}

	fetchedProfessorSubject, err := GetProfessorSubject(db, newProfessor.ID, 666)
	if err == nil {
		t.Error(err)
		return
	}

	// Assert
	if fetchedProfessorSubject.ProfessorID == newProfessor.ID && fetchedProfessorSubject.SubjectID == 666 {
		t.Errorf("Expected %v %v, got %v %v", newProfessor.ID, 666, fetchedProfessorSubject.ProfessorID, fetchedProfessorSubject.SubjectID)
		return
	}

	err = RemoveProfessorFromSubject(db, newProfessor.ID, 666)
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

func Test_AddNonExistantProfessorToSubject(t *testing.T) {
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
	err = AddProfessorToSubject(db, 666, newSubject.ID)
	if err == nil {
		t.Error(err)
		return
	}

	fetchedProfessorSubject, err := GetProfessorSubject(db, 666, newSubject.ID)
	if err == nil {
		t.Error(err)
		return
	}

	// Assert
	if fetchedProfessorSubject.ProfessorID == newSubject.ID && fetchedProfessorSubject.SubjectID == 666 {
		t.Errorf("Expected %v %v, got %v %v", 666, newSubject.ID, fetchedProfessorSubject.ProfessorID, fetchedProfessorSubject.SubjectID)
		return
	}

	err = RemoveProfessorFromSubject(db, 666, newSubject.ID)
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

func Test_AddRepeatedProfessorToSubject(t *testing.T) {
	// Given
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

	err = AddProfessorToSubject(db, newProfessor.ID, newSubject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	fetchedProfessorSubject, err := GetProfessorSubject(db, newProfessor.ID, newSubject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	err = AddProfessorToSubject(db, newProfessor.ID, newSubject.ID)
	if err == nil {
		t.Error(err)
		return
	}

	repeatedProfessorSubject, err := GetProfessorSubject(db, newProfessor.ID, newSubject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	// Assert
	if repeatedProfessorSubject != fetchedProfessorSubject {
		t.Errorf("Expected %v, got %v", fetchedProfessorSubject, repeatedProfessorSubject)
		return
	}

	err = RemoveProfessorFromSubject(db, newProfessor.ID, newSubject.ID)
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

func Test_RemoveProfessorFromSubject(t *testing.T) {
	// Given
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

	err = AddProfessorToSubject(db, newProfessor.ID, newSubject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	err = RemoveProfessorFromSubject(db, newProfessor.ID, newSubject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = GetProfessorSubject(db, newProfessor.ID, newSubject.ID)

	// Assert
	if err == nil {
		t.Error("ProfessorSubject not properly removed")
		return
	}

	err = RemoveProfessorFromSubject(db, newProfessor.ID, newSubject.ID)
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

func Test_GetProfessorSubject(t *testing.T) {
	// Given
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

	err = AddProfessorToSubject(db, newProfessor.ID, newSubject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	fetchedProfessorSubject, err := GetProfessorSubject(db, newProfessor.ID, newSubject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	// Assert
	if fetchedProfessorSubject.ProfessorID != newProfessor.ID || fetchedProfessorSubject.SubjectID != newSubject.ID {
		t.Errorf("Expected %v %v, got %v %v", newProfessor.ID, newSubject.ID, fetchedProfessorSubject.ProfessorID, fetchedProfessorSubject.SubjectID)
		return
	}

	err = RemoveProfessorFromSubject(db, newProfessor.ID, newSubject.ID)
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

func StartDB() (DB *gorm.DB, err error) {
	DB, err = gorm.Open("sqlite3", "../clip_holy_grail.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	DB.SingularTable(true)
	return
}
