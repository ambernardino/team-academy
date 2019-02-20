package shift

import (
	"encoding/json"
	"net/http"
	"strconv"
	"team-academy/component"

	"github.com/gorilla/mux"
)

func DeleteShiftController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shiftID := vars["shiftID"]

	ID, err := strconv.Atoi(shiftID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = DeleteShift(component.App.DB, ID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

func FetchShiftController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shiftID := vars["ID"]

	ID, err := strconv.Atoi(shiftID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	shift, err := GetShiftByID(component.App.DB, ID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	encodedShift, err := json.Marshal(shift)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(encodedShift)
}

func CreateShiftController(w http.ResponseWriter, r *http.Request) {
	var decodedShift Shift

	err := json.NewDecoder(r.Body).Decode(&decodedShift)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = CreateShift(component.App.DB, decodedShift)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

func UpdateShiftController(w http.ResponseWriter, r *http.Request) {
	var decodedShift Shift

	err := json.NewDecoder(r.Body).Decode(&decodedShift)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = UpdateShift(component.App.DB, decodedShift)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	encodedShift, err := json.Marshal(decodedShift)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(encodedShift)
}
