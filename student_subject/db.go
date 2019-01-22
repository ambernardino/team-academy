package student_subject

import (
	"team-academy/component"
	"team-academy/student"
	"team-academy/subject"

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
	_, err = subject.GetSubjectByID(db, subjectID)
	if err != nil {
		return
	}

	_, err = student.GetStudentByID(db, studentID)
	if err != nil {
		return
	}

	err = IsStudentRegisteredInSubject(db, studentID, subjectID)
	if err == nil {
		return component.ErrSomethingAlreadyExists
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

func IsStudentRegisteredInSubject(db *gorm.DB, studentID, subjectID int) (err error) {
	var studentSubject StudentSubject
	err = db.First(&studentSubject, &StudentSubject{StudentID: studentID, SubjectID: subjectID}).Error
	return
}
