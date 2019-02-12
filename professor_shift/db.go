package professor_shift

import (
	"team-academy/shift"

	"github.com/jinzhu/gorm"
)

type ProfessorShift struct {
	ID          int `gorm:"AUTO_INCREMENT"`
	ProfessorID int
	SubjectID   int
	ShiftID     int
}

func CreateTableIfNotExists(db *gorm.DB) (exists bool, err error) {
	if !db.HasTable(ProfessorShift{}) {
		return false, db.CreateTable(ProfessorShift{}).Error
	}
	return true, nil
}

func AddProfessorToSubjectShift(db *gorm.DB, professorID, subjectID, shiftID int) (err error) {
	return db.Save(&ProfessorShift{ProfessorID: professorID, SubjectID: subjectID, ShiftID: shiftID}).Error
}

func GetShiftsByProfessorID(db *gorm.DB, id int) (shifts []shift.Shift, err error) {
	err = db.Model(&ProfessorShift{}).Where(&ProfessorShift{ProfessorID: id}).Find(&shifts).Error
	return
}
	