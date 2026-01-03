package repository

import (
	"sync"

	"attendance-management-system/model"
)

type AttendanceRepository interface {
	Create(attendance model.Attendance) error
	GetAll() []model.Attendance
	GetByStudentID(studentID string) []model.Attendance
}

type InMemoryAttendanceRepository struct {
	data []model.Attendance
	mu   sync.RWMutex
}

func NewAttendanceRepository() AttendanceRepository {
	return &InMemoryAttendanceRepository{
		data: []model.Attendance{},
	}
}

func (r *InMemoryAttendanceRepository) Create(attendance model.Attendance) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.data = append(r.data, attendance)
	return nil
}

func (r *InMemoryAttendanceRepository) GetAll() []model.Attendance {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.data
}

func (r *InMemoryAttendanceRepository) GetByStudentID(studentID string) []model.Attendance {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := []model.Attendance{}
	for _, a := range r.data {
		if a.StudentID == studentID {
			result = append(result, a)
		}
	}
	return result
}
