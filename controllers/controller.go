package controllers

import (
	"ASimpleProgram/helper"
	"ASimpleProgram/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func InsertOneController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp model.Employee
	json.NewDecoder(r.Body).Decode(&emp)
	helper.InsertOneRecord(emp)
	json.NewEncoder(w).Encode(emp)
}

func GetOneEmpController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	empId := params["id"]
	emp := helper.GetOneRecord(empId)
	json.NewEncoder(w).Encode(emp)
}

func GetAllEmpController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	emp := helper.GetAllRecords()
	json.NewEncoder(w).Encode(emp)
}

func DeleteOneEmpController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	empId := params["id"]
	str := helper.DeleteRecord(empId)
	json.NewEncoder(w).Encode(str)
}

func UpdateOneEmpController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee model.Employee
	_ = json.NewDecoder(r.Body).Decode(&employee)
	str := helper.UpdateRecord(employee)
	json.NewEncoder(w).Encode(str)
}

func DeleteAllController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	str := helper.DeleteAllRecords()
	json.NewEncoder(w).Encode(str)
}
