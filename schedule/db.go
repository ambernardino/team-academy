package schedule

import (
	"team-academy/component"
	"team-academy/shift"

	"github.com/jinzhu/gorm"
)

type Schedule struct {
	ID        int `gorm:"AUTO_INCREMENT"`
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
	_, err = shift.GetShiftByID(db, schedule.ShiftID)
	if err != nil {
		err = component.ErrShiftDoesntExist
		return
	}

	if schedule.Weekday < 0 || schedule.Weekday > 6 {
		err = component.ErrWeekdayDoesntExist
		return
	}

	// 28800s = 8h, 82800s = 23h
	if schedule.StartTime < 28800 || schedule.StartTime > 82800 {
		err = component.ErrInvalidStartTime
		return
	}

	// 32400s = 9h, 86400s = 24h
	if schedule.EndTime < 32400 || schedule.EndTime > 86400 {
		err = component.ErrInvalidEndTime
		return
	}

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
