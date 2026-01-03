# Attendance Management System

A web-based **Attendance Management System** built using **Go (Golang)** and **RESTful APIs**. This project follows **clean architecture principles** and is designed for educational purposes to demonstrate core concepts of **web development** and **distributed systems**.

---

## 📌 Project Overview

The Attendance Management System allows instructors to manage students and mark their attendance through HTTP-based REST APIs. The system operates on a **client–server model**, supports multiple clients (browser, Postman, API tools), and exchanges data in **JSON format**. Data is stored in an **in-memory repository**, making the system lightweight and easy to deploy.

---

## 🎯 Objectives

* Develop a web-based attendance system
* Provide functionality via RESTful HTTP endpoints
* Support multiple clients (browser, API tools)
* Apply clean architecture & design patterns
* Demonstrate core concepts of web and distributed systems

---

## 🏗️ Architecture

### Technology Stack

* **Backend:** Go (Golang)
* **HTTP Server:** net/http
* **API Style:** REST
* **Data Format:** JSON
* **Storage:** In-memory
* **Deployment Model:** Client–Server

### Project Layers

* **Model:** Defines core entities (Student, Attendance)
* **Repository:** Data abstraction using Repository Pattern
* **Service:** Business logic layer
* **Handler:** HTTP request & response handling
* **Main:** Server initialization and dependency injection

This layered structure ensures **loose coupling**, **maintainability**, and **testability**.

---

## 📂 Folder Structure

```
attendance-management-system/
│
├── handler/        # HTTP handlers
├── service/        # Business logic
├── repository/     # In-memory data storage
├── model/          # Entity definitions
├── main.go         # Server entry point
├── go.mod
└── README.md
```

---

## 🔗 REST API Endpoints

### Student APIs

* `POST   /students` – Create student
* `GET    /students` – Get all students
* `GET    /students/{id}` – Get student by ID
* `PUT    /students/{id}` – Update student
* `DELETE /students/{id}` – Delete student

### Attendance APIs

* `POST   /attendance` – Mark attendance
* `GET    /attendance` – Get all attendance records
* `GET    /attendance/student/{studentId}` – Attendance by student
* `GET    /attendance/date/{date}` – Attendance by date
* `PUT    /attendance/{id}` – Update attendance
* `DELETE /attendance/{id}` – Delete attendance

---

## 🧪 Testing

All APIs were tested using **Postman**. CRUD operations for both students and attendance work correctly, including reporting endpoints by student and date.

---

## 🚀 How to Run the Project

### Prerequisites

* Go installed (v1.20+ recommended)

### Steps

```bash
git clone <your-repo-url>
cd attendance-management-system
go run main.go
```

The server will start at:

```
http://localhost:8080
```

Use Postman or any HTTP client to test the APIs.

---

## 📚 Course Topics Covered

* Web development with Go
* Distributed systems
* HTTP & RESTful APIs
* Go modules & interfaces
* Repository pattern
* Clean architecture
* Separation of concerns

---

## ⚠️ Limitations

* In-memory storage (data resets on server restart)
* No authentication or authorization
* No frontend UI (API-based system)

These limitations are acceptable as the project is intended for academic learning.

---

## ✅ Conclusion

This project fulfills all requirements of a distributed web-based Attendance Management System. It demonstrates proper REST API design, clean architecture, and software engineering principles, making it suitable for academic evaluation and learning purposes.

---

## 👤 Author

**Nirma Abro**

---

## 📄 License

This project is for educational purposes only.
