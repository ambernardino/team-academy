package subject

import (
	//"team-academy/subject"
	"github.com/jinzhu/gorm"
)

type Subject struct {
	ID          int `gorm:"AUTO_INCREMENT"`
	Name        string
	Description string
}

func CreateTableIfNotExist(db *gorm.DB) (err error) {
	if !db.HasTable(Subject{}) {
		return db.CreateTable(Subject{}).Error
	}
	return
}

func CreateSubject(db *gorm.DB, newSubject Subject) (err error) {
	err = db.Save(&newSubject).Error
	return
}

func RemoveSubject(db *gorm.DB, id int) (err error) {
	return db.Delete(&Subject{ID: id}).Error
}

func GetAllSubjects(db *gorm.DB) (subjects []Subject, err error) {
	err = db.Find(&subjects).Error
	return
}

func UpdateSubject (db *gorm.DB, subject Subject) (err error) {
	return db.Model(&Subject{}).Update(&subject).Error
}