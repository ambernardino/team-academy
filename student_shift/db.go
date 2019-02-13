package student_shift

import (
	"team-academy/component"
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

func AddStudentToShift(db *gorm.DB, studentID int, shiftID int) (err error) {
	rows, err := db.Table("student_shift").Select("student_shift.shift_id, student_shift.student_id").Joins("JOIN shift ON student_shift.shift_id = shift.id").Joins("JOIN student ON student_shift.student_id = student.id").Where(&StudentShift{ShiftID: shiftID, StudentID: studentID}).Rows()
	if rows.Next() {
		err = component.ErrStudentAlreadyInShift
		return
	}

	return db.Save(&StudentShift{StudentID: studentID, ShiftID: shiftID}).Error
}

func GetShiftsByStudentID(db *gorm.DB, id int) (shifts []shift.Shift, err error) {
	err = db.Model(&StudentShift{}).Where(&StudentShift{StudentID: id}).Find(&shifts).Error
	return
}
