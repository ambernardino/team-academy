package student

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Student struct {
	ID        int `gorm:"AUTO_INCREMENT"`
	FirstName string
	LastName  string
	DegreeID  int
	StartDate time.Time
}

// Creates a new table if non is found
func CreateTable(db *gorm.DB) (err error) {
	if !db.HasTable(Student{}) {
		return db.CreateTable(Student{}).Error
	}
	return
}

func CreateStudent(db *gorm.DB, newStudent Student) (err error) {
	err = db.Save(&newStudent).Error
	return
}

func GetStudents(db *gorm.DB) (students []Student, err error) {
	err = db.Find(&students).Error
	return
}

func GetStudentByID(db *gorm.DB, id int) (student Student, err error) {
	err = db.Find(&Student{ID: id}).Error
	return
}

func UpdateStudent(db *gorm.DB, updatedStudent Student) (err error) {
	return db.Model(&Student{}).Update(updatedStudent).Error
}

func DeleteStudent(db *gorm.DB, id int) (err error) {
	return db.Delete(&Student{ID: id}).Error
}
