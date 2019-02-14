package student_subject

import (
	"team-academy/component"

	"team-academy/subject"

	"github.com/jinzhu/gorm"
)

type StudentSubject struct {
	ID        int `gorm:"AUTO_INCREMENT"`
	StudentID int
	SubjectID int
	Date      int64
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

func AddStudentToSubject(db *gorm.DB, studentID, subjectID int, date int64) (err error) {
	rows, err := db.Table("student").Select("student_subject.student_id, student_subject.subject_id").Joins("JOIN student_subject ON student.id = student_subject.student_id").Joins("JOIN subject ON subject.id = student_subject.subject_id").Where(&StudentSubject{StudentID: studentID, SubjectID: subjectID}).Rows()

	defer rows.Close()
	if rows.Next() {
		err = component.ErrStudentAlreadyInSubject
		return
	}

	return db.Save(&StudentSubject{StudentID: studentID, SubjectID: subjectID, Date: date}).Error
}

func RemoveStudentFromSubject(db *gorm.DB, studentID, subjectID int) (err error) {
	return db.Where(&StudentSubject{StudentID: studentID, SubjectID: subjectID}).Delete(&StudentSubject{}).Error
}

func GetStudentSubject(db *gorm.DB, studentID, subjectID int) (studentSubject StudentSubject, err error) {
	err = db.First(&studentSubject, &StudentSubject{StudentID: studentID, SubjectID: subjectID}).Error
	return
}

func GetSubjectsByStudentID(db *gorm.DB, id int) (subjects []StudentSubject, err error) {
	err = db.Model(&StudentSubject{}).Where(&StudentSubject{StudentID: id}).Find(&subjects).Error
	return
}

func GetStudentsBySubjectID(db *gorm.DB, id int) (students []StudentSubject, err error) {
	err = db.Model(&StudentSubject{}).Where(&StudentSubject{SubjectID: id}).Find(&students).Error
	return
}

func GetSubjectsAndInfoByStudentID(db *gorm.DB, id int) (subjects []subject.Subject, err error) {
	err = db.Table("student").Select("subject.id, subject.name, subject.description").Joins("JOIN student_subject ON student.id = student_subject.student_id").Joins("JOIN subject ON subject.id = student_subject.subject_id").Where(&StudentSubject{StudentID: id}).Scan(&subjects).Error
	return
}

func GetStudentAndInfoBySubjectID(db *gorm.DB, id int) (infos []Information, err error) {
    err = db.Table("student_subject").Select("student.id, student.first_name, student.last_name, subject.id, subject.name").Joins("JOIN student ON student.id = student_subject.student_id").Joins("JOIN subject ON subject.id = student_subject.subject_id").Where(&StudentSubject{SubjectID: id}).Scan(&infos).Error
    return
}


func GetSubjectAndInfoByStudentIDAndTimeStamp (db *gorm.DB, id int, BeginningSchoolYear int64, EndingSchoolYear int64) (infos []Information, err error) {

	err = db.Table("student").Select("student.id, student.first_name, student.last_name, subject.id, subject.name").
		Joins("JOIN student_subject ON student.id = student_subject.student_id").
		Joins("JOIN subject ON subject.id = student_subject.subject_id").
		Where("student_subject.student_id = ? AND student_subject.date BETWEEN ? AND ?", id, BeginningSchoolYear, EndingSchoolYear).
		Scan(&infos).Error

	return
}