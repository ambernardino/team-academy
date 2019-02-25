package professor_shift

import (
	"encoding/json"
	"net/http"
	"strconv"
	"team-academy/component"

	"github.com/gorilla/mux"
)

func AddProfessorToShiftController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	professorID := vars["professorID"]
	shiftID := vars["shiftID"]

	professor, err := strconv.Atoi(professorID)
	if component.ControllerError(w, err, nil) {
		return
	}

	shift, err := strconv.Atoi(shiftID)
	if component.ControllerError(w, err, nil) {
		return
	}

	err = AddProfessorToShift(component.App.DB, professor, shift)
	if component.ControllerError(w, err, nil) {
		return
	}

	component.ReturnResponse(w, "Professor added to shift")
}

func RemoveProfessorFromShiftController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	professorID := vars["professorID"]
	shiftID := vars["shiftID"]

	professor, err := strconv.Atoi(professorID)
	if component.ControllerError(w, err, nil) {
		return
	}

	shift, err := strconv.Atoi(shiftID)
	if component.ControllerError(w, err, nil) {
		return
	}

	err = RemoveProfessorFromShift(component.App.DB, professor, shift)
	if component.ControllerError(w, err, nil) {
		return
	}

	component.ReturnResponse(w, "Professor removed from shift")
}

func FetchProfessorShiftController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	professorID := vars["professorID"]
	shiftID := vars["shiftID"]

	professor, err := strconv.Atoi(professorID)
	if component.ControllerError(w, err, nil) {
		return
	}

	shift, err := strconv.Atoi(shiftID)
	if component.ControllerError(w, err, nil) {
		return
	}

	professorShift, err := GetProfessorShift(component.App.DB, professor, shift)
	if component.ControllerError(w, err, nil) {
		return
	}

	encodedProfShift, err := json.Marshal(professorShift)
	if component.ControllerError(w, err, nil) {
		return
	}

	w.Write(encodedProfShift)
	component.ReturnResponse(w, "Fetch Professor Shift")
}

func FetchShiftsByProfessorIDController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	professorID := vars["professorID"]

	professor, err := strconv.Atoi(professorID)
	if component.ControllerError(w, err, nil) {
		return
	}

	shifts, err := GetShiftsByProfessorID(component.App.DB, professor)
	if component.ControllerError(w, err, nil) {
		return
	}

	encodedShifts, err := json.Marshal(shifts)
	if component.ControllerError(w, err, nil) {
		return
	}

	w.Write(encodedShifts)
	component.ReturnResponse(w, "Fetch List of Professor's Shifts")
}