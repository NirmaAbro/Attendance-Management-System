package handler

import (
	"encoding/json"
	"net/http"

	"attendance-management-system/service"
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
