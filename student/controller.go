package student

import (
	"encoding/json"
	"net/http"
	"team-academy/component"
)

func FetchAllStudentsController(w http.ResponseWriter, r *http.Request) {
	students, err := GetAllStudents(component.App.DB)
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
