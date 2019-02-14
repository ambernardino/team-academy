package student_subject

import (
	"encoding/json"
	"net/http"
	"strconv"
	"team-academy/component"
	"time"

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

	info, err := GetSubjectAndInfoByStudentID(component.App.DB, ID)
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

	err = AddStudentToSubject(component.App.DB, studID, subjID, int64(time.Now().UTC().Unix()))
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

func FetchSubjectAndInfoByStudentIDAndTimeStampController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["studentID"]
	beginSchool := vars["beginSchool"]
	endSchool := vars["endSchool"]

	studID, err := strconv.Atoi(studentID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	layout := "01/02/2006 3:04:05 PM"
	beginSchool = "09/01/" + beginSchool + " 0:00:00 AM"
	endSchool = "08/31/" + endSchool + " 0:00:00 AM"

	beginTime, err := time.Parse(layout, beginSchool)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	endTime, err := time.Parse(layout, endSchool)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	timeStart := beginTime.UTC().Unix()
	timeEnd := endTime.UTC().Unix()

	info, err := GetSubjectAndInfoByStudentIDAndTimeStamp(component.App.DB, studID, timeStart, timeEnd)
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
