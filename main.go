package main

import (
	"fmt"
	"net/http"
	"team-academy/classroom"
	"team-academy/component"
	"team-academy/config"
	"team-academy/department"
	"team-academy/grade"
	"team-academy/professor"
	"team-academy/professor_shift"
	"team-academy/professor_subject"
	"team-academy/schedule"
	"team-academy/shift"
	"team-academy/student"
	"team-academy/student_subject"
	"team-academy/subject"
	"team-academy/student_shift"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Creating and populating database...")
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

	//classroom
	r.HandleFunc("/classroom/{ID}", classroom.FetchClassroomByIDController).Methods("GET")
	r.HandleFunc("/classroom/", classroom.FetchAllClassroomsController).Methods("GET")
	r.HandleFunc("/classroom/", classroom.CreateClassroomController).Methods("POST")
	r.HandleFunc("/classroom/", classroom.UpdateClassroomController).Methods("PUT")
	r.HandleFunc("/classroom/{ID}", classroom.RemoveClassroomController).Methods("DELETE")

	//department
	r.HandleFunc("/department/{ID}", department.FetchDepartmentByIDController).Methods("GET")
	r.HandleFunc("/department/", department.FetchAllDepartmentsController).Methods("GET")
	r.HandleFunc("/department/", department.CreateDepartmentController).Methods("POST")
	r.HandleFunc("/department/", department.UpdateDepartmentController).Methods("PUT")
	r.HandleFunc("/department/{ID}", department.RemoveDepartmentController).Methods("DELETE")

	//grade
	r.HandleFunc("/grade/subject/{ID}/", grade.FetchGradeBySubjectController).Methods("GET")
	r.HandleFunc("/grade/student/{ID}/", grade.FetchGradeByStudentController).Methods("GET")
	r.HandleFunc("/grade/", grade.CreateGradeController).Methods("POST")
	r.HandleFunc("/grade/", grade.UpdateGradeController).Methods("PUT")
	r.HandleFunc("/grade/{studentID}/info/", grade.FetchStudentsGradesController).Methods("GET")
	r.HandleFunc("/grade/{studentID}/info/{beginSchool}/{endSchool}/", grade.FetchStudentsGradesbyTimeStampAndStudentID).Methods("GET")

	//professor
	r.HandleFunc("/professor/", professor.FetchAllProfessorsController).Methods("GET")
	r.HandleFunc("/professor/{ID}/", professor.FetchProfessorController).Methods("GET")
	r.HandleFunc("/professor/{email}/", professor.FetchProfessorByEmailController).Methods("GET")
	r.HandleFunc("/professor/", professor.CreateProfessorController).Methods("POST")
	r.HandleFunc("/professor/", professor.UpdateProfessorController).Methods("PUT")
	r.HandleFunc("/professor/{ID}/", professor.RemoveProfessorController).Methods("DELETE")

	//professor_shift
	r.HandleFunc("/professor/{professorID}/shift/{shiftID}/", professor_shift.AddProfessorToShiftController).Methods("POST")
	r.HandleFunc("/professor/{professorID}/shift/{shiftID}/", professor_shift.RemoveProfessorFromShiftController).Methods("DELETE")

	//professor_subject
	r.HandleFunc("/professor/{ID}/subject/", professor_subject.FetchSubjectsByProfessorIDController).Methods("GET")
	r.HandleFunc("/subject/{ID}/professor/", professor_subject.FetchProfessorsBySubjectIDController).Methods("GET")
	r.HandleFunc("/professor/{professorID}/subject/{subjectID}/", professor_subject.CreateProfessorToSubjectController).Methods("POST")
	r.HandleFunc("/subject/{professorID}/info/{beginSchool}/{endSchool}/", professor_subject.FetchSubjectAndInfobyProfessorIDAndTimeStampController).Methods("GET")

	//schedule
	r.HandleFunc("/schedule/{ID}/", schedule.FetchScheduleController).Methods("GET")
	r.HandleFunc("/schedule/", schedule.CreateScheduleController).Methods("POST")
	r.HandleFunc("/schedule/", schedule.UpdateScheduleController).Methods("PUT")
	r.HandleFunc("/schedule/", schedule.DeleteScheduleController).Methods("DELETE")

	//shift
	r.HandleFunc("/shift/{ID}/", shift.FetchShiftController).Methods("GET")
	r.HandleFunc("/shift/", shift.CreateShiftController).Methods("POST")
	r.HandleFunc("/shift/", shift.UpdateShiftController).Methods("PUT")
	r.HandleFunc("/shift/", shift.DeleteShiftController).Methods("DELETE")

	//student
	r.HandleFunc("/student/{studentID}/", student.FetchStudentController).Methods("GET")
	r.HandleFunc("/student/{email}/", student.FetchStudentByEmailController).Methods("GET")
	r.HandleFunc("/student/", student.FetchAllStudentsController).Methods("GET")
	r.HandleFunc("/student/", student.UpdateStudentController).Methods("PUT")
	r.HandleFunc("/student/", student.CreateStudentController).Methods("POST")
	r.HandleFunc("/student/{studentID}/", student.DeleteStudentController).Methods("DELETE")

	//student_shift
	r.HandleFunc("/student/{studentID}/shift/{shiftID}/", student_shift.AddStudentToShiftController).Methods("POST")
	r.HandleFunc("/student/{studentID}/shift/{shiftID}/", student_shift.RemoveStudentFromShiftController).Methods("DELETE")

	//student_subject
	r.HandleFunc("/student/{studentID}/subject/", student_subject.FetchSubjectsByStudentIDController).Methods("GET")
	r.HandleFunc("/student/{studentID}/info/", student_subject.FetchSubjectAndInfoByStudentIDController).Methods("GET")
	r.HandleFunc("/student/{studentID}/info/{beginSchool}/{endSchool}/", student_subject.FetchSubjectAndInfoByStudentIDAndTimeStampController).Methods("GET")
	r.HandleFunc("/subject/{subjectID}/students/", student_subject.FetchStudentsBySubjectIDController).Methods("GET")
	r.HandleFunc("/subject/{subjectID}/{studentID}/", student_subject.AddStudentToSubjectController).Methods("POST")
	r.HandleFunc("/subject/{subjectID}/{studentID}/", student_subject.RemoveStudentFromSubjectController).Methods("DELETE")
	r.HandleFunc("/subject/{subjectID}/info/", student_subject.FetchStudentAndInfoBySubjectIDController).Methods("GET")

	//subject
	r.HandleFunc("/subject/{ID}/", subject.FetchSubjectByIDController).Methods("GET")
	r.HandleFunc("/subject/", subject.FetchAllSubjectsController).Methods("GET")
	r.HandleFunc("/subject/", subject.CreateSubjectController).Methods("POST")

	err = config.GenerateSwaggerDocsAndEndpoints(r, "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := http.ListenAndServe(":8080", r); err != nil {
		return
	}
}
