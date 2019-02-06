package main

import (
	"fmt"
	"net/http"
	"team-academy/component"
	"team-academy/config"
	"team-academy/student"
	"team-academy/student_subject"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	err := component.StartDB()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = config.PopulateDatabase(component.App.DB)
	if err != nil {
		fmt.Println(err)
		return
	}

	r := mux.NewRouter()

	r.HandleFunc("/student/{studentID}/", student.FetchStudentController).Methods("GET")
	r.HandleFunc("/student/", student.FetchAllStudentsController).Methods("GET")
	r.HandleFunc("/student/", student.UpdateStudentController).Methods("PUT")
	r.HandleFunc("/student/", student.CreateStudentController).Methods("POST")
	r.HandleFunc("/student/{studentID}/", student.DeleteStudentController).Methods("DELETE")

	r.HandleFunc("/student/{studentID}/subjects/", student_subject.FetchSubjectsByStudentIDController).Methods("GET")
	r.HandleFunc("/student/{studentID}/info/", student_subject.FetchSubjectAndInfoByStudentIDController).Methods("GET")
	r.HandleFunc("/subject/{subjectID}/students/", student_subject.FetchStudentsBySubjectIDController).Methods("GET")
	r.HandleFunc("/subject/{subjectID}/{studentID}/", student_subject.AddStudentToSubjectController).Methods("POST")
	r.HandleFunc("/subject/{subjectID}/{studentID}/", student_subject.RemoveStudentFromSubjectController).Methods("DELETE")

	err = config.GenerateSwaggerDocsAndEndpoints(r, "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := http.ListenAndServe(":8080", r); err != nil {
		return
	}
}
