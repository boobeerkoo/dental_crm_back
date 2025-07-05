package requests

import (
	"github.com/GrassBusinessLabs/eduprog-go-back/internal/domain"
	"log"
	"time"
)

type CreateAppointmentRequest struct {
	PatientId       uint64    `json:"patient_id" validate:"required"`
	DoctorId        uint64    `json:"doctor_id" validate:"required"`
	AppointmentDate time.Time `json:"appointment_date" validate:"required"`
	Duration        string    `json:"duration" validate:"required"`
	Status          string    `json:"status" validate:"required"`
	Comment         string    `json:"comment" validate:"required"`
}

type UpdateAppointmentRequest struct {
	PatientId       uint64    `json:"patient_id" `
	DoctorId        uint64    `json:"doctor_id" `
	AppointmentDate time.Time `json:"appointment_date" `
	Duration        string    `json:"duration" `
	Status          string    `json:"status" `
	Comment         string    `json:"comment"`
}

func (r CreateAppointmentRequest) ToDomainModel() (interface{}, error) {
	duration, err := time.ParseDuration(r.Duration)
	if err != nil {
		log.Printf("AppointmentRequest: %s", err)
		return domain.Appointment{}, err
	}
	return domain.Appointment{
		PatientId:       r.PatientId,
		DoctorId:        r.DoctorId,
		AppointmentDate: r.AppointmentDate,
		Duration:        duration.String(),
		Status:          r.Status,
		Comment:         r.Comment,
	}, nil
}

func (r UpdateAppointmentRequest) ToDomainModel() (interface{}, error) {
	duration, err := time.ParseDuration(r.Duration)
	if err != nil {
		log.Printf("AppointmentRequest: %s", err)
		return domain.Appointment{}, err
	}
	return domain.Appointment{

		PatientId:       r.PatientId,
		DoctorId:        r.DoctorId,
		AppointmentDate: r.AppointmentDate,
		Duration:        duration.String(),
		Status:          r.Status,
		Comment:         r.Comment,
	}, nil
}
