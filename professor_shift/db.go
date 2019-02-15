package professor_shift

import (
	"team-academy/component"
	"team-academy/professor"
	"team-academy/shift"

	"github.com/jinzhu/gorm"
)

type ProfessorShift struct {
	ID          int `gorm:"AUTO_INCREMENT"`
	ProfessorID int
	ShiftID     int
}

func CreateTableIfNotExists(db *gorm.DB) (exists bool, err error) {
	if !db.HasTable(ProfessorShift{}) {
		return false, db.CreateTable(ProfessorShift{}).Error
	}
	return true, nil
}

func AddProfessorToShift(db *gorm.DB, professorID, shiftID int) (err error) {
	_, err = professor.GetProfessorByID(db, professorID)
	if err != nil {
		err = component.ErrProfessorDoesntExist
		return
	}

	_, err = shift.GetShiftByID(db, shiftID)
	if err != nil {
		err = component.ErrShiftDoesntExist
		return
	}

	rows, err := db.Table("professor_shift").Select("professor_id, subject_id").Joins("JOIN professor ON professor_shift.professor_id = professor.id").Joins("JOIN shift ON professor_shift.shift_id = shift.id").Where(&ProfessorShift{ShiftID: shiftID}).Rows()
	if err != nil {
		return
	}

	defer rows.Close()
	if rows.Next() {
		err = component.ErrProfessorAlreadyInShift
		return
	}

	return db.Save(&ProfessorShift{ProfessorID: professorID, ShiftID: shiftID}).Error
}

func RemoveProfessorFromShift(db *gorm.DB, professorID, shiftID int) (err error) {
	return db.Where(&ProfessorShift{ProfessorID: professorID, ShiftID: shiftID}).Delete(&ProfessorShift{}).Error
}

func GetProfessorShift(db *gorm.DB, professorID, shiftID int) (professorShift ProfessorShift, err error) {
	err = db.First(&professorShift, &ProfessorShift{ProfessorID: professorID, ShiftID: shiftID}).Error
	return
}

func GetShiftsByProfessorID(db *gorm.DB, id int) (shifts []shift.Shift, err error) {
	err = db.Model(&ProfessorShift{}).Where(&ProfessorShift{ProfessorID: id}).Find(&shifts).Error
	return
}
