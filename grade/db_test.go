package grade

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func Test_GiveGrades(t *testing.T) {
	//Given
	grade := Grade{StudentID: 1, SubjectID: 4, Rank: "15"}
	db, err := initializeDB()
	if err != nil {
		t.Error(err)
		return
	}
	//Perform
	err = GiveGrade(db, grade)
	if err != nil {
		t.Error(err)
		return
	}
	//Assert
	gradeGot, err := GetGradeByStudentIDAndSubjectID(db, 1, 4)
	if err != nil {
		t.Error(err)
		t.Fatalf("Expected: %v Got: %v", grade, gradeGot)
		return
	}
	if gradeGot.Rank != grade.Rank {
		t.Fatalf("Expected: %v Got: %v", grade.Rank, gradeGot.Rank)
		return
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
