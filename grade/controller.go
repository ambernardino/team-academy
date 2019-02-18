package grade

import (
	"encoding/json"
	"net/http"
	"strconv"
	"team-academy/component"
	"time"

	"github.com/gorilla/mux"
)

func CreateGradeController(w http.ResponseWriter, r *http.Request) {
	var g Grade
	err := json.NewDecoder(r.Body).Decode(&g)
	if err != nil {
		return
	}
	err = GiveGrade(component.App.DB, g)
	if err != nil {
		return
	}
	w.Write([]byte("Grade Added"))
}

func UpdateGradeController(w http.ResponseWriter, r *http.Request) {
	var g Grade
	err := json.NewDecoder(r.Body).Decode(&g)
	if err != nil {
		return
	}
	err = UpdateGrade(component.App.DB, g)
	if err != nil {
		return
	}
	w.Write([]byte("Grade Updated"))
}

func FetchGradeByStudentController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["ID"]
	id, err := strconv.Atoi(ID)
	if err != nil {
		return
	}
	g, err := GetGradeByStudentID(component.App.DB, id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	list, err := json.Marshal(g)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(list)
}

func FetchGradeBySubjectController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["ID"]
	id, err := strconv.Atoi(ID)
	if err != nil {
		return
	}
	g, err := GetGradeBySubjectID(component.App.DB, id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	list, err := json.Marshal(g)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(list)
}

func FetchStudentsGradesController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["studentID"]
	id, err := strconv.Atoi(studentID)
	if err != nil {
		return
	}
	g, err := GetStudentsGrades(component.App.DB, id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	list, err := json.Marshal(g)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(list)
}

func FetchStudentsGradesbyTimeStampAndStudentID(w http.ResponseWriter, r *http.Request) {
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

	info, err := GetStudentsGradesbyTimeStampAndStudentID(component.App.DB, studID, timeStart, timeEnd)
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
