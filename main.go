package main

import (
	"fmt"
	"net/http"
	"team-academy/component"
	"team-academy/config"
	"team-academy/grade"
	"team-academy/professor"

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
	r.HandleFunc("/professor/{ID}", professor.GetProfessorController).Methods("GET")
	r.HandleFunc("/professor/delete/{ID}", professor.DeleteProfessorController).Methods("DELETE")
	r.HandleFunc("/professor/update/{ID}", professor.UpdateProfessorController).Methods("PUT")
	r.HandleFunc("/professor/create", professor.PostProfessorController).Methods("POST")
	//Grade Handlers
	r.HandleFunc("/professor/givegrade", grade.PostGradeController).Methods("POST")
	r.HandleFunc("/professor/update", grade.PutGradeController).Methods("PUT")
	r.HandleFunc("/professor/get_by_subject/{ID}", grade.GetGradeBySubjectController).Methods("PUT")
	r.HandleFunc("/professor/get_by_student/{ID}", grade.GetGradeByStudentController).Methods("PUT")
	err = config.GenerateSwaggerDocsAndEndpoints(r, "localhost"+":80")
	if err != nil {
		fmt.Println(err)
		return
	}

	http.ListenAndServe(":8080", r)
}
