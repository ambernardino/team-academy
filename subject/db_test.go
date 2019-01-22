package subject

import (
	"team-academy/component"
	//"team-academy/subject"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func Test_CreateSubject(t *testing.T) {
	subject := Subject{ID: 2, Name: "IT", Description: "Amazing"}
	db, err := initializeDB()
	if err != nil {
		t.Error(err)
		return
	}
	err = CreateSubject(db, subject)
	if err != nil {
		t.Error(err)
		return
	}
	sub2, err := GetSubjectByID(db, subject.ID)
	if err != nil {
		t.Error(err)
		return
	}
	if sub2 == subject {
		return
	}
	t.Fatalf("Expected: %v Got: %v", subject, sub2)
}

func Test_CreateDuplicateSubject(t *testing.T) {
	sub := Subject{ID: 2, Name: "ET", Description: "wow"}
	subDup := Subject{ID: 2, Name: "IT", Description: "yay"}
	db, err := initializeDB()
	if err != nil {
		t.Error(err)
		return
	}
	err = CreateSubject(db, sub)
	if err != nil {
		t.Error(err)
		return
	}
	err = CreateSubject(db, subDup)
	if err != component.ErrSomethingAlreadyExists {
		subGot, err := GetSubjectByID(db, sub.ID)
		if err != nil {
			t.Error(err)
			return
		}
		t.Fatalf("Expected: %v Got: %v", sub, subGot)
	}
	defer DeleteSubject(db, sub.ID)
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
