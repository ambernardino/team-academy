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
var ErrProfessorAlreadyInSubject = &TeamAcademyError{http.StatusConflict, "Professor is already registered in subject"}
var ErrProfessorAlreadyInShift = &TeamAcademyError{http.StatusConflict, "Professor is already registered in the shift"}
var ErrProfessorNotInSubject = &TeamAcademyError{http.StatusConflict, "Professor isn't registered in the subject"}
var ErrProfessorDoesntExist = &TeamAcademyError{http.StatusConflict, "Professor not enroled in school"}
var ErrProfessorIDIsInvalid = &TeamAcademyError{http.StatusConflict, "Professor has invalid ID number"}
var ErrShiftDoesntExistInSubject = &TeamAcademyError{http.StatusConflict, "Shift is not assigned to the subject"}
var ErrMarshallingJSON = &TeamAcademyError{http.StatusBadRequest, "Unable to marshall parameters"}
var ErrUnmarshallingJSON = &TeamAcademyError{http.StatusBadRequest, "Unable to unmarshall parameters"}

//Error returns the user-readable message
func (err TeamAcademyError) Error() string {
	return err.msg
}

//Code returns the http response code of the problem
func (err TeamAcademyError) Code() int {
	return err.code
}
