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
	"team-academy/student_subject"
	"team-academy/subject"

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
	r.HandleFunc("/student/{studentID}/", student.FetchStudentController).Methods("GET")
	r.HandleFunc("/student/{email}/", student.FetchStudentByEmailController).Methods("GET")
	r.HandleFunc("/student/", student.FetchAllStudentsController).Methods("GET")
	r.HandleFunc("/student/", student.UpdateStudentController).Methods("PUT")
	r.HandleFunc("/student/", student.CreateStudentController).Methods("POST")
	r.HandleFunc("/student/{studentID}/", student.DeleteStudentController).Methods("DELETE")

	//student subject
	r.HandleFunc("/student/{studentID}/subjects/", student_subject.FetchSubjectsByStudentIDController).Methods("GET")
	r.HandleFunc("/student/{studentID}/info/", student_subject.FetchSubjectAndInfoByStudentIDController).Methods("GET")
	r.HandleFunc("/subject/{subjectID}/students/", student_subject.FetchStudentsBySubjectIDController).Methods("GET")
	r.HandleFunc("/subject/{subjectID}/{studentID}/", student_subject.AddStudentToSubjectController).Methods("POST")
	r.HandleFunc("/subject/{subjectID}/{studentID}/", student_subject.RemoveStudentFromSubjectController).Methods("DELETE")

	//professor
	r.HandleFunc("/professor/", professor.FetchAllProfessorsController).Methods("GET")
	r.HandleFunc("/professor/{ID}/", professor.FetchProfessorController).Methods("GET")
	r.HandleFunc("/professor/{email}/", professor.FetchProfessorByEmailController).Methods("GET")
	r.HandleFunc("/professor/", professor.CreateProfessorController).Methods("POST")
	r.HandleFunc("/professor/", professor.UpdateProfessorController).Methods("PUT")
	r.HandleFunc("/professor/{ID}/", professor.RemoveProfessorController).Methods("DELETE")

	//subject
	r.HandleFunc("/subject/{ID}/", subject.FetchSubjectByIDController).Methods("GET")
	r.HandleFunc("/subject/", subject.FetchAllSubjectsController).Methods("GET")
	r.HandleFunc("/subject/", subject.CreateSubjectController).Methods("POST")

	//professor subject
	r.HandleFunc("/professor/{ID}/subject/", professor_subject.FetchSubjectsByProfessorIDController).Methods("GET")
	r.HandleFunc("/subject/{ID}/professor/", professor_subject.FetchProfessorsBySubjectIDController).Methods("GET")
	r.HandleFunc("/professor/{professorID}/subject/{subjectID}/", professor_subject.CreateProfessorToSubjectController).Methods("POST")

	//grade
	r.HandleFunc("/grade/subject/{ID}/", grade.FetchGradeBySubjectController).Methods("GET")
	r.HandleFunc("/grade/student/{ID}/", grade.FetchGradeByStudentController).Methods("GET")
	r.HandleFunc("/grade/", grade.CreateGradeController).Methods("POST")
	r.HandleFunc("/grade/", grade.UpdateGradeController).Methods("PUT")

	err = config.GenerateSwaggerDocsAndEndpoints(r, "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := http.ListenAndServe(":8080", r); err != nil {
		return
	}
}
