package main

import (
	"team-academy/student"
	"team-academy/subject"
	"fmt"
	"net/http"
	"team-academy/component"
	"team-academy/config"
	"team-academy/grade"
	"team-academy/professor"
	"team-academy/professor_subject"

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
	//Professor Handlers
	r.HandleFunc("/professor/", professor.FetchAllProfessorsController).Methods("GET")
	r.HandleFunc("/professor/{ID}/", professor.FetchProfessorController).Methods("GET")
	r.HandleFunc("/professor/{ID}/", professor.RemoveProfessorController).Methods("DELETE")
	r.HandleFunc("/professor/", professor.UpdateProfessorController).Methods("PUT")
	r.HandleFunc("/professor/", professor.CreateProfessorController).Methods("POST")
	//Professor_Subject Handlers
	r.HandleFunc("/professor/{ID}/subject/", professor_subject.FetchSubjectsByProfessorIDController).Methods("GET")
	r.HandleFunc("/subject/{ID}/professor/", professor_subject.FetchProfessorsBySubjectIDController).Methods("GET")
	r.HandleFunc("/professor/{professorID}/subject/{subjectID}/", professor_subject.AddProfessorToSubjectController).Methods("POST")
	//Grade Handlers
	r.HandleFunc("/grade/", grade.CreateGradeController).Methods("POST")
	r.HandleFunc("/grade/", grade.UpdateGradeController).Methods("PUT")
	r.HandleFunc("/grade/subject/{ID}/", grade.FetchGradeBySubjectController).Methods("PUT")
	r.HandleFunc("/grade/student/{ID}/", grade.FetchGradeByStudentController).Methods("PUT")
	//Subject Handlers
	r.HandleFunc("/subject/", subject.FetchAllSubjectsController).Methods("GET")
	r.HandleFunc("/subject/", subject.CreateSubjectController).Methods("POST")
	r.HandleFunc("/subject/{ID}", subject.FetchSubjectByIDController).Methods("GET")
	//Student Handlers
	r.HandleFunc("/student/", student.FetchAllStudentsController).Methods("GET")

	err = config.GenerateSwaggerDocsAndEndpoints(r, "localhost"+":8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	http.ListenAndServe(":8080", r)
}
