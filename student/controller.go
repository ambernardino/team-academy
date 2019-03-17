package student

import (
	"encoding/json"
	"net/http"
	"strconv"
	"team-academy/component"
	"time"

	"github.com/gorilla/mux"
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

func FetchStudentController(w http.ResponseWriter, r *http.Request) {
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

func FetchStudentByEmailController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	student, err := GetStudentByEmail(component.App.DB, email)
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

func FetchAllStudentsController(w http.ResponseWriter, r *http.Request) {
	students, err := GetAllStudents(component.App.DB)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	encodedStudents, err := json.Marshal(students)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(encodedStudents)
}

func UpdateStudentController(w http.ResponseWriter, r *http.Request) {
	var decodedStudent Student

	err := json.NewDecoder(r.Body).Decode(&decodedStudent)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = UpdateStudent(component.App.DB, decodedStudent)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	encodedStudent, err := json.Marshal(decodedStudent)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(encodedStudent)
}

func CreateStudentController(w http.ResponseWriter, r *http.Request) {
	var decodedStudent Student

	err := json.NewDecoder(r.Body).Decode(&decodedStudent)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	decodedStudent.StartDate = time.Now().UTC().Unix()

	err = CreateStudent(component.App.DB, decodedStudent)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}


