package requests

import (
	"github.com/GrassBusinessLabs/eduprog-go-back/internal/domain"
)

type CreateTreatmentRequest struct {
	PatientId uint64 `json:"patient_id" validate:"required"`
	DoctorId  uint64 `json:"doctor_id" validate:"required"`
	Type      string `json:"type"`
	Comment   string `json:"comment" validate:"required"`
}

type UpdateTreatmentRequest struct {
	PatientId uint64 `json:"patient_id" `
	DoctorId  uint64 `json:"doctor_id" `
	Type      string `json:"type"`
	Comment   string `json:"comment"`
}

func (r CreateTreatmentRequest) ToDomainModel() (interface{}, error) {
	return domain.Treatment{
		PatientId: r.PatientId,
		DoctorId:  r.DoctorId,
		Type:      r.Type,
		Comment:   r.Comment,
	}, nil
}

func (r UpdateTreatmentRequest) ToDomainModel() (interface{}, error) {
	return domain.Treatment{
		PatientId: r.PatientId,
		DoctorId:  r.DoctorId,
		Type:      r.Type,
		Comment:   r.Comment,
	}, nil
}
