package component

import (
	"net/http"
)

type TeamAcademyError struct {
	code int
	msg  string
}

var ErrInvalidStartTime = &TeamAcademyError{http.StatusConflict, "Shift can't begin before 8 am"}
var ErrInvalidEndTime = &TeamAcademyError{http.StatusConflict, "Shift can't end after 12 am"}
var ErrInvalidDuration = &TeamAcademyError{http.StatusConflict, "Schedule can't last less than 1 hour"}
var ErrShiftDoesntExist = &TeamAcademyError{http.StatusConflict, "Shift doesn't exist"}
var ErrWeekdayDoesntExist = &TeamAcademyError{http.StatusConflict, "Invalid weekday"}
var ErrClassroomDoesntExist = &TeamAcademyError{http.StatusConflict, "Classroom doesn't exist"}
var ErrSubjectAlreadyExists = &TeamAcademyError{http.StatusConflict, "A subject with that name already exists"}
var ErrSubjectDoesntExist = &TeamAcademyError{http.StatusConflict, "Subject doesn't exist"}
var ErrStudentAlreadyInSubject = &TeamAcademyError{http.StatusConflict, "Student is already registered in subject"}
var ErrMissingParameters = &TeamAcademyError{http.StatusConflict, "Missing parameters"}
var ErrProfessorAlreadyInSubject = &TeamAcademyError{http.StatusConflict, "Professor is already registered in subject"}
var ErrProfessorAlreadyInShift = &TeamAcademyError{http.StatusConflict, "A Professor is already registered in this shift"}
var ErrProfessorNotInSubject = &TeamAcademyError{http.StatusConflict, "Professor isn't registered in the subject"}
var ErrStudentNotInSubject = &TeamAcademyError{http.StatusConflict, "Student isn't registered in the subject"}
var ErrProfessorDoesntExist = &TeamAcademyError{http.StatusConflict, "Professor not enroled in school"}
var ErrStudentDoesntExist = &TeamAcademyError{http.StatusConflict, "Student not enroled in school"}
var ErrProfessorIDIsInvalid = &TeamAcademyError{http.StatusConflict, "Professor has invalid ID number"}
var ErrShiftDoesntExistInSubject = &TeamAcademyError{http.StatusConflict, "Shift is not assigned to the subject"}
var ErrGradeAlreadyGiven = &TeamAcademyError{http.StatusConflict, "Grade was already given"}
var ErrDepartmentDoesntExist = &TeamAcademyError{http.StatusConflict, "Department doesn't exist"}
var ErrMarshallingJSON = &TeamAcademyError{http.StatusBadRequest, "Unable to marshall parameters"}
var ErrUnmarshallingJSON = &TeamAcademyError{http.StatusBadRequest, "Unable to unmarshall parameters"}
var ErrStudentAlreadyInShift = &TeamAcademyError{http.StatusBadRequest, "Student already registered in shift"}

//Error returns the user-readable message
func (err TeamAcademyError) Error() string {
	return err.msg
}

//Code returns the http response code of the problem
func (err TeamAcademyError) Code() int {
	return err.code
}
