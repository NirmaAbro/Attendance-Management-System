package main

import (
	"net/http"

	"attendance-management-system/handler"
	"attendance-management-system/repository"
	"attendance-management-system/service"
)

func main() {
	studentRepo := repository.NewStudentRepository()
	attendanceRepo := repository.NewAttendanceRepository()

	studentService := service.NewStudentService(studentRepo)
	attendanceService := service.NewAttendanceService(attendanceRepo, studentRepo)

	studentHandler := handler.NewStudentHandler(studentService)
	attendanceHandler := handler.NewAttendanceHandler(attendanceService)

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			studentHandler.CreateStudent(w, r)
		case http.MethodGet:
			studentHandler.GetAllStudents(w, r)
		}
	})

	http.HandleFunc("/attendance", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			attendanceHandler.MarkAttendance(w, r)
		case http.MethodGet:
			attendanceHandler.GetAllAttendance(w, r)
		}
	})

	http.ListenAndServe(":8080", nil)
}
