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
	db.SingularTable(true)
	if !db.HasTable(Professor{}) {
		return db.CreateTable(Professor{}).Error
	}
	return
}

func CreateProfessor(db *gorm.DB) (err error) {
	newProfessor := Professor{FirstName: "Paulo", LastName: "Pinto", CursoIDs: "MIEEC", CadeiraIDS: "PM", StartDate: time.Now()}
	return db.Save(&newProfessor).Error
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
