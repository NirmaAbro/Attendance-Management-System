package handler

import (
	"attendance-management-system/service"
	"encoding/json"
	"net/http"
	"strings"
)

type AttendanceHandler struct {
	service *service.AttendanceService
}

func NewAttendanceHandler(service *service.AttendanceService) *AttendanceHandler {
	return &AttendanceHandler{service: service}
}

func (h *AttendanceHandler) MarkAttendance(w http.ResponseWriter, r *http.Request) {
	var req struct {
		StudentID string `json:"student_id"`
		Date      string `json:"date"`
		Status    string `json:"status"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	attendance, err := h.service.MarkAttendance(req.StudentID, req.Date, req.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(attendance)
}

func (h *AttendanceHandler) GetAllAttendance(w http.ResponseWriter, r *http.Request) {
	attendance := h.service.GetAllAttendance()
	json.NewEncoder(w).Encode(attendance)
}

// 🔹 ADD: Get attendance by student ID
func (h *AttendanceHandler) GetAttendanceByStudent(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	studentID := parts[len(parts)-1]

	attendance := h.service.GetAttendanceByStudentID(studentID)
	json.NewEncoder(w).Encode(attendance)
}

// 🔹 ADD: Get attendance by date
func (h *AttendanceHandler) GetAttendanceByDate(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	date := parts[len(parts)-1]

	attendance := h.service.GetAttendanceByDate(date)
	json.NewEncoder(w).Encode(attendance)
}

func (h *AttendanceHandler) UpdateAttendance(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-1]

	var req struct {
		StudentID string `json:"student_id"`
		Date      string `json:"date"`
		Status    string `json:"status"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	attendance, err := h.service.UpdateAttendance(id, req.StudentID, req.Date, req.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(attendance)
}

func (h *AttendanceHandler) DeleteAttendance(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-1]

	err := h.service.DeleteAttendance(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

