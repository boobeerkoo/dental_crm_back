package requests

import (
	"github.com/GrassBusinessLabs/eduprog-go-back/internal/domain"
)

type CreateDoctorRequest struct {
	Name           string `json:"name" validate:"required,gte=1,max=255"`
	Phone1         string `json:"phone" validate:"required,gte=1,max=20"`
	Phone2         string `json:"phone2" `
	Sex            string `json:"sex" validate:"required,gte=1,max=10"`
	Specialization string `json:"specialization"`
	UserId         uint64 `json:"user_id"`
}

type UpdateDoctorRequest struct {
	Name           string `json:"name"`
	Phone1         string `json:"phone"`
	Phone2         string `json:"phone2"`
	Sex            string `json:"sex"`
	Specialization string `json:"specialization"`
	UserId         uint64 `json:"user_id"`
}

func (r CreateDoctorRequest) ToDomainModel() (interface{}, error) {
	return domain.Doctor{
		Name:           r.Name,
		Phone1:         r.Phone1,
		Phone2:         r.Phone2,
		Specialization: r.Specialization,
		Sex:            r.Sex,
		UserId:         r.UserId,
	}, nil
}

func (r UpdateDoctorRequest) ToDomainModel() (interface{}, error) {
	return domain.Doctor{
		Name:           r.Name,
		Phone1:         r.Phone1,
		Phone2:         r.Phone2,
		Sex:            r.Sex,
		Specialization: r.Specialization,
		UserId:         r.UserId,
	}, nil

}
