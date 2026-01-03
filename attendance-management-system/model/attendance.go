package model

type Attendance struct {
	ID        string `json:"id"`
	StudentID string `json:"student_id"`
	Date      string `json:"date"`
	Status    string `json:"status"`
}
