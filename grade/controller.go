package grade

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"team-academy/component"
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
