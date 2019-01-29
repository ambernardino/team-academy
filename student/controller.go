package student

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"team-academy/component"
	"time"
)

func DeleteStudentController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["studentID"]

	ID, err := strconv.Atoi(studentID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = DeleteStudent(component.App.DB, ID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

func GetStudentController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["studentID"]

	ID, err := strconv.Atoi(studentID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	student, err := GetStudentByID(component.App.DB, ID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	encodedStudent, err := json.Marshal(student)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(encodedStudent)
}

func UpdateStudentController(w http.ResponseWriter, r *http.Request) {
	var encodedStudent Student

	err := json.NewDecoder(r.Body).Decode(&encodedStudent)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = UpdateStudent(component.App.DB, encodedStudent)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Fprintf(w, "%v", encodedStudent)
}

func CreateStudentController(w http.ResponseWriter, r *http.Request) {
	var encodedStudent Student

	err := json.NewDecoder(r.Body).Decode(&encodedStudent)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	encodedStudent.StartDate = time.Now().UTC()

	err = CreateStudent(component.App.DB, encodedStudent)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Fprintf(w, "%v", encodedStudent)
}

