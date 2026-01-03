// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"strings"

// 	"attendance-management-system/handler"
// 	"attendance-management-system/repository"
// 	"attendance-management-system/service"
// )

// func main() {

// 	studentRepo := repository.NewStudentRepository()
// 	attendanceRepo := repository.NewAttendanceRepository()

// 	studentService := service.NewStudentService(studentRepo)
// 	attendanceService := service.NewAttendanceService(attendanceRepo, studentRepo)

// 	studentHandler := handler.NewStudentHandler(studentService)
// 	attendanceHandler := handler.NewAttendanceHandler(attendanceService)

// 	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
// 		switch r.Method {
// 		case http.MethodPost:
// 			studentHandler.CreateStudent(w, r)
// 		case http.MethodGet:
// 			studentHandler.GetAllStudents(w, r)
// 		}
// 	})

// 	http.HandleFunc("/attendance", func(w http.ResponseWriter, r *http.Request) {
// 		switch r.Method {
// 		case http.MethodPost:
// 			attendanceHandler.MarkAttendance(w, r)
// 		case http.MethodGet:
// 			attendanceHandler.GetAllAttendance(w, r)
// 		}
// 	})

// 	http.HandleFunc("/attendance/student/", func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method == http.MethodGet {
// 			attendanceHandler.GetAttendanceByStudent(w, r)
// 		}
// 	})

// 	http.HandleFunc("/attendance/date/", func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method == http.MethodGet {
// 			attendanceHandler.GetAttendanceByDate(w, r)
// 		}
// 	})

// 	http.HandleFunc("/attendance", func(w http.ResponseWriter, r *http.Request) {
// 		switch r.Method {
// 		case http.MethodPost:
// 			attendanceHandler.MarkAttendance(w, r)
// 		case http.MethodGet:
// 			attendanceHandler.GetAllAttendance(w, r)
// 		default:
// 			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
// 		}
// 	})

// 	http.HandleFunc("/attendance/", func(w http.ResponseWriter, r *http.Request) {
// 		switch r.Method {
// 		case http.MethodPut:
// 			attendanceHandler.UpdateAttendance(w, r)
// 		case http.MethodDelete:
// 			attendanceHandler.DeleteAttendance(w, r)
// 		case http.MethodGet:
// 			// for report endpoints
// 			if strings.Contains(r.URL.Path, "/student/") {
// 				attendanceHandler.GetAttendanceByStudent(w, r)
// 			} else if strings.Contains(r.URL.Path, "/date/") {
// 				attendanceHandler.GetAttendanceByDate(w, r)
// 			}
// 		default:
// 			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
// 		}
// 	})

// 	fmt.Println("✅ Server running on http://localhost:8080")
// 	http.ListenAndServe(":8080", nil)

// }

package main

import (
	"fmt"
	"net/http"
	"strings"

	"attendance-management-system/handler"
	"attendance-management-system/repository"
	"attendance-management-system/service"
)

func main() {

	// repositories
	studentRepo := repository.NewStudentRepository()
	attendanceRepo := repository.NewAttendanceRepository()

	// services
	studentService := service.NewStudentService(studentRepo)
	attendanceService := service.NewAttendanceService(attendanceRepo, studentRepo)

	// handlers
	studentHandler := handler.NewStudentHandler(studentService)
	attendanceHandler := handler.NewAttendanceHandler(attendanceService)

	// ================= STUDENTS =================
	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			studentHandler.CreateStudent(w, r)
		case http.MethodGet:
			studentHandler.GetAllStudents(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// http.HandleFunc("/students/", func(w http.ResponseWriter, r *http.Request) {
	// 	switch r.Method {
	// 	case http.MethodPut:
	// 		studentHandler.UpdateStudent(w, r)
	// 	case http.MethodDelete:
	// 		studentHandler.DeleteStudent(w, r)
	// 	default:
	// 		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	// 	}
	// })
	http.HandleFunc("/students/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case http.MethodGet:
			studentHandler.GetStudentByID(w, r)

		case http.MethodPut:
			studentHandler.UpdateStudent(w, r)

		case http.MethodDelete:
			studentHandler.DeleteStudent(w, r)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// ================= ATTENDANCE =================

	// Create & list attendance
	http.HandleFunc("/attendance", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			attendanceHandler.MarkAttendance(w, r)
		case http.MethodGet:
			attendanceHandler.GetAllAttendance(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Update, delete & reports
	http.HandleFunc("/attendance/", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodPut:
			attendanceHandler.UpdateAttendance(w, r)

		case http.MethodDelete:
			attendanceHandler.DeleteAttendance(w, r)

		case http.MethodGet:
			// Reports
			if strings.HasPrefix(r.URL.Path, "/attendance/student/") {
				attendanceHandler.GetAttendanceByStudent(w, r)
				return
			}
			if strings.HasPrefix(r.URL.Path, "/attendance/date/") {
				attendanceHandler.GetAttendanceByDate(w, r)
				return
			}
			http.Error(w, "not found", http.StatusNotFound)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("✅ Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
