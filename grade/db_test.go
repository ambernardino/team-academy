package grade

import (
	"team-academy/student"
	"team-academy/student_subject"
	"team-academy/subject"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func Test_GiveGrades(t *testing.T) {
	//Given
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

	err = student_subject.AddStudentToSubject(db, newStudent.ID, newSubject.ID, 666)
	if err != nil {
		t.Error(err)
		return
	}

	newGrade := Grade{ID: 666, StudentID: newStudent.ID, SubjectID: newSubject.ID, Rank: "6.66"}

	// Perform
	err = GiveGrade(db, newGrade)
	if err != nil {
		t.Error(err)
		return
	}

	fetchedGrade, err := GetGradeByStudentIDAndSubjectID(db, newGrade.StudentID, newGrade.SubjectID)

	// Assert
	if err != nil {
		t.Error(err)
		return
	}

	if fetchedGrade != newGrade {
		t.Errorf("Expected %v, got %v", newGrade, fetchedGrade)
		return
	}

	err = DeleteGrade(db, newGrade.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = student_subject.RemoveStudentFromSubject(db, newStudent.ID, newSubject.ID)
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

func Test_GiveGradeToNonExistantStudent(t *testing.T) {
	//Given
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

	err = student_subject.AddStudentToSubject(db, 666, newSubject.ID, 666)
	if err == nil {
		t.Error(err)
		return
	}

	newGrade := Grade{ID: 666, StudentID: 666, SubjectID: newSubject.ID, Rank: "6.66"}

	// Perform
	err = GiveGrade(db, newGrade)
	if err == nil {
		t.Error(err)
		return
	}

	fetchedGrade, err := GetGradeByStudentIDAndSubjectID(db, newGrade.StudentID, newGrade.SubjectID)

	// Assert
	if err == nil {
		t.Error(err)
		return
	}

	if fetchedGrade == newGrade {
		t.Errorf("Expected %v, got %v", newGrade, fetchedGrade)
		return
	}

	err = DeleteGrade(db, newGrade.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = student_subject.RemoveStudentFromSubject(db, 666, newSubject.ID)
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

func Test_GiveGradeToNonExistantSubject(t *testing.T) {
	//Given
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

	err = student_subject.AddStudentToSubject(db, newStudent.ID, 666, 666)
	if err == nil {
		t.Error(err)
		return
	}

	newGrade := Grade{ID: 666, StudentID: newStudent.ID, SubjectID: 666, Rank: "6.66"}

	// Perform
	err = GiveGrade(db, newGrade)
	if err == nil {
		t.Error(err)
		return
	}

	fetchedGrade, err := GetGradeByStudentIDAndSubjectID(db, newGrade.StudentID, 666)

	// Assert
	if err == nil {
		t.Error(err)
		return
	}

	if fetchedGrade == newGrade {
		t.Errorf("Expected %v, got %v", newGrade, fetchedGrade)
		return
	}

	err = DeleteGrade(db, newGrade.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = student_subject.RemoveStudentFromSubject(db, newStudent.ID, 666)
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

func Test_GiveGradeToNonExistantStudentSubject(t *testing.T) {
	//Given
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

	newGrade := Grade{ID: 666, StudentID: newStudent.ID, SubjectID: newSubject.ID, Rank: "6.66"}

	// Perform
	err = GiveGrade(db, newGrade)
	if err == nil {
		t.Error(err)
		return
	}

	fetchedGrade, err := GetGradeByStudentIDAndSubjectID(db, newGrade.StudentID, newGrade.SubjectID)

	// Assert
	if err == nil {
		t.Error(err)
		return
	}

	if fetchedGrade == newGrade {
		t.Errorf("Expected %v, got %v", newGrade, fetchedGrade)
		return
	}

	err = DeleteGrade(db, newGrade.ID)
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

func Test_GiveRepeatedGrade(t *testing.T) {
	//Given
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

	err = student_subject.AddStudentToSubject(db, newStudent.ID, newSubject.ID, 666)
	if err != nil {
		t.Error(err)
		return
	}

	newGrade := Grade{ID: 666, StudentID: newStudent.ID, SubjectID: newSubject.ID, Rank: "6.66"}
	err = GiveGrade(db, newGrade)
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	repeatedGrade := Grade{ID: 666, StudentID: newStudent.ID, SubjectID: newSubject.ID}

	// Assert
	err = GiveGrade(db, repeatedGrade)
	if err == nil {
		t.Error(err)
		return
	}

	fetchedGrade, err := GetGradeByStudentIDAndSubjectID(db, newStudent.ID, newSubject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	if fetchedGrade != newGrade {
		t.Errorf("Expected %v, got %v", newGrade, fetchedGrade)
		return
	}

	err = DeleteGrade(db, newGrade.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = student_subject.RemoveStudentFromSubject(db, newStudent.ID, newSubject.ID)
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

func Test_UpdateGrade(t *testing.T) {
	//Given
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

	err = student_subject.AddStudentToSubject(db, newStudent.ID, newSubject.ID, 666)
	if err != nil {
		t.Error(err)
		return
	}

	newGrade := Grade{ID: 666, StudentID: newStudent.ID, SubjectID: newSubject.ID, Rank: "6.66", Date: 666}
	err = GiveGrade(db, newGrade)
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	updatedGrade := Grade{ID: 666, StudentID: newStudent.ID, SubjectID: newSubject.ID, Rank: "66.6", Date: 666}
	err = UpdateGrade(db, updatedGrade)
	if err != nil {
		t.Error("Can't update grade")
		return
	}

	fetchedGrade, err := GetGradeByStudentIDAndSubjectID(db, newStudent.ID, newSubject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	// Assert
	if fetchedGrade != updatedGrade {
		t.Errorf("Expected %v, got %v", updatedGrade, fetchedGrade)
		return
	}

	err = DeleteGrade(db, newGrade.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = student_subject.RemoveStudentFromSubject(db, newStudent.ID, newSubject.ID)
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

func Test_DeleteGrade(t *testing.T) {
	//Given
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

	err = student_subject.AddStudentToSubject(db, newStudent.ID, newSubject.ID, 666)
	if err != nil {
		t.Error(err)
		return
	}

	newGrade := Grade{ID: 666, StudentID: newStudent.ID, SubjectID: newSubject.ID, Rank: "6.66"}
	err = GiveGrade(db, newGrade)
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	err = DeleteGrade(db, newGrade.ID)

	// Assert
	if err != nil {
		t.Error(err)
		return
	}

	_, err = GetGradeByStudentIDAndSubjectID(db, newGrade.StudentID, newGrade.SubjectID)
	if err == nil {
		t.Error(err)
		return
	}

	err = student_subject.RemoveStudentFromSubject(db, newStudent.ID, newSubject.ID)
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

func Test_GetGrade(t *testing.T) {
	//Given
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

	err = student_subject.AddStudentToSubject(db, newStudent.ID, newSubject.ID, 666)
	if err != nil {
		t.Error(err)
		return
	}

	newGrade := Grade{ID: 666, StudentID: newStudent.ID, SubjectID: newSubject.ID, Rank: "6.66", Date: 666}
	err = GiveGrade(db, newGrade)
	if err != nil {
		t.Error(err)
		return
	}

	// Perform
	fetchedGrade, err := GetGradeByStudentIDAndSubjectID(db, newStudent.ID, newSubject.ID)
	if err != nil {
		t.Error(err)
		return
	}

	// Assert
	if fetchedGrade != newGrade {
		t.Errorf("Expected %v, got %v", newGrade, fetchedGrade)
		return
	}

	err = DeleteGrade(db, newGrade.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = student_subject.RemoveStudentFromSubject(db, newStudent.ID, newSubject.ID)
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

func StartDB() (DB *gorm.DB, err error) {
	DB, err = gorm.Open("sqlite3", "../clip_holy_grail.db")
	if err != nil {
		return
	}
	DB.SingularTable(true)
	return
}
