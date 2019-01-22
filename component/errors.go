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
//var ErrStudentAlreadyExist = &TeamAcademyError
