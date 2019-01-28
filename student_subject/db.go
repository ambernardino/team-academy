package student_subject

import (
	"fmt"
	"team-academy/component"

	"github.com/jinzhu/gorm"
)

type StudentSubject struct {
	ID        int `gorm:"AUTO_INCREMENT"`
	StudentID int
	SubjectID int
}

type Information struct {
	StudentID        int
	StudentFirstName string
	StudentLastName  string
	SubjectID        int
	SubjectName      string
}

func CreateTableIfNotExists(db *gorm.DB) (err error) {
	if !db.HasTable(StudentSubject{}) {
		return db.CreateTable(StudentSubject{}).Error
	}
	return
}

func AddStudentToSubject(db *gorm.DB, studentID, subjectID int) (err error) {
	rows, err := db.Table("student").Select("student_subject.student_id, student_subject.subject_id").Joins("JOIN student_subject ON student.id = student_subject.student_id").Joins("JOIN subject ON subject.id = student_subject.subject_id").Where(&StudentSubject{StudentID: studentID, SubjectID: subjectID}).Rows()

	if rows.Next() {
		err = component.ErrSomethingAlreadyExists
		return
	}

	return db.Save(&StudentSubject{StudentID: studentID, SubjectID: subjectID}).Error
}

func RemoveStudentFromSubject(db *gorm.DB, studentID, subjectID int) (err error) {
	return db.Delete(&StudentSubject{StudentID: studentID, SubjectID: subjectID}).Error
}

func GetSubjectsFromStudentID(db *gorm.DB, id int) (subjects []StudentSubject, err error) {
	err = db.Model(&StudentSubject{}).Find(&subjects).Where(&StudentSubject{SubjectID: id}).Error
	return
}

func GetStudentsBySubjectID(db *gorm.DB, id int) (students []StudentSubject, err error) {
	err = db.Model(&StudentSubject{}).Find(&students).Where(&StudentSubject{StudentID: id}).Error
	return
}

func Delete(db *gorm.DB, id int) (err error) {
	return db.Delete(&StudentSubject{ID: id}).Error
}

func IsStudentRegisteredInSubject(db *gorm.DB, studentID, subjectID int) (err error) {
	var studentSubject StudentSubject
	err = db.First(&studentSubject, &StudentSubject{StudentID: studentID, SubjectID: subjectID}).Error
	return
}

func GetSubjectAndInfoByStudentID(db *gorm.DB, id int) (infos []Information, err error) {
	err = db.Table("student").Select("student.id, student.first_name, student.last_name, subject.id, subject.name").Joins("JOIN student_subject ON student.id = student_subject.student_id").Joins("JOIN subject ON subject.id = student_subject.subject_id").Where(&StudentSubject{StudentID: id}).Scan(&infos).Error

	for _, v := range infos {
		fmt.Println(v)
	}
	return
}
