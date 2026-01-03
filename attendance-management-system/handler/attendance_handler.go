package handler

import (
	"encoding/json"
	"net/http"

	"attendance-management-system/service"
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
