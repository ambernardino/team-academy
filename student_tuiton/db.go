package student_tuiton

import (
	"team-academy/component"

	"github.com/jinzhu/gorm"
)

type StudentTuition struct {
	ID          int `gorm:"AUTO_INCREMENT"`
	StudentID   int
	StudentDebt string
	Reference   string
	Entity      string
	Date        int64
}

func CreateTableIfNotExists(db *gorm.DB) (exists bool, err error) {
	if !db.HasTable(StudentTuition{}) {
		return false, db.CreateTable(StudentTuition{}).Error
	}
	return true, nil
}

func AddStudentTuition(db *gorm.DB, st StudentTuition) (err error) {
	return db.Save(&st).Error
}

func UpdateStudentTuition(db *gorm.DB, st StudentTuition) (err error) {
	_, err = GetStudentTuitionByStudentID(db, st.StudentID)
	if err != nil {
		return component.ErrStudentDoesntExist
	} else if st.StudentID <= 0 {
		return component.ErrMissingParameters
	}

	return db.Model(&StudentTuition{}).Updates(&st).Error
}

func DeleteStudentTuition(db *gorm.DB, id int) (err error) {
	return db.Delete(&StudentTuition{ID: id}).Error
}

func GetStudentTuitionByStudentID(db *gorm.DB, studentID int) (st StudentTuition, err error) {
	err = db.First(&st, &StudentTuition{StudentID: studentID}).Error
	return
}
