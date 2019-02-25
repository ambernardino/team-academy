package student_shift

import (
	"encoding/json"
	"net/http"
	"strconv"
	"team-academy/component"

	"github.com/gorilla/mux"
)

func AddStudentToShiftController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["studentID"]
	shiftID := vars["shiftID"]

	student, err := strconv.Atoi(studentID)
	if component.ControllerError(w, err, nil) {
		return
	}

	shift, err := strconv.Atoi(shiftID)
	if component.ControllerError(w, err, nil) {
		return
	}

	err = AddStudentToShift(component.App.DB, student, shift)
	if component.ControllerError(w, err, nil) {
		return
	}

	component.ReturnResponse(w, "Student added to shift")
}

func RemoveStudentFromShiftController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["studentID"]
	shiftID := vars["shiftID"]

	student, err := strconv.Atoi(studentID)
	if component.ControllerError(w, err, nil) {
		return
	}

	shift, err := strconv.Atoi(shiftID)
	if component.ControllerError(w, err, nil) {
		return
	}

	err = RemoveStudentFromShift(component.App.DB, student, shift)
	if component.ControllerError(w, err, nil) {
		return
	}

	component.ReturnResponse(w, "Student removed from shift")
}

func FetchStudentShiftController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["studentID"]
	shiftID := vars["shiftID"]

	student, err := strconv.Atoi(studentID)
	if component.ControllerError(w, err, nil) {
		return
	}

	shift, err := strconv.Atoi(shiftID)
	if component.ControllerError(w, err, nil) {
		return
	}

	studentShift, err := GetStudentShift(component.App.DB, student, shift)
	if component.ControllerError(w, err, nil) {
		return
	}

	encodedStudentShift, err := json.Marshal(studentShift)
	if component.ControllerError(w, err, nil) {
		return
	}

	w.Write(encodedStudentShift)
	component.ReturnResponse(w, "Fetch Student Shift")
}

func FetchShiftsByStudentIDController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["studentID"]

	student, err := strconv.Atoi(studentID)
	if component.ControllerError(w, err, nil) {
		return
	}

	shifts, err := GetShiftsByStudentID(component.App.DB, student)
	if component.ControllerError(w, err, nil) {
		return
	}

	encodedShifts, err := json.Marshal(shifts)
	if component.ControllerError(w, err, nil) {
		return
	}

	w.Write(encodedShifts)
	component.ReturnResponse(w, "Fetch List of Student's Shifts")
}
