
	r := mux.NewRouter()

	//student
	r.HandleFunc("/student/{studentID}/", student.FetchStudentController).Methods("GET")
	r.HandleFunc("/student/{email}/", student.FetchStudentByEmailController).Methods("GET")
	r.HandleFunc("/student/", student.FetchAllStudentsCont