package professor_subject

import (
	"encoding/json"
	"net/http"
	"strconv"
	"team-academy/component"
	"time"

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

func FetchSubjectAndInfobyProfessorIDAndTimeStampController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	professorID := vars["professorID"]
	beginSchool := vars["beginSchool"]
	endSchool := vars["endSchool"]

	profID, err := strconv.Atoi(professorID)
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

	info, err := GetSubjectAndInfobyProfessorIDAndTimeStamp(component.App.DB, profID, timeStart, timeEnd)
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

	err = AddProfessorToSubject(component.App.DB, professor, subject, int64(time.Now().UTC().Unix()))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Professor added to Subject"))
}
