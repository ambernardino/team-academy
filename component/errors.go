package component

import "net/http"

type TeamAcademyError struct {
	code int
	msg  string
}

//Error returns the user-readable message
func (err TeamAcademyError) Error() string {
	return err.msg
}

//Code returns the http response code of the problem
func (err TeamAcademyError) Code() int {
	return err.code
}

var ErrSomethingAlreadyExists = &TeamAcademyError{http.StatusConflict, "Something is duplicated"}
var ErrStudentAlreadyInSubject = &TeamAcademyError{http.StatusConflict, "Student is already registered in subject"}
var ErrProfessorDoesntExist = &TeamAcademyError{http.StatusConflict, "Professor not enroled in school"}
var ErrProfessorIDIsInvalid = &TeamAcademyError{http.StatusConflict, "Professor has invalid ID number"}
var ErrMissingParameters = &TeamAcademyError{http.StatusConflict, "Missing Parameters"}
var ErrProfessorAlreadyInSubject = &TeamAcademyError{http.StatusConflict, "Professor is already registered in subject"}

