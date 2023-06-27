package handler

import (
	"encoding/json"
	"net/http"

	M "github.com/farhaanfsk/Go-Gorilla-Mux-CRUD/models"
	"github.com/farhaanfsk/Go-Gorilla-Mux-CRUD/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	S *service.Service
}

func (h *Handler) GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.S.GetEmployees())
}

func (h *Handler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var e M.Employee
	json.NewDecoder(r.Body).Decode(&e)
	e = h.S.CreateEmployee(e)
	json.NewEncoder(w).Encode(e)
}

func (h *Handler) GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	e := h.S.GetEmployee(id)
	if e != (M.Employee{}) {
		json.NewEncoder(w).Encode(e)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Employee not found")

}

func (h *Handler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	isDeleted := h.S.DeleteEmployee(id)
	if isDeleted {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Employee not found")
}

func (h *Handler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var e M.Employee
	json.NewDecoder(r.Body).Decode(&e)
	isUpdated := h.S.UpdateEmployee(e)
	if isUpdated {
		json.NewEncoder(w).Encode(e)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Employee not found")
}

func NewHandler(s *service.Service) Handler {
	return Handler{s}
}
