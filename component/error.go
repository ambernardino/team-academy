package component

import (
	"net/http"
)

type TeamAcademyError struct {
	code int
	msg  string
}

var ErrSomethingAlreadyExists = &TeamAcademyError{http.StatusConflict, "Something is duplicated"}
var ErrStudentAlreadyInSubject = &TeamAcademyError{http.StatusConflict, "Student is already registered in subject"}
var ErrMissingParameters = &TeamAcademyError{http.StatusConflict, "Missing parameters"}

//Error returns the user-readable message
func (err TeamAcademyError) Error() string {
	return err.msg
}

//Code returns the http response code of the problem
func (err TeamAcademyError) Code() int {
	return err.code
}
