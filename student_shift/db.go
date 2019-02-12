package student_shift

import (
	"team-academy/shift"

	"github.com/jinzhu/gorm"
)

type StudentShift struct {
	ID        int `gorm:"AUTO_INCREMENT"`
	StudentID int
	SubjectID int
	ShiftID   int
}

func CreateTableIfNotExists(db *gorm.DB) (exists bool, err error) {
	if !db.HasTable(StudentShift{}) {
		return false, db.CreateTable(StudentShift{}).Error
	}
	return true, nil
}

func AddProfessorToSubjectShift(db *gorm.DB, studentID, subjectID, shiftID int) (err error) {
	return db.Save(&StudentShift{StudentID: studentID, SubjectID: subjectID, ShiftID: shiftID}).Error
}

func GetShiftsByProfessorID(db *gorm.DB, id int) (shifts []shift.Shift, err error) {
	err = db.Model(&StudentShift{}).Where(&StudentShift{StudentID: id}).Find(&shifts).Error
	return
}
