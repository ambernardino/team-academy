package grades

import (
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func Test_GetStudentsGrades(t *testing.T) {
	// Given
	db, err := initializeDB()
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	grades, err := GetStudentsGrades(db)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(grades)

	// Assert
	t.Error("string")
}

func initializeDB() (DB *gorm.DB, err error) {
	DB, err = gorm.Open("sqlite3", "../clip_holy_grail.db")
	if err != nil {
		return
	}
	DB.SingularTable(true)
	return
}
