package student_subject

import (
	"github.com/jinzhu/gorm"
)

type StudentSubject struct {
	ID        int `gorm:"AUTO_INCREMENT"`
	StudentID int
	SubjectID int
}

func CreateTableIfNotExists(db *gorm.DB) (err error) {
	if !db.HasTable(StudentSubject{}) {
		return db.CreateTable(StudentSubject{}).Error
	}

	return
}

func AddStudentToSubject(db *gorm.DB, studentID, subjectID int) (err error) {
	newStudentSubject := StudentSubject{StudentID: studentID, SubjectID: subjectID}
	return db.Save(&newStudentSubject).Error
}

func RemoveStudentFromSubject(db *gorm.DB, studentID, subjectID int) (err error) {
	return db.Delete(&StudentSubject{StudentID: studentID, SubjectID: subjectID}).Error
}

func GetSubjectsByStudentID(db *gorm.DB, id int) (subjects []StudentSubject, err error) {
	err = db.Model(&StudentSubject{}).Find(&subjects).Where(&StudentSubject{StudentID: id}).Error
	return
}

func GetStudentsBySubjectID(db *gorm.DB, id int) (students []StudentSubject, err error) {
	err = db.Model(&StudentSubject{}).Find(&students).Where(&StudentSubject{SubjectID: id}).Error
	return
}
