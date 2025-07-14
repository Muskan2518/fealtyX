package main

import (
	"net/http"

	"github.com/gorilla/mux"
    "github.com/Muskan2518/fealtyX/handles"

)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/students", handles.CreateStudent).Methods("POST")
	r.HandleFunc("/students", handles.GetAllStudents).Methods("GET")
	r.HandleFunc("/students/{id}", handles.GetStudent).Methods("GET")
	r.HandleFunc("/students/{id}", handles.UpdateStudent).Methods("PUT")
	r.HandleFunc("/students/{id}", handles.DeleteStudent).Methods("DELETE")
	r.HandleFunc("/students/{id}/summary", handles.GetStudentSummary).Methods("GET")

	http.ListenAndServe(":8080", r)
}
