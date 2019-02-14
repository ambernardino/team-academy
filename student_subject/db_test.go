package student_subject

import (
	"fmt"
	"team-academy/student"
	"team-academy/subject"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func Test_AddStudentToSubject(t *testing.T) {
		return
	}

	newStudent := student.Student{ID: 666, FirstName: "Test", LastName: "Test", CursoID: 666, StartDate: time.Now().UTC().Unix(), Email: "t.test_666@campus.fct.unl.pt"}
	err = student.CreateStudent(db, newStudent)
	if err != nil {
		return
	}

	newSubject := subject.Subject{ID: 666, Name: "Test", Description: "Test"}
	err = subject.CreateSubject(db, newSubject)
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	err = AddStudentToSubject(db, newStudent.ID, newSubject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	fetchedStudentSubject, err := GetStudentSubject(db, newStudent.ID, newSubject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	// Assert
	if fetchedStudentSubject.StudentID != newStudent.ID || fetchedStudentSubject.SubjectID != newSubject.ID {
		t.Errorf("Expected %v %v, got %v %v", newStudent.ID, newSubject.ID, fetchedStudentSubject.StudentID, fetchedStudentSubject.SubjectID)
		return
	}

	err = RemoveStudentFromSubject(db, newStudent.ID, newSubject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = subject.DeleteSubject(db, newSubject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = student.DeleteStudent(db, newStudent.ID)
	if err != nil {
		t.Error(err)
		return
	}

	return
}

func Test_AddStudentToNonExistantSubject(t *testing.T) {
		return
	}

	newStudent := student.Student{ID: 666, FirstName: "Test", LastName: "Test", CursoID: 666, StartDate: time.Now().UTC().Unix(), Email: "t.test_666@campus.fct.unl.pt"}
	err = student.CreateStudent(db, newStudent)
	if err != nil {
		return
	}

	// Perform
	err = AddStudentToSubject(db, newStudent.ID, 666)
	if err == nil {
		t.Error(err)
		return
	}

	fetchedStudentSubject, err := GetStudentSubject(db, newStudent.ID, 666)
	if err == nil {
		t.Error(err)
		return
	}

	// Assert
	if fetchedStudentSubject.StudentID == newStudent.ID && fetchedStudentSubject.SubjectID == 666 {
		t.Errorf("Expected %v %v, got %v %v", newStudent.ID, 666, fetchedStudentSubject.StudentID, fetchedStudentSubject.SubjectID)
		return
	}

	err = RemoveStudentFromSubject(db, newStudent.ID, 666)
	if err != nil {
		t.Error(err)
		return
	}

	err = student.DeleteStudent(db, newStudent.ID)
	if err != nil {
		t.Error(err)
		return
	}

	return
}

func Test_AddNonExistantStudentToSubject(t *testing.T) {
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
	err = AddStudentToSubject(db, 666, newSubject.ID)
	if err == nil {
		t.Error(err)
		return
	}

	fetchedStudentSubject, err := GetStudentSubject(db, 666, newSubject.ID)
	if err == nil {
		t.Error(err)
		return
	}

	// Assert
	if fetchedStudentSubject.StudentID == newSubject.ID && fetchedStudentSubject.SubjectID == 666 {
		t.Errorf("Expected %v %v, got %v %v", 666, newSubject.ID, fetchedStudentSubject.StudentID, fetchedStudentSubject.SubjectID)
		return
	}

	err = RemoveStudentFromSubject(db, 666, newSubject.ID)
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

func Test_AddRepeatedStudentToSubject(t *testing.T) {
	// Given
	db, err := StartDB()
	if err != nil {
		t.Error(err)
		return
	}

	newStudent := student.Student{ID: 666, FirstName: "Test", LastName: "Test", CursoID: 666, StartDate: time.Now().UTC().Unix(), Email: "t.test_666@campus.fct.unl.pt"}
	err = student.CreateStudent(db, newStudent)
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

	err = AddStudentToSubject(db, newStudent.ID, newSubject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	fetchedStudentSubject, err := GetStudentSubject(db, newStudent.ID, newSubject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	err = AddStudentToSubject(db, newStudent.ID, newSubject.ID)
	if err == nil {
		t.Error(err)
		return
	}

	repeatedStudentSubject, err := GetStudentSubject(db, newStudent.ID, newSubject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	// Assert
	if repeatedStudentSubject != fetchedStudentSubject {
		t.Errorf("Expected %v, got %v", fetchedStudentSubject, repeatedStudentSubject)
		return
	}

	err = RemoveStudentFromSubject(db, newStudent.ID, newSubject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = subject.DeleteSubject(db, newSubject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = student.DeleteStudent(db, newStudent.ID)
	if err != nil {
		t.Error(err)
		return
	}

	return
}

		return
	}

	newStudent := student.Student{ID: 666, FirstName: "Test", LastName: "Test", CursoID: 666, StartDate: time.Now().UTC().Unix(), Email: "t.test_666@campus.fct.unl.pt"}
	err = student.CreateStudent(db, newStudent)
	if err != nil {
		t.Error(err)
	}

	newSubject := subject.Subject{ID: 666, Name: "Test", Description: "Test"}
	err = subject.CreateSubject(db, newSubject)
	if err != nil {
		t.Error(err)
		return
	}

	err = AddStudentToSubject(db, newStudent.ID, newSubject.ID)
	if err != nil {
		t.Error(err)
		return
	}

<<<<<<< HEAD
	time := time.Now().UTC().Unix()

	err = AddStudentToSubject(db, testStudent.ID, testSubject.ID, time)
=======
	// Perform
	}

		t.Error("StudentSubject not properly removed")
		return

		return
	}

	if err != nil {
		t.Error(err)
		return
	}

	err = student.DeleteStudent(db, newStudent.ID)
	if err != nil {
		t.Error(err)
		return
	}

	return
}

	if err != nil {
		t.Error(err)
		return
	}

	newSubject := subject.Subject{ID: 666, Name: "Test", Description: "Test"}
	err = subject.CreateSubject(db, newSubject)
	if err != nil {
		t.Error(err)
	}

	err = AddStudentToSubject(db, newStudent.ID, newSubject.ID)
	if err != nil {
		t.Error(err)
	}

	// Perform
	fetchedStudentSubject, err := GetStudentSubject(db, newStudent.ID, newSubject.ID)
	if err != nil {
		t.Error(err)
		return
	}


	err = RemoveStudentFromSubject(db, newStudent.ID, newSubject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	if err != nil {
		t.Error(err)
		return
	}

	err = student.DeleteStudent(db, newStudent.ID)
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