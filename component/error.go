package component

import (
	"net/http"
)

type TeamAcademyError struct {
	code int
	msg  string
}

var ErrSubjectAlreadyExists = &TeamAcademyError{http.StatusConflict, "A subject with that name already exists"}
var ErrStudentAlreadyInSubject = &TeamAcademyError{http.StatusConflict, "Student is already registered in subject"}
var ErrMissingParameters = &TeamAcademyError{http.StatusConflict, "Missing parameters"}
var ErrProfessorAlreadyInSubject = &TeamAcademyError{http.StatusConflict, "Professor is already registered in subject"}
var ErrProfessorDoesntExist = &TeamAcademyError{http.StatusConflict, "Professor not enroled in school"}
var ErrStudentDoesntExist = &TeamAcademyError{http.StatusConflict, "Student not enroled in school"}
var ErrProfessorIDIsInvalid = &TeamAcademyError{http.StatusConflict, "Professor has invalid ID number"}
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
