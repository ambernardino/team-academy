package schedule

import (
	"fmt"
	"team-academy/classroom"
	"team-academy/department"
	"team-academy/shift"
	"team-academy/subject"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func Test_CreateSchedule(t *testing.T) {
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

	newShift := shift.Shift{ID: 666, ShiftNum: 666, SubjectID: 666, Type: "T", ClassroomID: 666}
	err = shift.CreateShift(db, newShift)
	if err != nil {
		t.Error(err)
		return
	}

	newSchedule := Schedule{ID: 666, ShiftID: 666, Weekday: 0, StartTime: 28800, EndTime: 32400}
	err = CreateSchedule(db, newSchedule)
	if err != nil {
		t.Error("Couldn't create schedule")
		return
	}

	_, err = GetScheduleByID(db, newSchedule.ID)
	if err != nil {
		t.Error("Couldn't fetch created schedule")
		return
	}

	err = DeleteSchedule(db, newSchedule.ID)
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

func Test_CreateScheduleOnNonExistantShift(t *testing.T) {
	db, err := StartDB()
	if err != nil {
		t.Error(err)
		return
	}

	newSchedule := Schedule{ID: 666, ShiftID: 666, Weekday: 0, StartTime: 28800, EndTime: 32400}
	err = CreateSchedule(db, newSchedule)
	if err == nil {
		t.Error("Created schedule on non existant shift")
		return
	}

	_, err = GetScheduleByID(db, newSchedule.ID)
	if err == nil {
		t.Error("Fetched schedule on non existant shift")
		return
	}

	err = DeleteSchedule(db, newSchedule.ID)
	if err != nil {
		t.Error(err)
		return
	}

	return
}

func Test_CreateScheduleOnInvalidWeekday(t *testing.T) {
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

	newShift := shift.Shift{ID: 666, ShiftNum: 666, SubjectID: 666, Type: "T", ClassroomID: 666}
	err = shift.CreateShift(db, newShift)
	if err != nil {
		t.Error(err)
		return
	}

	newSchedule := Schedule{ID: 666, ShiftID: 666, Weekday: 666, StartTime: 28800, EndTime: 32400}
	err = CreateSchedule(db, newSchedule)
	if err == nil {
		t.Error("Created schedule on invalid weekday")
		return
	}

	_, err = GetScheduleByID(db, newSchedule.ID)
	if err == nil {
		t.Error("Fetched schedule on invalid weekday")
		return
	}

	err = DeleteSchedule(db, newSchedule.ID)
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

func Test_CreateScheduleOnInvalidStartTime(t *testing.T) {
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

	newShift := shift.Shift{ID: 666, ShiftNum: 666, SubjectID: 666, Type: "T", ClassroomID: 666}
	err = shift.CreateShift(db, newShift)
	if err != nil {
		t.Error(err)
		return
	}

	newSchedule := Schedule{ID: 666, ShiftID: 666, Weekday: 0, StartTime: 0, EndTime: 32400}
	err = CreateSchedule(db, newSchedule)
	if err == nil {
		t.Error("Created schedule on invalid start time")
		return
	}

	_, err = GetScheduleByID(db, newSchedule.ID)
	if err == nil {
		t.Error("Fetched schedule with invalid start time")
		return
	}

	err = DeleteSchedule(db, newSchedule.ID)
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

func Test_CreateScheduleOnInvalidEndTime(t *testing.T) {
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

	newShift := shift.Shift{ID: 666, ShiftNum: 666, SubjectID: 666, Type: "T", ClassroomID: 666}
	err = shift.CreateShift(db, newShift)
	if err != nil {
		t.Error(err)
		return
	}

	newSchedule := Schedule{ID: 666, ShiftID: 666, Weekday: 0, StartTime: 28800, EndTime: 86401}
	err = CreateSchedule(db, newSchedule)
	if err == nil {
		t.Error("Created schedule on invalid end time")
		return
	}

	_, err = GetScheduleByID(db, newSchedule.ID)
	if err == nil {
		t.Error("Fetched schedule with invalid end time")
		return
	}

	err = DeleteSchedule(db, newSchedule.ID)
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

func Test_DeleteSchedule(t *testing.T) {
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

	newShift := shift.Shift{ID: 666, ShiftNum: 666, SubjectID: 666, Type: "T", ClassroomID: 666}
	err = shift.CreateShift(db, newShift)
	if err != nil {
		t.Error(err)
		return
	}

	newSchedule := Schedule{ID: 666, ShiftID: 666, Weekday: 0, StartTime: 28800, EndTime: 32400}
	err = CreateSchedule(db, newSchedule)
	if err != nil {
		t.Error(err)
		return
	}

	err = DeleteSchedule(db, newSchedule.ID)
	if err != nil {
		t.Error("Couldn't delete schedule")
		return
	}

	_, err = GetScheduleByID(db, newSchedule.ID)
	if err == nil {
		t.Error("Fetched deleted schedule")
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

func Test_GetSubjectByID(t *testing.T) {
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

	newShift := shift.Shift{ID: 666, ShiftNum: 666, SubjectID: 666, Type: "T", ClassroomID: 666}
	err = shift.CreateShift(db, newShift)
	if err != nil {
		t.Error(err)
		return
	}

	newSchedule := Schedule{ID: 666, ShiftID: 666, Weekday: 0, StartTime: 28800, EndTime: 32400}
	err = CreateSchedule(db, newSchedule)
	if err != nil {
		t.Error(err)
		return
	}

	fetchedSchedule, err := GetScheduleByID(db, newSchedule.ID)
	if err != nil {
		t.Error("Couldn't fetch schedule")
		return
	}

	if fetchedSchedule != newSchedule {
		t.Errorf("Expected %v, got %v", newSchedule, fetchedSchedule)
		return
	}

	err = DeleteSchedule(db, newSchedule.ID)
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

func StartDB() (db *gorm.DB, err error) {
	db, err = gorm.Open("sqlite3", "../clip_holy_grail.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	db.SingularTable(true)
	return
}
