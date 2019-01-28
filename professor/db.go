package professor

import (
	"team-academy/component"
	"time"

	"github.com/jinzhu/gorm"
)

type Professor struct {
	ID         int    `gorm:"AUTO_INCREMENT"`
	FirstName  string `json:"first_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	CursoIDs   string
	CadeiraIDS string
	StartDate  time.Time
}

func CreateTableIfNotExists(db *gorm.DB) (exists bool, err error) {
	if !db.HasTable(Professor{}) {
		return false, db.CreateTable(Professor{}).Error
	}
	return true, nil
}

func CreateProfessor(db *gorm.DB, professor Professor) (err error) {
	return db.Save(&professor).Error
}

func GetAllProfessors(db *gorm.DB) (professors []Professor, err error) {
	err = db.Find(&professors).Error
	return
}

func UpdateProfessorInfo(db *gorm.DB, professor Professor) (err error) {
	_, err = GetProfessorByID(db, professor.ID)
	if err != nil {
		return component.ErrProfessorDoesntExist
	} else if professor.ID <= 0 {
		return component.ErrMissingParameters
	}
	return db.Model(&Professor{}).Updates(&professor).Error
}

func DeleteProfessor(db *gorm.DB, id int) (err error) {
	return db.Delete(&Professor{ID: id}).Error
}

func GetProfessorByID(db *gorm.DB, id int) (prof Professor, err error) {
	err = db.Model(&Professor{}).Where(&Professor{ID: id}).Find(&prof).Error
	return
}
