package app

import (
	"github.com/GrassBusinessLabs/eduprog-go-back/internal/domain"
	"github.com/GrassBusinessLabs/eduprog-go-back/internal/infra/database"
	"log"
)

type AppointmentService interface {
	Save(appointment domain.Appointment) (domain.Appointment, error)
	Update(appointment domain.Appointment, id uint64) (domain.Appointment, error)
	ShowList() ([]domain.Appointment, error)
	FindById(id uint64) (domain.Appointment, error)
	FindByPatientId(patientId uint64) ([]domain.Appointment, error)
	Delete(id uint64) error
}

type appointmentService struct {
	appointmentRepo database.AppointmentRepository
}

func NewAppointmentService(ar database.AppointmentRepository) AppointmentService {
	return appointmentService{
		appointmentRepo: ar,
	}
}

func (s appointmentService) Save(appointment domain.Appointment) (domain.Appointment, error) {
	a, err := s.appointmentRepo.Save(appointment)
	if err != nil {
		log.Printf("AppointmentService: %s", err)
		return domain.Appointment{}, err
	}
	return a, nil
}

func (s appointmentService) Update(appointment domain.Appointment, id uint64) (domain.Appointment, error) {
	a, err := s.appointmentRepo.Update(appointment, id)
	if err != nil {
		log.Printf("AppointmentService: %s", err)
		return domain.Appointment{}, err
	}
	return a, nil
}

func (s appointmentService) ShowList() ([]domain.Appointment, error) {
	a, err := s.appointmentRepo.ShowList()
	if err != nil {
		log.Printf("AppointmentService: %s", err)
		return []domain.Appointment{}, err
	}
	return a, nil
}

func (s appointmentService) FindById(id uint64) (domain.Appointment, error) {
	a, err := s.appointmentRepo.FindById(id)
	if err != nil {
		log.Printf("AppointmentService: %s", err)
		return domain.Appointment{}, err
	}
	return a, nil
}

func (s appointmentService) FindByPatientId(patientId uint64) ([]domain.Appointment, error) {
	a, err := s.appointmentRepo.FindByPatientId(patientId)
	if err != nil {
		log.Printf("AppointmentService: %s", err)
		return []domain.Appointment{}, err
	}
	return a, nil
}

func (s appointmentService) Delete(id uint64) error {
	err := s.appointmentRepo.Delete(id)
	if err != nil {
		log.Printf("AppointmentService: %s", err)
		return err
	}
	return nil
}
