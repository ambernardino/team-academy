package student_subject

import (
	"team-academy/component"

	"github.com/jinzhu/gorm"
)

type StudentSubject struct {
	ID        int `gorm:"AUTO_INCREMENT"`
	StudentID int
	SubjectID int
}

func CreateTableIfNotExists(db *gorm.DB) (exists bool, err error) {
	if !db.HasTable(StudentSubject{}) {
		return false, db.CreateTable(StudentSubject{}).Error
	}

	return true, nil
}

func AddStudentToSubject(db *gorm.DB, studentID, subjectID int) (err error) {
	rows, err := db.Table("student").Select("student_subject.student_id, student_subject.subject_id").Joins("JOIN student_subject ON student.id = student_subject.student_id").Joins("JOIN subject ON subject.id = student_subject.subject_id").Where(&StudentSubject{StudentID: studentID, SubjectID: subjectID}).Rows()

	if rows.Next() {
		err = component.ErrStudentAlreadyInSubject
		return
	}

	return db.Save(&StudentSubject{StudentID: studentID, SubjectID: subjectID}).Error
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
