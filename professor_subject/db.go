package professor_subject

import (
	"team-academy/component"
	"team-academy/subject"

	"github.com/jinzhu/gorm"
)

type ProfessorSubject struct {
	ID          int `gorm:"AUTO_INCREMENT"`
	ProfessorID int
	SubjectID   int
}

func CreateTableIfNotExists(db *gorm.DB) (exists bool, err error) {
	if !db.HasTable(ProfessorSubject{}) {
		return false, db.CreateTable(ProfessorSubject{}).Error
	}
	return true, nil
}

func AddProfessorToSubject(db *gorm.DB, professorID, subjectID int) (err error) {
	rows, err := db.Table("professor").Select("professor_subject.professor_id, professor_subject.subject_id").Joins("JOIN professor_subject ON professor.id = professor_subject.professor_id").Joins("JOIN subject ON subject.id = professor_subject.subject_id").Where(&ProfessorSubject{ProfessorID: professorID, SubjectID: subjectID}).Rows()

	if rows.Next() {
		err = component.ErrProfessorAlreadyInSubject
		return
	}

	return db.Save(&ProfessorSubject{ProfessorID: professorID, SubjectID: subjectID}).Error
}

func GetProfessorsBySubjectID(db *gorm.DB, id int) (professors []ProfessorSubject, err error) {
	err = db.Model(&ProfessorSubject{}).Where(&ProfessorSubject{SubjectID: id}).Find(&professors).Error
	return
}

func GetSubjectsByProfessorID(db *gorm.DB, id int) (subjects []ProfessorSubject, err error) {
	err = db.Model(&ProfessorSubject{}).Where(&ProfessorSubject{ProfessorID: id}).Find(&subjects).Error
	return
}

func GetSubjectsAndInfoByProfessorID(db *gorm.DB, id int) (subjects []subject.Subject, err error) {
	err = db.Table("professor").Select("subject.id, subject.name, subject.description").Joins("JOIN professor_subject ON professor.id = professor_subject.professor_id").Joins("JOIN subject ON subject.id = professor_subject.subject_id").Where(&ProfessorSubject{ProfessorID: id}).Scan(&subjects).Error
	return
}
