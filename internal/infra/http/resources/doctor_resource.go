package resources

import (
	"github.com/GrassBusinessLabs/eduprog-go-back/internal/domain"
)

type DoctorDto struct {
	Id             uint64 `json:"id"`
	Name           string `json:"name"`
	Phone1         string `json:"phone1"`
	Phone2         string `json:"phone2"`
	Sex            string `json:"sex"`
	Specialization string `json:"specialization"`
	UserId         uint64 `json:"user_id"`
	CreatedDate    string `json:"created_date"`
	UpdatedDate    string `json:"updated_date"`
}

func (d DoctorDto) DomainToDto(doctor domain.Doctor) DoctorDto {

	return DoctorDto{
		Id:             doctor.Id,
		Name:           doctor.Name,
		Phone1:         doctor.Phone1,
		Phone2:         doctor.Phone2,
		Sex:            doctor.Sex,
		Specialization: doctor.Specialization,
		UserId:         doctor.UserId,
		CreatedDate:    doctor.CreatedDate.Format("2006-01-02 15:04:05"),
		UpdatedDate:    doctor.UpdatedDate.Format("2006-01-02 15:04:05"),
	}
}

func (d DoctorDto) DomainToDtoCollection(doctors []domain.Doctor) []DoctorDto {
	result := make([]DoctorDto, len(doctors))

	for i := range doctors {
		result[i] = d.DomainToDto(doctors[i])
	}

	return result
}
