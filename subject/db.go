package subject

import (

	//"team-academy/subject"
	"team-academy/component"

	"github.com/jinzhu/gorm"
)

type Subject struct {
	ID          int `gorm:"AUTO_INCREMENT"`
	Name        string
	Description string
}

func CreateTableIfNotExists(db *gorm.DB) (exists bool, err error) {
	if !db.HasTable(Subject{}) {
		return false, db.CreateTable(Subject{}).Error
	}
	return true, nil
}

func CreateSubject(db *gorm.DB, newSubject Subject) (err error) {
	_, err = GetSubjectByID(db, newSubject.ID)
	if err != nil {
		return db.Save(&newSubject).Error
	}
	_, err = GetSubjectByName(db, newSubject.Name)
	if err != nil {
		return db.Save(&newSubject).Error
	}
	err = component.ErrSomethingAlreadyExists
	return
}

func DeleteSubject(db *gorm.DB, id int) (err error) {
	return db.Delete(&Subject{ID: id}).Error
}

func UpdateSubjectInfo(db *gorm.DB, subject Subject) (err error) {
	return db.Model(&Subject{}).Updates(&subject).Error
}

func GetAllSubjects(db *gorm.DB) (subjects []Subject, err error) {
	err = db.Find(&subjects).Error
	return
}

func GetSubjectByID(db *gorm.DB, id int) (subject Subject, err error) {
	err = db.First(&subject, &Subject{ID: id}).Error
	return
}

func GetSubjectByName(db *gorm.DB, name string) (subject Subject, err error) {
	err = db.First(&subject, &Subject{Name: name}).Error
	return
}
