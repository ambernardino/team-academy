package schedule

import (
	"team-academy/component"

	"github.com/jinzhu/gorm"
)

type Schedule struct {
	ID        int `gorm:"AUTO_INCREMENT"`
	SubjectID int
	ShiftID   int
	Weekday   int
	StartTime int64
	EndTime   int64
}

func CreateTableIfNotExists(db *gorm.DB) (exists bool, err error) {
	if !db.HasTable(Schedule{}) {
		return false, db.CreateTable(Schedule{}).Error
	}
	return true, nil
}

func CreateSchedule(db *gorm.DB, schedule Schedule) (err error) {
	return db.Save(&schedule).Error
}

func UpdateSchedule(db *gorm.DB, schedule Schedule) (err error) {
	if schedule.ID <= 0 {
		err = component.ErrMissingParameters
		return
	}

	_, err = GetScheduleByID(db, schedule.ID)
	if err != nil {
		return
	}

	return db.Model(&Schedule{}).Update(schedule).Error
}

func DeleteSchedule(db *gorm.DB, id int) (err error) {
	return db.Delete(&Schedule{ID: id}).Error
}

func GetScheduleByID(db *gorm.DB, id int) (schedule Schedule, err error) {
	err = db.First(&schedule, &Schedule{ID: id}).Error
	return
}
