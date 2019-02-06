package professor

import (
    "encoding/json"
    "net/http"
    "strconv"
    "team-academy/component"
    "time"

    "github.com/gorilla/mux"
)

func FetchProfessorController(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    ID := vars["ID"]
    id, err := strconv.Atoi(ID)
    if err != nil {
        return
    }
    professor, err := GetProfessorByID(component.App.DB, id)
    if err != nil {
        w.Write([]byte(err.Error()))
        return
    }

    profInfo, err := json.Marshal(professor)
    if err != nil {
        w.Write([]byte(err.Error()))
        return
    }
    w.Write(profInfo)
}

func UpdateProfessorController(w http.ResponseWriter, r *http.Request) {
    var prof Professor

    err := json.NewDecoder(r.Body).Decode(&prof)
    if err != nil {
        return
    }

    err = UpdateProfessorInfo(component.App.DB, prof)
    if err != nil {
        w.Write([]byte(err.Error()))
        return
    }
    w.Write([]byte("Professor Updated"))
}

func CreateProfessorController(w http.ResponseWriter, r *http.Request) {
    var prof Professor
    err := json.NewDecoder(r.Body).Decode(&prof)
    if err != nil {
        return
    }
    prof.StartDate = time.Now().UTC()
    err = CreateProfessor(component.App.DB, prof)
    if err != nil {
        return
    }
    w.Write([]byte("Professor Created"))
}

func RemoveProfessorController(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    ID := vars["ID"]
    id, err := strconv.Atoi(ID)
    if err != nil {
        return
    }
    err = DeleteProfessor(component.App.DB, id)
    if err != nil {
        w.Write([]byte(err.Error()))
        return
    }
    w.Write([]byte("Professor Deleted"))
}

func FetchAllProfessorsController(w http.ResponseWriter, r *http.Request) {
    professors, err := GetAllProfessors(component.App.DB)
    if err != nil {
        w.Write([]byte(err.Error()))
        return
    }

    profsInfo, err := json.Marshal(professors)
    if err != nil {
        w.Write([]byte(err.Error()))
        return
    }
    w.Write(profsInfo)
}