package student_subject

import (
	"github.com/jinzhu/gorm"
)

type StudentSubject struct {
	ID        int `gorm:"AUTO_INCREMENT"`
	StudentID int
	SubjectID int
}

func CreateTable(db *gorm.DB) (err error) {
	if !db.HasTable(StudentSubject{}) {
		return db.CreateTable(StudentSubject{}).Error
	}
	return
}

func Add(db *gorm.DB, studentID, subjectID int) (err error) {
	newSubject := StudentSubject{StudentID: studentID, SubjectID: subjectID}
	err = db.Save(&newSubject).Error
	return
}

func Remove(db *gorm.DB, id int) (err error) {
	return db.Delete(&StudentSubject{ID: id}).Error
}

func GetStudent(db *gorm.DB, studentID int) (err error) {
	return db.First(&StudentSubject{StudentID: studentID}).Error
}

func GetStudents(db *gorm.DB, id int) (subject []StudentSubject, err error) {
	err = db.Model(&StudentSubject{}).Find(&subject).Where(&StudentSubject{ID: id}).Error
	return
}
