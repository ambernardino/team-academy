package component

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func ReturnResponse(w http.ResponseWriter, v interface{}) {
	encodedResponse, err := json.Marshal(v)
	if err != nil {
		ErrorHandler(w, ErrMarshallingJSON)
		return
	}

	w.Write(encodedResponse)
}

func ErrorHandler(w http.ResponseWriter, err error) bool {
	serr, ok := err.(*TeamAcademyError)
	if ok {
		return handle(w, serr.Code(), serr)
	}

	return handle(w, http.StatusInternalServerError, err)
}

func ProcessError(err *error) {
	if *err != nil {
		z := *err
		_, ok := z.(*Error)
		if ok {
			return
		}

		*err = New(*err)
	}
}

func ControllerError(w http.ResponseWriter, err error, serr *TeamAcademyError) bool {
	if err != nil {
		log.Printf("Error was: %s", err.Error())
		p, ok := err.(*Error)
		if ok {
			for _, frame := range p.StackFrames() {
				log.Println(frame)
			}
			if serr != nil {
				return ErrorHandler(w, serr)
			}

			se, ok := p.Err.(*TeamAcademyError)
			if ok {
				return ErrorHandler(w, se)
			}
		}

		if serr != nil {
			return ErrorHandler(w, serr)
		}

		return ErrorHandler(w, err)
	}

	return false
}

func handle(w http.ResponseWriter, status int, err error) bool {
	w.WriteHeader(status)
	fmt.Fprint(w, err)
	log.Println(err)
	return true
}
