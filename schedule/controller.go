package schedule

import (
	"encoding/json"
	"net/http"
	"strconv"
	"team-academy/component"

	"github.com/gorilla/mux"
)

func DeleteScheduleController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	scheduleID := vars["ID"]

	ID, err := strconv.Atoi(scheduleID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = DeleteSchedule(component.App.DB, ID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

func FetchScheduleController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	scheduleID := vars["ID"]

	ID, err := strconv.Atoi(scheduleID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	schedule, err := GetScheduleByID(component.App.DB, ID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	encodedSchedule, err := json.Marshal(schedule)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(encodedSchedule)
}

func CreateScheduleController(w http.ResponseWriter, r *http.Request) {
	var decodedSchedule Schedule

	err := json.NewDecoder(r.Body).Decode(&decodedSchedule)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = CreateSchedule(component.App.DB, decodedSchedule)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

func UpdateScheduleController(w http.ResponseWriter, r *http.Request) {
	var decodedSchedule Schedule

	err := json.NewDecoder(r.Body).Decode(&decodedSchedule)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = UpdateSchedule(component.App.DB, decodedSchedule)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	encodedSchedule, err := json.Marshal(decodedSchedule)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(encodedSchedule)
}
