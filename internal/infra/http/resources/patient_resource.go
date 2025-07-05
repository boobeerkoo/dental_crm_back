package resources

import (
	"github.com/GrassBusinessLabs/eduprog-go-back/internal/domain"
	"time"
)

type PatientDto struct {
	Id            uint64    `json:"id"`
	Name          string    `json:"name"`
	Phone1        string    `json:"phone1"`
	Phone2        string    `json:"phone2"`
	Address       string    `json:"address"`
	Email         string    `json:"email"`
	Sex           string    `json:"sex"`
	ImportantInfo string    `json:"important_info"`
	Comment       string    `json:"comment"`
	Status        string    `json:"status"`
	DateOfBirth   time.Time `json:"date_of_birth"`
	CreatedDate   string    `json:"created_date"`
	UpdatedDate   string    `json:"updated_date"`
}

type PatientsDto struct {
	Items []PatientDto `json:"items"`
	Total uint64       `json:"total"`
	Pages uint         `json:"pages"`
}

func (d PatientDto) DomainToDto(patient domain.Patient) PatientDto {

	return PatientDto{
		Id:            patient.Id,
		Name:          patient.Name,
		Phone1:        patient.Phone1,
		Phone2:        patient.Phone2,
		Address:       patient.Address,
		Email:         patient.Email,
		Sex:           patient.Sex,
		ImportantInfo: patient.ImportantInfo,
		Comment:       patient.Comment,
		Status:        patient.Status,
		DateOfBirth:   patient.DateOfBirth,
		CreatedDate:   patient.CreatedDate.Format("2006-01-02 15:04:05"),
		UpdatedDate:   patient.UpdatedDate.Format("2006-01-02 15:04:05"),
	}
}

func (d PatientsDto) DomainToDtoCollection(patients domain.Patients) PatientsDto {
	var items []PatientDto
	for _, patient := range patients.Items {
		items = append(items, PatientDto{}.DomainToDto(patient))
	}
	return PatientsDto{
		Items: items,
		Total: patients.Total,
		Pages: patients.Pages,
	}
}
