package student

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"team-academy/config"
	"team-academy/repository"
	"time"

	"github.com/gorilla/mux"
)

func CreateStudentHandler(w http.ResponseWriter, r *http.Request) {
	var st repository.Student
	err := json.NewDecoder(r.Body).Decode(&st)
	st.StartDate = time.Now().UTC()

	err = repository.CreateStudent(config.Application.Db, st)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	fmt.Fprintf(w, "Student %v registered with sucess", st)
}

func GetStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	student_id := vars["student_id"]

	st_id, err := strconv.Atoi(student_id)
	if err != nil {
		fmt.Fprintln(w, "Error converting student_id to int")
		return
	}

	s, err := repository.GetStudentByID(config.Application.Db, st_id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	stEncoded, err := json.Marshal(s)
	if err != nil {
		fmt.Fprintln(w, "Error using json on student")
		return
	}

	w.Write(stEncoded)
}

func UpdateStudentHandler(w http.ResponseWriter, r *http.Request) {
	var st repository.Student
	err := json.NewDecoder(r.Body).Decode(&st)

	fmt.Fprintf(w, "Student %v\n", st)

	err = repository.UpdateStudent(config.Application.Db, st)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Fprintf(w, "Student %v was updated", st)
}

func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	s_studentid := vars["student_id"]

	st_id, err := strconv.Atoi(s_studentid)
	if err != nil {
		fmt.Fprintln(w, "Error converting student_id to int")
		return
	}

	err = repository.DeleteStudent(config.Application.Db, st_id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Fprintf(w, "Student %d was deleted", st_id)
}
