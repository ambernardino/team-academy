package main

import (
	"fmt"
	"net/http"
	"team-academy/component"
	"team-academy/config"
	"team-academy/grade"
	"team-academy/professor"
	"team-academy/professor_subject"
	"team-academy/student"
	"team-academy/subject"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	DB *gorm.DB
}

var app App

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
	r.HandleFunc("/professor/", professor.CreateProfessorController).Methods("POST")
	r.HandleFunc("/professor/", professor.UpdateProfessorController).Methods("PUT")
	r.HandleFunc("/professor/{ID}/", professor.RemoveProfessorController).Methods("DELETE")
	//Subject Handlers
	r.HandleFunc("/subject/{ID}/", subject.FetchSubjectByIDController).Methods("GET")
	r.HandleFunc("/subject/", subject.FetchAllSubjectsController).Methods("GET")
	r.HandleFunc("/subject/", subject.CreateSubjectController).Methods("POST")
	//ProfessorSubject Handlers
	r.HandleFunc("/professor/{ID}/subject/", professor_subject.FetchSubjectsByProfessorIDController).Methods("GET")
	r.HandleFunc("/subject/{ID}/professor/", professor_subject.FetchProfessorsBySubjectIDController).Methods("GET")
	r.HandleFunc("/professor/{professorID}/subject/{subjectID}/", professor_subject.CreateProfessorToSubjectController).Methods("POST")
	//Student Handlers
	r.HandleFunc("/student/", student.FetchAllStudentsController).Methods("GET")
	//Grade Handlers
	r.HandleFunc("/grade/subject/{ID}", grade.FetchGradeBySubjectController).Methods("GET")
	r.HandleFunc("/grade/student/{ID}", grade.FetchGradeByStudentController).Methods("GET")
	r.HandleFunc("/grade/", grade.CreateGradeController).Methods("POST")
	r.HandleFunc("/grade/", grade.UpdateGradeController).Methods("PUT")

	err = config.GenerateSwaggerDocsAndEndpoints(r, "localhost:8080")
	if err != nil {
		return
	}

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
