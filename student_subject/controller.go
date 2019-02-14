package student_subject

import (
	"encoding/json"
	"net/http"
	"strconv"
	"team-academy/component"

	"github.com/gorilla/mux"
)

func FetchSubjectsByStudentIDController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["studentID"]

	ID, err := strconv.Atoi(studentID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	subjects, err := GetSubjectsByStudentID(component.App.DB, ID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	encodedSubjects, err := json.Marshal(subjects)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(encodedSubjects)
}

func FetchStudentsBySubjectIDController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subjectID := vars["subjectID"]

	ID, err := strconv.Atoi(subjectID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	students, err := GetStudentsBySubjectID(component.App.DB, ID)
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

func FetchStudentAndInfoBySubjectIDController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subjectID := vars["subjectID"]

	ID, err := strconv.Atoi(subjectID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	info, err := GetStudentAndInfoBySubjectID(component.App.DB, ID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	encodedInfo, err := json.Marshal(info)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(encodedInfo)
}

func FetchSubjectAndInfoByStudentIDController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["studentID"]

	ID, err := strconv.Atoi(studentID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	info, err := GetSubjectsAndInfoByStudentID(component.App.DB, ID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	encodedInfo, err := json.Marshal(info)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(encodedInfo)
}

func AddStudentToSubjectController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subjectID := vars["subjectID"]
	studentID := vars["studentID"]

	subjID, err := strconv.Atoi(subjectID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	studID, err := strconv.Atoi(studentID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = AddStudentToSubject(component.App.DB, studID, subjID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

func RemoveStudentFromSubjectController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subjectID := vars["subjectID"]
	studentID := vars["studentID"]

	subjID, err := strconv.Atoi(subjectID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	studID, err := strconv.Atoi(studentID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = RemoveStudentFromSubject(component.App.DB, studID, subjID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}
