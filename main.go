package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Employee struct {
	Id      uuid.UUID
	Name    string
	Address Address
}

type Address struct {
	City    string
	State   string
	Country string
}

var Employees []Employee

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(Employees)

	json.NewEncoder(w).Encode(Employees)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var e Employee
	json.NewDecoder(r.Body).Decode(&e)
	e.Id = uuid.New()
	Employees = append(Employees, e)
	json.NewEncoder(w).Encode(e)
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	for i := range Employees {
		if Employees[i].Id.String() == id {
			json.NewEncoder(w).Encode(Employees[i])
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Employee not found")

}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	for i := range Employees {
		if Employees[i].Id.String() == id {
			Employees = append(Employees[:i], Employees[i+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Employee not found")
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var e Employee
	json.NewDecoder(r.Body).Decode(&e)
	for i := range Employees {
		if Employees[i].Id == e.Id {
			Employees[i] = e
			json.NewEncoder(w).Encode(e)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Employee not found")
}

func main() {
	router := mux.NewRouter()
	Employees = append(Employees, Employee{
		Id:   uuid.New(),
		Name: "Test1",
		Address: Address{
			City:    "hyd",
			State:   "AP",
			Country: "Ind",
		},
	})
	router.HandleFunc("/employees", getEmployees).Methods("GET")
	router.HandleFunc("/employees/{id}", getEmployee).Methods("GET")
	router.HandleFunc("/employees", createEmployee).Methods("POST")
	router.HandleFunc("/employees/{id}", deleteEmployee).Methods("DELETE")
	router.HandleFunc("/employees", updateEmployee).Methods("PUT")
	http.ListenAndServe(":1234", router)
}
