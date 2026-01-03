package service

import (
	"attendance-management-system/model"
	"attendance-management-system/repository"

	"github.com/google/uuid"
)

type StudentService struct {
	repo repository.StudentRepository
}

func NewStudentService(repo repository.StudentRepository) *StudentService {
	return &StudentService{repo: repo}
}

func (s *StudentService) CreateStudent(name, email string) (model.Student, error) {
	student := model.Student{
		ID:    uuid.New().String(),
		Name:  name,
		Email: email,
	}
	err := s.repo.Create(student)
	return student, err
}

func (s *StudentService) GetAllStudents() []model.Student {
	return s.repo.GetAll()
}

func (s *StudentService) GetStudentByID(id string) (model.Student, error) {
	return s.repo.GetByID(id)
}
