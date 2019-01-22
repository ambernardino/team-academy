package student

import (
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func Test_CreateStudent(t *testing.T) {
	db, err := initializeDB()
	if err != nil {
		t.Error("DB not initialized")
		return
	}
	st := Student{ID: 1, FirstName: "aaaaa", LastName: "bbbbb", DegreeID: 10, StartDate: time.Now().UTC()}
	err = CreateStudent(db, st)
	if err != nil {
		t.Error("Student not created")
		return
	}

	var st2 Student
	st2, err = GetStudentByID(db, st.ID)
	if err != nil {
		t.Error("Can't find student")
		return
	}

	if st == st2 {
		return
	}
	t.Fatalf("Expected: %v Got: %v", st, st2)
}

func Test_UpdateStudent(t *testing.T) {
	db, err := initializeDB()
	if err != nil {
		t.Error("DB not initialized")
		return
	}

	newStudent := Student{ID: 1, FirstName: "ccccc", LastName: "ddddd", DegreeID: 10, StartDate: time.Now().UTC()}
	err = UpdateStudent(db, newStudent)
	if err != nil {
		t.Error("Update Student Error")
		return
	}

	//var newStudent2 Student
	newStudent2, err := GetStudentByID(db, newStudent.ID)
	if err != nil {
		t.Error("Can't find student")
		return
	}

	if newStudent == newStudent2 {
		return
	}
	t.Fatalf("Expected: %v Got: %v", newStudent, newStudent2)
}

func Test_DeleteStudent(t *testing.T) {
	db, err := initializeDB()
	if err != nil {
		t.Error("DB not initialized")
		return
	}

	st := Student{ID: 1, FirstName: "ccccc", LastName: "ddddd", DegreeID: 10, StartDate: time.Now().UTC()}
	err = CreateStudent(db, st)
	if err != nil {
		t.Error("Couldn't create Student")
		return
	}

	err = DeleteStudent(db, st.ID)
	if err != nil {
		t.Error("Couldn't delete student")
		return
	}

	st2, err := GetStudentByID(db, st.ID)
	if err != nil {
		return
	}
	t.Fatalf("Found unexpected student: %v", st2)
}

func initializeDB() (DB *gorm.DB, err error) {
	DB, err = gorm.Open("sqlite3", "../clip_holy_grail.db")
	if err != nil {
		return
	}

	DB.SingularTable(true)
	return
}
