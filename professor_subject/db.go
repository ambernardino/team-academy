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
	Date        int64
}


type Information struct {
	ProfessorID        int    `gorm:"column:id"`
	ProfessorFirstName string `gorm:"column:first_name"`
	ProfessorLastName  string `gorm:"column:last_name"`
	SubjectID        int    `gorm:"column:id"`
	SubjectName      string `gorm:"column:name"`
}

func CreateTableIfNotExists(db *gorm.DB) (exists bool, err error) {
	if !db.HasTable(ProfessorSubject{}) {
		return false, db.CreateTable(ProfessorSubject{}).Error
	}
	return true, nil
}

func AddProfessorToSubject(db *gorm.DB, professorID, subjectID int, date int64) (err error) {
	rows, err := db.Table("professor").Select("professor_subject.professor_id, professor_subject.subject_id").Joins("JOIN professor_subject ON professor.id = professor_subject.professor_id").Joins("JOIN subject ON subject.id = professor_subject.subject_id").Where(&ProfessorSubject{ProfessorID: professorID, SubjectID: subjectID}).Rows()

	if rows.Next() {
		err = component.ErrProfessorAlreadyInSubject
		return
	}

	return db.Save(&ProfessorSubject{ProfessorID: professorID, SubjectID: subjectID, Date: date}).Error
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
	err = db.Table("professor").Select("subject.id, subject.name, subject.description").
	Joins("JOIN professor_subject ON professor.id = professor_subject.professor_id").
	Joins("JOIN subject ON subject.id = professor_subject.subject_id").
	Where(&ProfessorSubject{ProfessorID: id}).
	Scan(&subjects).Error
	return
}


func GetSubjectAndInfobyProfessorIDAndTimeStamp (db *gorm.DB, id int, BeginningSchoolYear int64, EndingSchoolYear int64) (infos []Information, err error) {
	err = db.Table("professor").Select("professor.id, professor.first_name, professor.last_name, subject.id, subject.name").
	Joins("JOIN professor_subject ON professor.id = professor_subject.professor_id").
	Joins("JOIN subject ON subject.id = professor_subject.subject_id").
	Where("professor_subject.professor_id = ? AND professor_subject.date BETWEEN ? AND ?", id, BeginningSchoolYear, EndingSchoolYear).
	Scan(&infos).Error

	return
}