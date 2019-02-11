package subject

import (
	"encoding/json"
	"net/http"
	"strconv"
	"team-academy/component"

	"github.com/gorilla/mux"
)

func CreateSubjectController(w http.ResponseWriter, r *http.Request) {
	var subject Subject
	err := json.NewDecoder(r.Body).Decode(&subject)
	if component.ControllerError(w, err, component.ErrUnmarshallingJSON) {
		return
	}

	err = CreateSubject(component.App.DB, subject)
	if component.ControllerError(w, err, nil) {
		return
	}

	newSubject, err := GetSubjectByName(component.App.DB, subject.Name)
	if component.ControllerError(w, err, nil) {
		return
	}

	component.ReturnResponse(w, newSubject)
}

func FetchAllSubjectsController(w http.ResponseWriter, r *http.Request) {
	subjects, err := GetAllSubjects(component.App.DB)
	if component.ControllerError(w, err, nil) {
		return
	}

	component.ReturnResponse(w, subjects)
}

func FetchSubjectByIDController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]

	subjectID, err := strconv.Atoi(id)
	if component.ControllerError(w, err, nil) {
		return
	}

	subject, err := GetSubjectByID(component.App.DB, subjectID)
	if component.ControllerError(w, err, nil) {
		return
	}

	component.ReturnResponse(w, subject)
}
