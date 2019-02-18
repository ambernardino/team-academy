package department

import (
	"encoding/json"
	"net/http"
	"strconv"
	"team-academy/component"

	"github.com/gorilla/mux"
)

func CreateDepartmentController(w http.ResponseWriter, r *http.Request) {
	var department Department
	err := json.NewDecoder(r.Body).Decode(&department)
	if component.ControllerError(w, err, component.ErrUnmarshallingJSON) {
		return
	}

	err = CreateDepartment(component.App.DB, department)
	if component.ControllerError(w, err, nil) {
		return
	}

	newDepartment, err := GetDepartmentByID(component.App.DB, department.ID)
	if component.ControllerError(w, err, nil) {
		return
	}

	component.ReturnResponse(w, newDepartment)
}

func UpdateDepartmentController(w http.ResponseWriter, r *http.Request) {
	var department Department

	err := json.NewDecoder(r.Body).Decode(&department)
	if component.ControllerError(w, err, nil) {
		return
	}

	err = UpdateDepartment(component.App.DB, department)
	if component.ControllerError(w, err, nil) {
		return
	}

	component.ReturnResponse(w, department)
}

func FetchDepartmentByIDController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]

	departmentID, err := strconv.Atoi(id)
	if component.ControllerError(w, err, nil) {
		return
	}

	department, err := GetDepartmentByID(component.App.DB, departmentID)
	if component.ControllerError(w, err, nil) {
		return
	}

	component.ReturnResponse(w, department)
}

func FetchAllDepartmentsController(w http.ResponseWriter, r *http.Request) {
	departments, err := GetAllDepartments(component.App.DB)
	if component.ControllerError(w, err, nil) {
		return
	}

	component.ReturnResponse(w, departments)
}

func RemoveDepartmentController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["ID"]

	departmentID, err := strconv.Atoi(ID)
	if component.ControllerError(w, err, nil) {
		return
	}

	err = DeleteDepartment(component.App.DB, departmentID)
	if component.ControllerError(w, err, nil) {
		return
	}

	component.ReturnResponse(w, "Department deleted")
}
