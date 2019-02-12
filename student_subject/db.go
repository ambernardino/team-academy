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

type Information struct {
	StudentID        int    `gorm:"column:id"`
	StudentFirstName string `gorm:"column:first_name"`
	StudentLastName  string `gorm:"column:last_name"`
	SubjectID        int    `gorm:"column:id"`
	SubjectName      string `gorm:"column:name"`
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
	return db.Where(&StudentSubject{StudentID: studentID, SubjectID: subjectID}).Delete(&StudentSubject{}).Error
}

func GetSubjectsByStudentID(db *gorm.DB, id int) (subjects []StudentSubject, err error) {
	err = db.Model(&StudentSubject{}).Where(&StudentSubject{StudentID: id}).Find(&subjects).Error
	return
}

func GetStudentsBySubjectID(db *gorm.DB, id int) (students []StudentSubject, err error) {
	err = db.Model(&StudentSubject{}).Where(&StudentSubject{SubjectID: id}).Find(&students).Error
	return
}

func GetStudentAndInfoBySubjectID(db *gorm.DB, id int) (infos []Information, err error) {
	err = db.Table("student").Select("student.id, student.first_name, student.last_name, subject.id, subject.name").Joins("JOIN student_subject ON student.id = student_subject.student_id").Joins("JOIN subject ON subject.id = student_subject.subject_id").Where(&StudentSubject{SubjectID: id}).Scan(&infos).Error
	return
}

func GetSubjectAndInfoByStudentID(db *gorm.DB, id int) (infos []Information, err error) {
	err = db.Table("student").Select("student.id, student.first_name, student.last_name, subject.id, subject.name").Joins("JOIN student_subject ON student.id = student_subject.student_id").Joins("JOIN subject ON subject.id = student_subject.subject_id").Where(&StudentSubject{StudentID: id}).Scan(&infos).Error
	return
}
