package subject

import (
	"github.com/jinzhu/gorm"
)

type Subject struct {
	ID          int `gorm:"AUTO_INCREMENT"`
	Name        string
	Description string
}

func CreateTableIfNotExists(db *gorm.DB) (err error) {
	if !db.HasTable(Subject{}) {
		return db.CreateTable(Subject{}).Error
	}
	return
}

func CreateSubject(db *gorm.DB, newSubject Subject) (err error) {
	return db.Save(&newSubject).Error
}

func DeleteSubject(db *gorm.DB, id int) (err error) {
	return db.Delete(&Subject{ID: id}).Error
}

func UpdateSubjectInfo(db *gorm.DB, subject Subject) (err error) {
	return db.Model(&Subject{}).Updates(&subject).Error
}

func GetSubjectByID(db *gorm.DB, id int) (subject Subject, err error) {
	err = db.First(&subject, &Subject{ID: id}).Error
	return
}

func GetAllSubjects(db *gorm.DB) (subjects []Subject, err error) {
	err = db.Find(&subjects).Error
	return
}
