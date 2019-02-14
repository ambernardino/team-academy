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

func CreateTableIfNotExists(db *gorm.DB) (exists bool, err error) {
	if !db.HasTable(ProfessorSubject{}) {
		return false, db.CreateTable(ProfessorSubject{}).Error
	}
	return true, nil
}

func AddProfessorToSubject(db *gorm.DB, professorID, subjectID int, date int64) (err error) {
	
	_, err = subject.GetSubjectByID(db, subjectID)
	if err != nil {
		return
	}

	rows, err := db.Table("professor_subject").Select("professor_subject.professor_id, professor_subject.subject_id, professor.first_name, professor.last_name, subject.name").Joins("JOIN professor ON professor.id = professor_subject.professor_id").Joins("JOIN subject ON subject.id = professor_subject.subject_id").Where(&ProfessorSubject{ProfessorID: professorID, SubjectID: subjectID}).Rows()
	if err != nil {
		return
	}

	defer rows.Close()
	if rows.Next() {
		err = component.ErrProfessorAlreadyInSubject
		return
	}

	return db.Save(&ProfessorSubject{ProfessorID: professorID, SubjectID: subjectID, Date: date}).Error
}

func RemoveProfessorFromSubject(db *gorm.DB, professorID, subjectID int) (err error) {
	return db.Where(&ProfessorSubject{ProfessorID: professorID, SubjectID: subjectID}).Delete(&ProfessorSubject{}).Error
}

func GetProfessorSubject(db *gorm.DB, professorID, subjectID int) (professorSubject ProfessorSubject, err error) {
	err = db.First(&professorSubject, &ProfessorSubject{ProfessorID: professorID, SubjectID: subjectID}).Error
	return
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


func GetSubjectAndInfobyProfessorIDAndTimeStamp (db *gorm.DB, id int, BeginningSchoolYear int64, EndingSchoolYear int64) (subjects []subject.Subject, err error) {
	err = db.Table("professor").Select("subject.id, subject.name, subject.description").
	Joins("JOIN professor_subject ON professor.id = professor_subject.professor_id").
	Joins("JOIN subject ON subject.id = professor_subject.subject_id").
	Where("professor_subject.professor_id = ? AND professor_subject.date BETWEEN ? AND ?", id, BeginningSchoolYear, EndingSchoolYear).
	Scan(&subjects).Error

	return
}