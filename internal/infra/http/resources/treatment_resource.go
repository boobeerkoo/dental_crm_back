package resources

import "github.com/GrassBusinessLabs/eduprog-go-back/internal/domain"

type TreatmentDto struct {
	Id          uint64 `json:"id"`
	PatientId   uint64 `json:"patient_id"`
	DoctorId    uint64 `json:"doctor_id"`
	Type        string `json:"type"`
	Comment     string `json:"comment"`
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`
}

func (d TreatmentDto) DomainToDto(treatment domain.Treatment) TreatmentDto {

	return TreatmentDto{
		Id:          treatment.Id,
		PatientId:   treatment.PatientId,
		DoctorId:    treatment.DoctorId,
		Type:        treatment.Type,
		Comment:     treatment.Comment,
		CreatedDate: treatment.CreatedDate.Format("2006-01-02 15:04:05"),
		UpdatedDate: treatment.UpdatedDate.Format("2006-01-02 15:04:05"),
	}
}

func (d TreatmentDto) DomainToDtoCollection(treatments []domain.Treatment) []TreatmentDto {
	result := make([]TreatmentDto, len(treatments))

	for i := range treatments {
		result[i] = d.DomainToDto(treatments[i])
	}

	return result
}
