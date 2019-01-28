package student

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Student struct {
	ID        int `gorm:"AUTO_INCREMENT"`
	FirstName string
	LastName  string
	CursoID   int
	StartDate time.Time
}

func CreateTableIfNotExists(db *gorm.DB) (exists bool, err error) {
	if !db.HasTable(Student{}) {
		return false, db.CreateTable(Student{}).Error
	}
	return true, nil
}

func CreateStudent(db *gorm.DB, student Student) (err error) {
	return db.Save(&student).Error
}

func UpdateStudent(db *gorm.DB, student Student) (err error) {
	return db.Model(&Student{}).Update(student).Error
}

func DeleteStudent(db *gorm.DB, id int) (err error) {
	return db.Delete(&Student{ID: id}).Error
}

func GetStudentByID(db *gorm.DB, id int) (student Student, err error) {
	err = db.First(&student, &Student{ID: id}).Error
	return
}

func GetAllStudents(db *gorm.DB) (students []Student, err error) {
	err = db.Find(&students).Error
	return
}
