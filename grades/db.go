package grades

import (
	"fmt"
	"team-academy/student"
	"team-academy/student_subject"
	"team-academy/subject"

	"github.com/jinzhu/gorm"
)

type Grade struct {
	SubjectID int
	StudentID int
	Rank      string
}

func GiveGrade(db *gorm.DB, grade Grade) (err error) {
	_, err = subject.GetSubjectByID(db, grade.SubjectID)
	if err != nil {
		return
	}
	_, err = student.GetStudentByID(db, grade.StudentID)
	if err != nil {
		return
	}
	st, err := student_subject.GetStudentsBySubjectID(db, grade.SubjectID)
	if err != nil {
		return
	}
	isRegistered := false
	for _, v := range st {
		if v.ID == grade.StudentID {
			isRegistered = true
		}
	}
	if !isRegistered {
		fmt.Println("Student not registered in this class")
		return
	}
	return db.Save(&grade).Error
}

func UpdateGrade(db *gorm.DB, grade Grade) (err error) {

}
