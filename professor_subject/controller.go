package professor_subject

import (
	"encoding/json"
	"net/http"
	"strconv"
	"team-academy/component"

	"github.com/gorilla/mux"
)

func GetSubjectsByProfessorIDController(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    professorID := vars["professorID"]
    id, err := strconv.Atoi(professorID)
    if err != nil {
        return
    }

    subjects, err := GetSubjectsAndInfoByProfessorID(component.App.DB, id)
    if err != nil {
        w.Write([]byte(err.Error()))
        return
    }

    subjectsInfo, err := json.Marshal(subjects)
    if err != nil {
        w.Write([]byte(err.Error()))
        return
    }
    w.Write(subjectsInfo)
}

func GetProfessorsBySubjectIDController(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	subjectID := vars["subjectID"]
	id, err := strconv.Atoi(subjectID)
	if err != nil {
		return
	}
	listOfProfessors, err := GetProfessorsBySubjectID(component.App.DB, id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	profsInfo, err := json.Marshal(listOfProfessors)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(profsInfo)
}
