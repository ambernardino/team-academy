package department

import (
	"team-academy/component"

	"github.com/jinzhu/gorm"
)

type Department struct {
	ID   int `gorm:"AUTO_INCREMENT"`
	Name string
}

func CreateTableIfNotExists(db *gorm.DB) (exists bool, err error) {
	if !db.HasTable(Department{}) {
		return false, db.CreateTable(Department{}).Error
	}

	return true, nil
}

func CreateDepartment(db *gorm.DB, department Department) (err error) {
	return db.Save(&department).Error
}

func UpdateDepartment(db *gorm.DB, department Department) (err error) {
	_, err = GetDepartmentByID(db, department.ID)
	if err != nil {
		err = component.ErrDepartmentDoesntExist
		return
	}

	if department.ID <= 0 {
		err = component.ErrMissingParameters
		return
	}

	return db.Model(&Department{}).Updates(&department).Error
}

func GetDepartmentByID(db *gorm.DB, id int) (department Department, err error) {
	err = db.First(&department, &Department{ID: id}).Error
	return
}

func GetAllDepartments(db *gorm.DB) (departments []Department, err error) {
	err = db.Find(&departments).Error
	return
}

func DeleteDepartment(db *gorm.DB, id int) (err error) {
	return db.Delete(&Department{ID: id}).Error
}
