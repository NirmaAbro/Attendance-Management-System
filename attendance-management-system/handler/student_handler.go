package handler

import (
	"attendance-management-system/service"
	"encoding/json"
	"net/http"
	"strings"
)

type StudentHandler struct {
	service *service.StudentService
}

func NewStudentHandler(service *service.StudentService) *StudentHandler {
	return &StudentHandler{service: service}
}

func (h *StudentHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	student, err := h.service.CreateStudent(req.Name, req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}

func (h *StudentHandler) GetAllStudents(w http.ResponseWriter, r *http.Request) {
	students := h.service.GetAllStudents()
	json.NewEncoder(w).Encode(students)
}

// update student
func (h *StudentHandler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-1]

	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	student, err := h.service.UpdateStudent(id, req.Name, req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(student)
}

// delete studnet
func (h *StudentHandler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-1]

	err := h.service.DeleteStudent(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *StudentHandler) GetStudentByID(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-1]

	student, err := h.service.GetStudentByID(id)
	if err != nil {
		http.Error(w, "student not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

