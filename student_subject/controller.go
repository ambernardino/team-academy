package student_subject

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"team-academy/config"
	"team-academy/repository"

	"github.com/gorilla/mux"
)

func CreateStudentSubjectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Sstudentid := vars["student_id"]
	Ssubjectid := vars["subject_id"]

	studentid, err := strconv.Atoi(Sstudentid)
	if err != nil {
		fmt.Fprintln(w, "Error converting student_id to int")
		return
	}
	subjectid, err := strconv.Atoi(Ssubjectid)
	if err != nil {
		fmt.Fprintln(w, "Error converting student_id to int")
		return
	}

	err = repository.AddStudentToSubject(config.Application.Db, studentid, subjectid)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Fprintf(w, "Student %v added to subject %v with sucess", studentid, subjectid)
}

func DeleteStudentSubjectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Sstudentid := vars["student_id"]
	Ssubjectid := vars["subject_id"]

	studentid, err := strconv.Atoi(Sstudentid)
	if err != nil {
		fmt.Fprintln(w, "Error converting student_id to int")
		return
	}
	subjectid, err := strconv.Atoi(Ssubjectid)
	if err != nil {
		fmt.Fprintln(w, "Error converting student_id to int")
		return
	}

	err = repository.RemoveStudentFromSubject(config.Application.Db, studentid, subjectid)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

func GetStudentsInSubjectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Ssubjectid := vars["subject_id"]

	subjectid, err := strconv.Atoi(Ssubjectid)
	if err != nil {
		fmt.Fprintln(w, "Error converting student_id to int")
		return
	}

	students, err := repository.GetStudentsBySubjectID(config.Application.Db, subjectid)
	if err != nil {
		fmt.Fprintln(w, "Error couldn't find students in subject")
	}

	studentsEncoded, err := json.Marshal(students)
	if err != nil {
		fmt.Fprintln(w, "Error using json on student")
		return
	}

	w.Write(studentsEncoded)
}
