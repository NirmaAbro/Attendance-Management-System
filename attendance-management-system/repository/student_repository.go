package repository

import (
	"errors"
	"sync"

	"attendance-management-system/model"
)

type StudentRepository interface {
	Create(student model.Student) error
	GetAll() []model.Student
	GetByID(id string) (model.Student, error)
}

type InMemoryStudentRepository struct {
	data map[string]model.Student
	mu   sync.RWMutex
}

func NewStudentRepository() StudentRepository {
	return &InMemoryStudentRepository{
		data: make(map[string]model.Student),
	}
}

func (r *InMemoryStudentRepository) Create(student model.Student) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.data[student.ID] = student
	return nil
}

func (r *InMemoryStudentRepository) GetAll() []model.Student {
	r.mu.RLock()
	defer r.mu.RUnlock()

	students := []model.Student{}
	for _, s := range r.data {
		students = append(students, s)
	}
	return students
}

func (r *InMemoryStudentRepository) GetByID(id string) (model.Student, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	student, ok := r.data[id]
	if !ok {
		return model.Student{}, errors.New("student not found")
	}
	return student, nil
}
