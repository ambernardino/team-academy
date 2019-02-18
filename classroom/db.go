package classroom

import (
	"team-academy/component"
	"team-academy/department"

	"github.com/jinzhu/gorm"
)

type Classroom struct {
	ID           int `gorm:"AUTO_INCREMENT"`
	Name         string
	DepartmentID int
}

func CreateTableIfNotExists(db *gorm.DB) (exists bool, err error) {
	if !db.HasTable(Classroom{}) {
		return false, db.CreateTable(Classroom{}).Error
	}
	return true, nil
}

func CreateClassroom(db *gorm.DB, classroom Classroom) (err error) {
	_, err = department.GetDepartmentByID(db, classroom.DepartmentID)
	if err != nil {
		err = component.ErrDepartmentDoesntExist
		return
	}

	return db.Save(&classroom).Error
}

func UpdateClassroom(db *gorm.DB, classroom Classroom) (err error) {
	_, err = GetClassroomByID(db, classroom.ID)
	if err != nil {
		err = component.ErrClassroomDoesntExist
		return
	}

	if classroom.ID <= 0 {
		err = component.ErrMissingParameters
		return
	}

	return db.Model(&Classroom{}).Updates(&classroom).Error
}

func DeleteClassroom(db *gorm.DB, id int) (err error) {
	return db.Delete(&Classroom{ID: id}).Error
}

func GetClassroomByID(db *gorm.DB, id int) (classroom Classroom, err error) {
	err = db.First(&classroom, &Classroom{ID: id}).Error
	return
}

func GetAllClassrooms(db *gorm.DB) (classrooms []Classroom, err error) {
	err = db.Find(&classrooms).Error
	return
}
