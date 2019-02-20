package student_shift

import (
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
