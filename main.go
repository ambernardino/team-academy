package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"team-academy/config"
	"team-academy/repository"
	"team-academy/student"
	"team-academy/student_subject"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	summerfish "github.com/plicca/summerfish-swagger"
)

func GenerateSwaggerDocsAndEndpoints(router *mux.Router, endpoint string) (err error) {
	config := summerfish.Config{
		Schemes:                []string{"http", "https"},
		SwaggerFileRoute:       summerfish.SwaggerFileRoute,
		SwaggerFilePath:        summerfish.SwaggerFileRoute,
		SwaggerFileHeaderRoute: summerfish.SwaggerFileRoute,
		SwaggerUIRoute:         summerfish.SwaggerUIRoute,
		BaseRoute:              "/",
	}

	config.SwaggerFilePath, err = filepath.Abs("res/swagger.json")
	if err != nil {
		return
	}

	routerInformation, err := summerfish.GetInfoFromRouter(router)
	if err != nil {
		return
	}

	scheme := summerfish.SchemeHolder{Schemes: config.Schemes, Host: endpoint, BasePath: config.BaseRoute}
	err = scheme.GenerateSwaggerFile(routerInformation, config.SwaggerFilePath)
	if err != nil {
		return
	}

	log.Println("Swagger documentation generated")
	return summerfish.AddSwaggerUIEndpoints(router, config)
}

func main() {
	err := config.StartDatabase()
	if err != nil {
		fmt.Println(err)
		return
	}

	/*err = populateDatabase(db)
	if err != nil {
		fmt.Println(err)
		return
	}*/

	// ------------------- Gorilla-mux -------------------
	r := mux.NewRouter()

	// Student handle function
	r.HandleFunc("/student/register_student", student.CreateStudentHandler).Methods("POST")
	r.HandleFunc("/student/{student_id}", student.GetStudentHandler).Methods("GET")
	r.HandleFunc("/student/update_student", student.UpdateStudentHandler).Methods("PUT")
	r.HandleFunc("/student/{student_id}", student.DeleteStudentHandler).Methods("DELETE")

	// Student subject
	r.HandleFunc("/subject/{subject_id}/student/{student_id}", student_subject.CreateStudentSubjectHandler).Methods("POST")
	r.HandleFunc("/subject/{subject_id}/student/{student_id}", student_subject.DeleteStudentSubjectHandler).Methods("DELETE")
	r.HandleFunc("/subject/{subject_id}", student_subject.GetStudentsInSubjectHandler).Methods("GET")

	err = GenerateSwaggerDocsAndEndpoints(r, "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Main router
	http.ListenAndServe(":8080", r)
}

// ------------------- Gorilla-mux -------------------

func PopulateDatabase(db *gorm.DB) (err error) {
	err = repository.CreateStudentTableIfNotExists(db)
	if err != nil {
		return
	}

	err = repository.CreateStudentSubjectTableIfNotExists(db)
	if err != nil {
		return
	}

	err = repository.CreateProfessorTableIfNotExists(db)
	if err != nil {
		return
	}

	err = repository.CreateSubjectTableIfNotExists(db)
	if err != nil {
		return
	}

	err = repository.CreateGradeTableIfNotExists(db)
	if err != nil {
		return
	}

	newSubject := repository.Subject{ID: 1, Name: "Cadeira 1", Description: "Nothing"}
	err = repository.CreateSubject(db, newSubject)
	if err != nil {
		return
	}

	newSubject = repository.Subject{ID: 2, Name: "Cadeira 2", Description: "Nothing"}
	err = repository.CreateSubject(db, newSubject)
	if err != nil {
		return
	}

	newSubject = repository.Subject{ID: 3, Name: "Cadeira 3", Description: "Nothing"}
	err = repository.CreateSubject(db, newSubject)
	if err != nil {
		return
	}

	newProfessor := repository.Professor{ID: 1, FirstName: "Prof 1", LastName: "Prof 1", CursoIDs: "Curso 1", CadeiraIDS: "Cadeira 1", StartDate: time.Now().UTC()}
	err = repository.CreateProfessor(db, newProfessor)
	if err != nil {
		return
	}

	newProfessor = repository.Professor{ID: 2, FirstName: "Prof 2", LastName: "Prof 2", CursoIDs: "Curso 2", CadeiraIDS: "Cadeira 2", StartDate: time.Now().UTC()}
	err = repository.CreateProfessor(db, newProfessor)
	if err != nil {
		return
	}

	newProfessor = repository.Professor{ID: 3, FirstName: "Prof 3", LastName: "Prof 3", CursoIDs: "Curso 3", CadeiraIDS: "Cadeira 3", StartDate: time.Now().UTC()}
	err = repository.CreateProfessor(db, newProfessor)
	if err != nil {
		return
	}

	newStudent := repository.Student{ID: 1, FirstName: "Student 1", LastName: "Student 1", CursoID: 1, StartDate: time.Now().UTC()}
	err = repository.CreateStudent(db, newStudent)
	if err != nil {
		return
	}

	newStudent = repository.Student{ID: 2, FirstName: "Student 2", LastName: "Student 2", CursoID: 2, StartDate: time.Now().UTC()}
	err = repository.CreateStudent(db, newStudent)
	if err != nil {
		return
	}

	newStudent = repository.Student{ID: 3, FirstName: "Student 3", LastName: "Student 3", CursoID: 3, StartDate: time.Now().UTC()}
	err = repository.CreateStudent(db, newStudent)
	if err != nil {
		return
	}

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			err = repository.AddStudentToSubject(db, i, j)
			if err != nil {
				return
			}
		}
	}

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			newGrade := repository.Grade{SubjectID: i, StudentID: j, Rank: "Failed"}
			err = repository.GiveGrade(db, newGrade)
			if err != nil {
				return
			}
		}
	}

	return
}
