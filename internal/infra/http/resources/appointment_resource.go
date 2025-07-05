package resources

import (
	"github.com/GrassBusinessLabs/eduprog-go-back/internal/domain"
	"time"
)

type AppointmentDto struct {
	Id              uint64    `json:"id"`
	PatientId       uint64    `json:"patient_id"`
	DoctorId        uint64    `json:"doctor_id"`
	AppointmentDate time.Time `json:"appointment_date"`
	Duration        string    `json:"duration"`
	Status          string    `json:"status"`
	Comment         string    `json:"comment"`
	CreatedDate     string    `json:"created_date"`
	UpdatedDate     string    `json:"updated_date"`
}

func (d AppointmentDto) DomainToDto(appointment domain.Appointment) AppointmentDto {

	return AppointmentDto{
		Id:              appointment.Id,
		PatientId:       appointment.PatientId,
		DoctorId:        appointment.DoctorId,
		AppointmentDate: appointment.AppointmentDate,
		Duration:        appointment.Duration,
		Status:          appointment.Status,
		Comment:         appointment.Comment,
		CreatedDate:     appointment.CreatedDate.Format("2006-01-02 15:04:05"),
		UpdatedDate:     appointment.UpdatedDate.Format("2006-01-02 15:04:05"),
	}
}

func (d AppointmentDto) DomainToDtoCollection(appointments []domain.Appointment) []AppointmentDto {
	result := make([]AppointmentDto, len(appointments))

	for i := range appointments {
		result[i] = d.DomainToDto(appointments[i])
	}

	return result
}
