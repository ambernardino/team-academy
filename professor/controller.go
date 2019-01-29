package professor

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"team-academy/component"
	"time"
)

func GetProfessorController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["ID"]
	id, err := strconv.Atoi(ID)
	if err != nil {
		return
	}
	professor, err := GetProfessorByID(component.App.DB, id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	profInfo, err := json.Marshal(professor)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(profInfo)
}

func UpdateProfessorController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["ID"]
	id, err := strconv.Atoi(ID)
	if err != nil {
		return
	}
	var prof Professor
	err = json.NewDecoder(r.Body).Decode(&prof)
	if err != nil {
		return
	}
	prof.ID = id
	err = UpdateProfessorInfo(component.App.DB, prof)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("Professor Updated"))
}

func PostProfessorController(w http.ResponseWriter, r *http.Request) {
	var prof Professor
	err := json.NewDecoder(r.Body).Decode(&prof)
	if err != nil {
		return
	}
	prof.StartDate = time.Now().UTC()
	err = CreateProfessor(component.App.DB, prof)
	if err != nil {
		return
	}
	w.Write([]byte("Professor Created"))
}

func DeleteProfessorController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["ID"]
	id, err := strconv.Atoi(ID)
	if err != nil {
		return
	}
	err = DeleteProfessor(component.App.DB, id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("Professor Deleted"))
}