package professor_subject

import (
	"encoding/json"
	"net/http"
	"strconv"
	"team-academy/component"

	"github.com/gorilla/mux"
)

func FetchSubjectsByProfessorIDController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	professorID := vars["ID"]
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

func FetchProfessorsBySubjectIDController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subjectID := vars["ID"]
	id, err := strconv.Atoi(subjectID)
	if err != nil {
		return
	}

	listOfProfessors, err := GetProfessorsBySubjectID(component.App.DB, id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	professorsInfo, err := json.Marshal(listOfProfessors)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(professorsInfo)
}

func CreateProfessorToSubjectController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	professorID := vars["professorID"]
	subjectID := vars["subjectID"]
	professor, err := strconv.Atoi(professorID)
	if err != nil {
		return
	}

	subject, err := strconv.Atoi(subjectID)
	if err != nil {
		return
	}

	err = AddProfessorToSubject(component.App.DB, professor, subject)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Professor added to Subject"))
}
