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

	//student
	r.HandleFunc("/student/{studentID}", student.GetStudentController).Methods("GET")
	r.HandleFunc("/students", student.GetAllStudentsController).Methods("GET")
	r.HandleFunc("/student/update", student.UpdateStudentController).Methods("PUT")
	r.HandleFunc("/student/create", student.CreateStudentController).Methods("POST")
	r.HandleFunc("/student/delete/{studentID}", student.DeleteStudentController).Methods("DELETE")

	//student subject
	r.HandleFunc("/student/{studentID}/subjects", student_subject.GetSubjectsByStudentIDController).Methods("GET")
	r.HandleFunc("/student/{studentID}/subjects/info", student_subject.GetSubjectAndInfoByStudentIDController).Methods("GET")
	r.HandleFunc("/subject/{subjectID}/students", student_subject.GetStudentsBySubjectIDController).Methods("GET")
	r.HandleFunc("/subject/{subjectID}/student/{studentID}", student_subject.AddStudentToSubjectController).Methods("POST")
	r.HandleFunc("/subject/{subjectID}/remove/{studentID}", student_subject.RemoveStudentFromSubjectController).Methods("DELETE")

	err = config.GenerateSwaggerDocsAndEndpoints(r, "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := http.ListenAndServe(":8080", r); err != nil {
		return
	}
}
