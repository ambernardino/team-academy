package professor_shift

import (
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
