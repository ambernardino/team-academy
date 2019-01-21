package professor

import (
	//"team-academy/professor"
	"time"
	"github.com/jinzhu/gorm"
)

type Professor struct {
	ID         int `gorm:"AUTO_INCREMENT"`
	FirstName  string
	LastName   string
	CursoIds   string
	CadeiraIds string
	StartDate  time.Time
}

func CreateTableIfNotExist(db *gorm.DB) (err error) {
	if !db.HasTable(Professor{}) {
		return db.CreateTable(Professor{}).Error
	}
	return
}

func CreateProfessors(db *gorm.DB) (err error) {
	newProfessor := Professor{FirstName: "Paulo", LastName: "Montezuma", CursoIds: "MIEEC", CadeiraIds: "IT", StartDate: time.Now()}
	err = db.Save(&newProfessor).Error
	return
}

func GetAllProfessors(db *gorm.DB) (professors []Professor, err error) {
	err = db.Find(&professors).Error
	return
}

func UpdateProfessor (db *gorm.DB, professor Professor) (err error) {
	return db.Model(&Professor{}).Update(&professor).Error
}

func RemoveProfessor (db *gorm.DB, id int) (err error) {
	return db.Delete(&Professor{ID: id}).Error
}