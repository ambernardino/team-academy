package professor

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Professor struct {
	ID         int `gorm:"AUTO_INCREMENT"`
	FirstName  string
	LastName   string
	CursoIDs   string
	CadeiraIDS string
	StartDate  time.Time
}

func CreateTableIfNotExists(db *gorm.DB) (err error) {
	if !db.HasTable(Professor{}) {
		return db.CreateTable(Professor{}).Error
	}
	return
}

func CreateProfessor(db *gorm.DB, professor Professor) (err error) {
	return db.Save(&professor).Error
}

func GetAllProfessors(db *gorm.DB) (professors []Professor, err error) {
	err = db.Find(&professors).Error
	return
}

func UpdateProfessorInfo(db *gorm.DB, professor Professor) (err error) {
	return db.Model(&Professor{}).Updates(&professor).Error
}

func DeleteProfessor(db *gorm.DB, id int) (err error) {
	return db.Delete(&Professor{ID: id}).Error
}

func GetProfessorByID(db *gorm.DB, id int) (professor Professor, err error) {
	err = db.Model(&Professor{}).Where(&Professor{ID: id}).Find(&professor).Error
	return
}
