package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"team-academy/grade"
	"team-academy/professor"
	"team-academy/student"
	"team-academy/student_subject"
	"team-academy/subject"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	summerfish "github.com/plicca/summerfish-swagger"
)

type App struct {
	Db *gorm.DB
}

var app App

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
	db, err := gorm.Open("sqlite3", "clip_holy_grail.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	app.Db = db
	db.SingularTable(true)
	/*err = populateDatabase(db)
	if err != nil {
		fmt.Println(err)
		return
	}*/

	// ------------------- Gorilla-mux -------------------
	r := mux.NewRouter()

	// CreateStudent doesn't need student_id
	r.HandleFunc("/student/register_student", CreateStudent).Methods("POST")
	r.HandleFunc("/student/{student_id}", GetStudent).Methods("GET")
	r.HandleFunc("/student/update_student", UpdateStudent).Methods("PUT")
	r.HandleFunc("/student/{student_id}", DeleteStudent).Methods("DELETE")

	err = GenerateSwaggerDocsAndEndpoints(r, "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Main router
	http.ListenAndServe(":8080", r)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var st student.Student
	err := json.NewDecoder(r.Body).Decode(&st)
	st.StartDate = time.Now().UTC()

	err = student.CreateStudent(app.Db, st)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	fmt.Fprintf(w, "Student %v registered with sucess", st)
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	student_id := vars["student_id"]

	st_id, err := strconv.Atoi(student_id)
	if err != nil {
		fmt.Fprintln(w, "Error converting student_id to int")
		return
	}

	s, err := student.GetStudentByID(app.Db, st_id)
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

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var st student.Student
	err := json.NewDecoder(r.Body).Decode(&st)

	fmt.Fprintf(w, "Student %v\n", st)

	err = student.UpdateStudent(app.Db, st)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Fprintf(w, "Student %v was updated", st)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	s_studentid := vars["student_id"]

	st_id, err := strconv.Atoi(s_studentid)
	if err != nil {
		fmt.Fprintln(w, "Error converting student_id to int")
		return
	}

	err = student.DeleteStudent(app.Db, st_id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Fprintf(w, "Student %d was deleted", st_id)
}

// ------------------- Gorilla-mux -------------------

func populateDatabase(db *gorm.DB) (err error) {
	err = professor.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	err = student.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	err = subject.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	err = student_subject.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	err = grade.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

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

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			err = student_subject.AddStudentToSubject(db, i, j)
			if err != nil {
				return
			}
		}
	}

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			newGrade := grade.Grade{SubjectID: i, StudentID: j, Rank: "Failed"}
			err = grade.GiveGrade(db, newGrade)
			if err != nil {
				return
			}
		}
	}

	return
}
