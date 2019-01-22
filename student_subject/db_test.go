package student_subject

import (
	"team-academy/student"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func Test_Add(t *testing.T) {
	subjectID := 1
	db, err := initializeDB()
	if err != nil {
		t.Error("Couldn't initialize DB")
		return
	}

	st := student.Student{ID: 1, FirstName: "ccccc", LastName: "ddddd", DegreeID: 1, StartDate: time.Now().UTC()}
	err = student.CreateStudent(db, st)
	if err != nil {
		t.Error("Couldn't create Student")
		return
	}

	err = Remove(db, 1)
	if err != nil {
		t.Error("Couldn't delete subject")
		return
	}

	err = Add(db, StudentSubject{StudentID: st.ID, SubjectID: subjectID})
	if err != nil {
		t.Error(err)
		return
	}

	err = Add(db, StudentSubject{StudentID: st.ID, SubjectID: subjectID})
	if err != nil {
		t.Error(err)
		return
	}
	t.Fatalf("Student added twice to subject")
}
func initializeDB() (DB *gorm.DB, err error) {
	DB, err = gorm.Open("sqlite3", "../clip_holy_grail.db")
	if err != nil {
		return
	}

	DB.SingularTable(true)
	return
}
