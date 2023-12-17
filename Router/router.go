package router

import (
	"ASimpleProgram/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/employees", controllers.GetAllEmpController).Methods("GET")
	r.HandleFunc("/api/employee/{id}", controllers.GetOneEmpController).Methods("GET")
	r.HandleFunc("/api/employee", controllers.InsertOneController).Methods("POST")
	r.HandleFunc("/api/employee", controllers.UpdateOneEmpController).Methods("PUT")
	r.HandleFunc("/api/employees", controllers.DeleteAllController).Methods("DELETE")
	r.HandleFunc("/api/employee/{id}", controllers.DeleteOneEmpController).Methods("DELETE")
	return r
}
