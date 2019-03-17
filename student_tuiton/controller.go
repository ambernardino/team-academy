package student_tuiton

import (
	"encoding/json"
	"net/http"
	"strconv"
	"team-academy/component"

	"github.com/gorilla/mux"
)

func FetchStudentTuitionByStudentIDController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["studentID"]

	ID, err := strconv.Atoi(studentID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	st, err := GetStudentTuitionByStudentID(component.App.DB, ID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	encodedSt, err := json.Marshal(st)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(encodedSt)
}

func AddStudentTuitionController(w http.ResponseWriter, r *http.Request) {
	var decodedSt StudentTuition

	err := json.NewDecoder(r.Body).Decode(&decodedSt)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = AddStudentTuition(component.App.DB, decodedSt)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

func UpdateStudentTuitionController(w http.ResponseWriter, r *http.Request) {
	var decodedSt StudentTuition

	err := json.NewDecoder(r.Body).Decode(&decodedSt)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = UpdateStudentTuition(component.App.DB, decodedSt)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	encodedSt, err := json.Marshal(decodedSt)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(encodedSt)
}

func RemoveStudentTuitionController (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["studentID"]

	ID, err := strconv.Atoi(studentID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = DeleteStudentTuition(component.App.DB, ID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}
