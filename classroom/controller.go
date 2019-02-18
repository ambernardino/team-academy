package classroom

import (
	"encoding/json"
	"net/http"
	"strconv"
	"team-academy/component"

	"github.com/gorilla/mux"
)

func CreateClassroomController(w http.ResponseWriter, r *http.Request) {
	var classroom Classroom
	err := json.NewDecoder(r.Body).Decode(&classroom)
	if component.ControllerError(w, err, component.ErrUnmarshallingJSON) {
		return
	}

	err = CreateClassroom(component.App.DB, classroom)
	if component.ControllerError(w, err, nil) {
		return
	}

	newClassroom, err := GetClassroomByID(component.App.DB, classroom.ID)
	if component.ControllerError(w, err, nil) {
		return
	}

	component.ReturnResponse(w, newClassroom)
}

func UpdateClassroomController(w http.ResponseWriter, r *http.Request) {
	var classroom Classroom

	err := json.NewDecoder(r.Body).Decode(&classroom)
	if component.ControllerError(w, err, nil) {
		return
	}

	err = UpdateClassroom(component.App.DB, classroom)
	if component.ControllerError(w, err, nil) {
		return
	}

	component.ReturnResponse(w, classroom)
}

func FetchClassroomByIDController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]

	classroomID, err := strconv.Atoi(id)
	if component.ControllerError(w, err, nil) {
		return
	}

	classroom, err := GetClassroomByID(component.App.DB, classroomID)
	if component.ControllerError(w, err, nil) {
		return
	}

	component.ReturnResponse(w, classroom)
}

func FetchAllClassroomsController(w http.ResponseWriter, r *http.Request) {
	classrooms, err := GetAllClassrooms(component.App.DB)
	if component.ControllerError(w, err, nil) {
		return
	}

	component.ReturnResponse(w, classrooms)
}

func RemoveClassroomController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["ID"]

	classroomID, err := strconv.Atoi(ID)
	if component.ControllerError(w, err, nil) {
		return
	}

	err = DeleteClassroom(component.App.DB, classroomID)
	if component.ControllerError(w, err, nil) {
		return
	}

	component.ReturnResponse(w, "Classroom deleted")
}
