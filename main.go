package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"team-academy/grade"
	"team-academy/professor"
	"team-academy/student"
	"team-academy/student_subject"
	"team-academy/subject"
	"time"

	"github.com/gorilla/mux"
	summerfish "github.com/plicca/summerfish-swagger"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	DB *gorm.DB
}

var app App

func main() {
	db, err := gorm.Open("sqlite3", "clip_holy_grail.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	db.SingularTable(true)
	err = populateDatabase(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	app = App{DB: db}

	r := mux.NewRouter()
	r.HandleFunc("/professor/{professorID}", GetProfessor).Methods("GET")
	r.HandleFunc("/professor/create", CreateProfessor).Methods("POST")
	r.HandleFunc("/professor/update", UpdateProfessor).Methods("PUT")
	r.HandleFunc("/professor/{professorID}", DeleteProfessor).Methods("DELETE")
	r.HandleFunc("/professor/get_grade_by_student_id/{studentID}", GetGradeByStudentID).Methods("GET")
	r.HandleFunc("/professor/get_grade_by_subject_id/{subjectID}", GetGradeBySubjectID).Methods("GET")
	r.HandleFunc("/professor/give_grade", GiveGrade).Methods("POST")
	r.HandleFunc("/professor/update_grade", UpdateGrade).Methods("PUT")
	err = GenerateSwaggerDocsAndEndpoints(r, "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := http.ListenAndServe(":8080", r); err != nil {
		return
	}
}

func GetProfessor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	professorID := vars["professorID"]
	ID, err := strconv.Atoi(professorID)
	if err != nil {
		return
	}

	prof, err := professor.GetProfessorByID(app.DB, ID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	array, err := json.Marshal(prof)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(array)
}

func CreateProfessor(w http.ResponseWriter, r *http.Request) {
	var prof professor.Professor
	err := json.NewDecoder(r.Body).Decode(&prof)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	prof.StartDate = time.Now().UTC()
	err = professor.CreateProfessor(app.DB, prof)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Fprintf(w, "%v", prof)
}

func UpdateProfessor(w http.ResponseWriter, r *http.Request) {
	var prof professor.Professor
	err := json.NewDecoder(r.Body).Decode(&prof)
	if err != nil {
		return
	}

	err = professor.UpdateProfessorInfo(app.DB, prof)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Fprintf(w, "%v", prof)
}

func DeleteProfessor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	professorID := vars["professorID"]
	ID, err := strconv.Atoi(professorID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Println(ID)

	err = professor.DeleteProfessor(app.DB, ID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

func GetGradeByStudentID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["studentID"]
	ID, err := strconv.Atoi(studentID)
	if err != nil {
		return
	}

	grd, err := grade.GetGradeByStudentID(app.DB, ID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	array, err := json.Marshal(grd)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(array)
}

func GetGradeBySubjectID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subjectID := vars["subjectID"]
	ID, err := strconv.Atoi(subjectID)
	if err != nil {
		return
	}

	grd, err := grade.GetGradeBySubjectID(app.DB, ID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	array, err := json.Marshal(grd)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(array)
}

func GiveGrade(w http.ResponseWriter, r *http.Request) {
	var grd grade.Grade
	err := json.NewDecoder(r.Body).Decode(&grd)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = grade.GiveGrade(app.DB, grd)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Fprintf(w, "%v", grd)
}

func UpdateGrade(w http.ResponseWriter, r *http.Request) {
	var grd grade.Grade
	err := json.NewDecoder(r.Body).Decode(&grd)
	if err != nil {
		return
	}

	err = grade.UpdateGrade(app.DB, grd)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Fprintf(w, "%v", grd)
}

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

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

func populateDatabase(db *gorm.DB) (err error) {
	existsProfessorTable, err := professor.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	existsStudentTable, err := student.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	existsSubjectTable, err := subject.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	existsStudentSubjectTable, err := student_subject.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	existsGradeTable, err := grade.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	if !existsSubjectTable {
		newSubject := subject.Subject{ID: 1, Name: "Cadeira 1", Description: "Nothing"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{ID: 2, Name: "Cadeira 2", Description: "Nothing"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{ID: 3, Name: "Cadeira 3", Description: "Nothing"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}
	}

	if !existsProfessorTable {
		newProfessor := professor.Professor{ID: 1, FirstName: "Prof 1", LastName: "Prof 1", CursoIDs: "Curso 1", CadeiraIDS: "Cadeira 1", StartDate: time.Now().UTC()}
		err = professor.CreateProfessor(db, newProfessor)
		if err != nil {
			return
		}

		newProfessor = professor.Professor{ID: 2, FirstName: "Prof 2", LastName: "Prof 2", CursoIDs: "Curso 2", CadeiraIDS: "Cadeira 2", StartDate: time.Now().UTC()}
		err = professor.CreateProfessor(db, newProfessor)
		if err != nil {
			return
		}

		newProfessor = professor.Professor{ID: 3, FirstName: "Prof 3", LastName: "Prof 3", CursoIDs: "Curso 3", CadeiraIDS: "Cadeira 3", StartDate: time.Now().UTC()}
		err = professor.CreateProfessor(db, newProfessor)
		if err != nil {
			return
		}
	}

	if !existsStudentTable {
		newStudent := student.Student{ID: 1, FirstName: "Student 1", LastName: "Student 1", CursoID: 1, StartDate: time.Now().UTC()}
		err = student.CreateStudent(db, newStudent)
		if err != nil {
			return
		}

		newStudent = student.Student{ID: 2, FirstName: "Student 2", LastName: "Student 2", CursoID: 2, StartDate: time.Now().UTC()}
		err = student.CreateStudent(db, newStudent)
		if err != nil {
			return
		}

		newStudent = student.Student{ID: 3, FirstName: "Student 3", LastName: "Student 3", CursoID: 3, StartDate: time.Now().UTC()}
		err = student.CreateStudent(db, newStudent)
		if err != nil {
			return
		}
	}

	if !existsStudentSubjectTable {
		for i := 1; i <= 3; i++ {
			for j := 1; j <= 3; j++ {
				err = student_subject.AddStudentToSubject(db, i, j)
				if err != nil {
					return
				}
			}
		}
	}

	if !existsGradeTable {
		for i := 1; i <= 3; i++ {
			for j := 1; j <= 3; j++ {
				newGrade := grade.Grade{SubjectID: i, StudentID: j, Rank: "Failed"}
				err = grade.GiveGrade(db, newGrade)
				if err != nil {
					return
				}
			}
		}
	}

	return
}
