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

func CreateTable(db *gorm.DB) (err error) {
	if !db.HasTable(StudentSubject{}) {
		return db.CreateTable(StudentSubject{}).Error
	}
	return
}

func Add(db *gorm.DB, studentSubject StudentSubject) (err error) {
	registeredStudentSubject, err := FindStudentSubject(db, studentSubject.StudentID, studentSubject.SubjectID)

	if registeredStudentSubject.StudentID == studentSubject.StudentID &&
		registeredStudentSubject.SubjectID == studentSubject.SubjectID {
		return component.ErrSomethingAlreadyExists
	}

	newSubject := StudentSubject{StudentID: studentSubject.StudentID, SubjectID: studentSubject.SubjectID}
	err = db.Save(&newSubject).Error
	return
}

func Remove(db *gorm.DB, id int) (err error) {
	return db.Debug().Delete(&StudentSubject{ID: id}).Error
}

// Needs to be rewritten
func GetStudent(db *gorm.DB, studentID int) (err error) {
	return db.First(&StudentSubject{StudentID: studentID}).Error
}

// Needs to be rewritten
func GetStudentsInSubject(db *gorm.DB, subjectID int) (studentSubjectArray []StudentSubject, err error) {
	err = db.Model(&StudentSubject{}).Find(&studentSubjectArray).Where(&StudentSubject{SubjectID: subjectID}).Error
	return
}

func FindStudentSubject(db *gorm.DB, studentID, subjectID int) (studentSubject StudentSubject, err error) {
	err = db.Debug().First(&studentSubject, StudentSubject{StudentID: studentID, SubjectID: subjectID}).Error
	return
}
