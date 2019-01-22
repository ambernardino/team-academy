package grades

import (
	"team-academy/subject"
	"github.com/jinzhu/gorm"
)

type Grade struct {
	SubjectID int
	StudentID int
	Rank string
}

func GiveGrade (db *gorm.DB, grade Grade) (e error) {
	err = subject.GetSubjectByID (db, grade.SubjectID)
	if err != nil {
		
	}
	return db.Save(&grade).Error
}

func UpdateGrade (db *gorm.DB, grade Grade) (err error) {

}
