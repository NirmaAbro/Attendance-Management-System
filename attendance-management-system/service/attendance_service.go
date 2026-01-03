package service

import (
	"errors"

	"attendance-management-system/model"
	"attendance-management-system/repository"

	"github.com/google/uuid"
)

type AttendanceService struct {
	attendanceRepo repository.AttendanceRepository
	studentRepo    repository.StudentRepository
}

func NewAttendanceService(aRepo repository.AttendanceRepository, sRepo repository.StudentRepository) *AttendanceService {
	return &AttendanceService{
		attendanceRepo: aRepo,
		studentRepo:    sRepo,
	}
}

func (s *AttendanceService) MarkAttendance(studentID, date, status string) (model.Attendance, error) {
	_, err := s.studentRepo.GetByID(studentID)
	if err != nil {
		return model.Attendance{}, errors.New("student does not exist")
	}

	attendance := model.Attendance{
		ID:        uuid.New().String(),
		StudentID: studentID,
		Date:      date,
		Status:    status,
	}

	err = s.attendanceRepo.Create(attendance)
	return attendance, err
}

func (s *AttendanceService) GetAllAttendance() []model.Attendance {
	return s.attendanceRepo.GetAll()
}

func (s *AttendanceService) GetAttendanceByStudentID(studentID string) []model.Attendance {
	return s.attendanceRepo.GetByStudentID(studentID)
}

func (s *AttendanceService) GetAttendanceByDate(date string) []model.Attendance {
	return s.attendanceRepo.GetByDate(date)
}

func (s *AttendanceService) UpdateAttendance(id, studentID, date, status string) (model.Attendance, error) {
	attendance := model.Attendance{
		ID:        id,
		StudentID: studentID,
		Date:      date,
		Status:    status,
	}
	err := s.attendanceRepo.Update(attendance)
	return attendance, err
}

func (s *AttendanceService) DeleteAttendance(id string) error {
	return s.attendanceRepo.Delete(id)
}
