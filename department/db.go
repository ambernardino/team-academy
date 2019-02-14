package department

import "github.com/jinzhu/gorm"

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

func GetDepartmentByID(db *gorm.DB, id int) (department Department, err error) {
	err = db.First(&department, &Department{ID: id}).Error
	return
}

func DeleteDepartment(db *gorm.DB, id int) (err error) {
	return db.Delete(&Department{ID: id}).Error
}
