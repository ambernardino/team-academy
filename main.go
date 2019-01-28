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
	app = App{DB: db}
	if err != nil {
		fmt.Println(err)
		return
	}

	r := mux.NewRouter()

	r.HandleFunc("/student/{studentID}", GetStudent).Methods("GET")
	r.HandleFunc("/student/{studentID}/subjects", GetSubjectsByStudentID).Methods("GET")
	r.HandleFunc("/subject/{subjectID}/students", GetStudentsBySubjectID).Methods("GET")
	r.HandleFunc("/student/{studentID}/info", GetSubjectAndInfoByStudentID).Methods("GET")
	r.HandleFunc("/student/create", CreateStudent).Methods("POST")
	r.HandleFunc("/subject/{subjectID}/student/{studentID}", AddStudentToSubject).Methods("POST")
	r.HandleFunc("/student/update", UpdateStudent).Methods("PUT")
	r.HandleFunc("/student/delete/{studentID}", DeleteStudent).Methods("DELETE")
	r.HandleFunc("/subject/{subjectID}/remove/{studentID}", RemoveStudentFromSubject).Methods("DELETE")

	err = GenerateSwaggerDocsAndEndpoints(r, "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := http.ListenAndServe(":8080", r); err != nil {
		return
	}
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["studentID"]

	ID, err := strconv.Atoi(studentID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	student, err := student.GetStudentByID(app.DB, ID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	encodedStudent, err := json.Marshal(student)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(encodedStudent)
}

func GetSubjectsByStudentID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["studentID"]

	ID, err := strconv.Atoi(studentID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	subjects, err := student_subject.GetSubjectsByStudentID(app.DB, ID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	encodedSubjects, err := json.Marshal(subjects)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(encodedSubjects)
}

func GetStudentsBySubjectID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subjectID := vars["subjectID"]

	ID, err := strconv.Atoi(subjectID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	students, err := student_subject.GetStudentsBySubjectID(app.DB, ID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	encodedStudents, err := json.Marshal(students)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(encodedStudents)
}

func GetSubjectAndInfoByStudentID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["studentID"]

	ID, err := strconv.Atoi(studentID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Println(ID)

	info, err := student_subject.GetSubjectAndInfoByStudentID(app.DB, ID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	encodedInfo, err := json.Marshal(info)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(encodedInfo)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var encodedStudent student.Student

	err := json.NewDecoder(r.Body).Decode(&encodedStudent)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	encodedStudent.StartDate = time.Now().UTC()

	err = student.CreateStudent(app.DB, encodedStudent)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Fprintf(w, "%v", encodedStudent)
}

func AddStudentToSubject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subjectID := vars["subjectID"]
	studentID := vars["studentID"]

	subjID, err := strconv.Atoi(subjectID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	studID, err := strconv.Atoi(studentID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = student_subject.AddStudentToSubject(app.DB, studID, subjID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

func RemoveStudentFromSubject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subjectID := vars["subjectID"]
	studentID := vars["studentID"]

	subjID, err := strconv.Atoi(subjectID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	studID, err := strconv.Atoi(studentID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = student_subject.RemoveStudentFromSubject(app.DB, studID, subjID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var encodedStudent student.Student

	err := json.NewDecoder(r.Body).Decode(&encodedStudent)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = student.UpdateStudent(app.DB, encodedStudent)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Fprintf(w, "%v", encodedStudent)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["studentID"]

	ID, err := strconv.Atoi(studentID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = student.DeleteStudent(app.DB, ID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
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

func populateDatabase(db *gorm.DB) (err error) {

	professorTable, err := professor.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	studentTable, err := student.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	subjectTable, err := subject.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	gradeTable, err := grade.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	studentSubjectTable, err := student_subject.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	if !subjectTable {
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

	if !professorTable {
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

	if !studentTable {
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

	if !studentSubjectTable {
		for i := 1; i <= 3; i++ {
			for j := 1; j <= 3; j++ {
				err = student_subject.AddStudentToSubject(db, i, j)
				if err != nil {
					return
				}
			}
		}
	}

	if !gradeTable {
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
