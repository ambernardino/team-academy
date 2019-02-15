package student_shift

import (
	"fmt"
	"team-academy/classroom"
	"team-academy/department"
	"team-academy/shift"
	"team-academy/student"
	"team-academy/subject"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func Test_AddStudentToShift(t *testing.T) {
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

	newDepartment := department.Department{ID: 666, Name: "Test"}
	err = department.CreateDepartment(db, newDepartment)
	if err != nil {
		t.Error(err)
		return
	}

	newClassroom := classroom.Classroom{ID: 666, Name: "Test", DepartmentID: 666}
	err = classroom.CreateClassroom(db, newClassroom)
	if err != nil {
		t.Error(err)
		return
	}

	newShift := shift.Shift{ID: 666, Type: "T", ShiftNum: 666, SubjectID: newSubject.ID, ClassroomID: newClassroom.ID}
	err = shift.CreateShift(db, newShift)
	if err != nil {
		t.Error(err)
		return
	}

	err = AddStudentToShift(db, newStudent.ID, newShift.ID)
	if err != nil {
		t.Error("Couldn't add student to shift")
		return
	}

	fetchedStudentShift, err := GetStudentShift(db, newStudent.ID, newShift.ID)
	if err != nil {
		t.Error("Couldn't fetch student from shift")
		return
	}

	if fetchedStudentShift.StudentID != newStudent.ID || fetchedStudentShift.ShiftID != newShift.ID {
		t.Error("The fetched student/shift is diferent from the created")
		return
	}

	err = RemoveStudentFromShift(db, newStudent.ID, newShift.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = shift.DeleteShift(db, newShift.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = classroom.DeleteClassroom(db, newClassroom.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = department.DeleteDepartment(db, newDepartment.ID)
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

func Test_AddNonExistantStudentToShift(t *testing.T) {
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

	newDepartment := department.Department{ID: 666, Name: "Test"}
	err = department.CreateDepartment(db, newDepartment)
	if err != nil {
		t.Error(err)
		return
	}

	newClassroom := classroom.Classroom{ID: 666, Name: "Test", DepartmentID: 666}
	err = classroom.CreateClassroom(db, newClassroom)
	if err != nil {
		t.Error(err)
		return
	}

	newShift := shift.Shift{ID: 666, Type: "T", ShiftNum: 666, SubjectID: newSubject.ID, ClassroomID: newClassroom.ID}
	err = shift.CreateShift(db, newShift)
	if err != nil {
		t.Error(err)
		return
	}

	err = AddStudentToShift(db, 666, newShift.ID)
	if err == nil {
		t.Error("Added non existant student to shift")
		return
	}

	_, err = GetStudentShift(db, 666, newShift.ID)
	if err == nil {
		t.Error("Fetched non existant student from shift")
		return
	}

	err = RemoveStudentFromShift(db, 666, newShift.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = shift.DeleteShift(db, newShift.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = classroom.DeleteClassroom(db, newClassroom.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = department.DeleteDepartment(db, newDepartment.ID)
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

func Test_AddStudentToNonExistantShift(t *testing.T) {
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

	err = AddStudentToShift(db, newStudent.ID, 666)
	if err == nil {
		t.Error("Added student to non existant shift")
		return
	}

	_, err = GetStudentShift(db, newStudent.ID, 666)
	if err == nil {
		t.Error("Fetched student from non existant shift")
		return
	}

	err = RemoveStudentFromShift(db, newStudent.ID, 666)
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

func Test_RemoveStudentFromShift(t *testing.T) {
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

	newDepartment := department.Department{ID: 666, Name: "Test"}
	err = department.CreateDepartment(db, newDepartment)
	if err != nil {
		t.Error(err)
		return
	}

	newClassroom := classroom.Classroom{ID: 666, Name: "Test", DepartmentID: 666}
	err = classroom.CreateClassroom(db, newClassroom)
	if err != nil {
		t.Error(err)
		return
	}

	newShift := shift.Shift{ID: 666, Type: "T", ShiftNum: 666, SubjectID: newSubject.ID, ClassroomID: newClassroom.ID}
	err = shift.CreateShift(db, newShift)
	if err != nil {
		t.Error(err)
		return
	}

	err = AddStudentToShift(db, newStudent.ID, newShift.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = RemoveStudentFromShift(db, newStudent.ID, newShift.ID)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = GetStudentShift(db, newStudent.ID, newShift.ID)
	if err == nil {
		t.Error("Fetched non existant student/shift")
		return
	}

	err = shift.DeleteShift(db, newShift.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = classroom.DeleteClassroom(db, newClassroom.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = department.DeleteDepartment(db, newDepartment.ID)
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

func Test_GetStudentShift(t *testing.T) {
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

	newDepartment := department.Department{ID: 666, Name: "Test"}
	err = department.CreateDepartment(db, newDepartment)
	if err != nil {
		t.Error(err)
		return
	}

	newClassroom := classroom.Classroom{ID: 666, Name: "Test", DepartmentID: 666}
	err = classroom.CreateClassroom(db, newClassroom)
	if err != nil {
		t.Error(err)
		return
	}

	newShift := shift.Shift{ID: 666, Type: "T", ShiftNum: 666, SubjectID: newSubject.ID, ClassroomID: newClassroom.ID}
	err = shift.CreateShift(db, newShift)
	if err != nil {
		t.Error(err)
		return
	}

	err = AddStudentToShift(db, newStudent.ID, newShift.ID)
	if err != nil {
		t.Error(err)
		return
	}

	fetchedStudentShift, err := GetStudentShift(db, newStudent.ID, newShift.ID)
	if err != nil {
		t.Error("Couldn't fetch student/shift")
		return
	}

	if fetchedStudentShift.StudentID != newStudent.ID || fetchedStudentShift.ShiftID != newShift.ID {
		t.Error("The fetched student/shift is diferent from the created")
		return
	}

	err = RemoveStudentFromShift(db, newStudent.ID, newShift.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = shift.DeleteShift(db, newShift.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = classroom.DeleteClassroom(db, newClassroom.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = department.DeleteDepartment(db, newDepartment.ID)
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

func StartDB() (db *gorm.DB, err error) {
	db, err = gorm.Open("sqlite3", "../clip_holy_grail.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	db.SingularTable(true)
	return
}
