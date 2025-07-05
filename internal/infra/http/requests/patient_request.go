package requests

import (
	"github.com/GrassBusinessLabs/eduprog-go-back/internal/domain"
	"time"
)

type CreatePatientRequest struct {
	Name          string    `json:"name" validate:"required,gte=1,max=255"`
	Phone1        string    `json:"phone" validate:"required,gte=1,max=20"`
	Phone2        string    `json:"phone2" `
	Address       string    `json:"address" `
	Email         string    `json:"email" `
	Sex           string    `json:"sex" validate:"required,gte=1,max=10"`
	ImportantInfo string    `json:"important_info"`
	Comment       string    `json:"comment"`
	Status        string    `json:"status"`
	DateOfBirth   time.Time `json:"date_of_birth"`
}

type UpdatePatientRequest struct {
	Name          string    `json:"name" `
	Phone1        string    `json:"phone"`
	Phone2        string    `json:"phone2"`
	Address       string    `json:"address" `
	Email         string    `json:"email" `
	Sex           string    `json:"sex" `
	ImportantInfo string    `json:"important_info"`
	Comment       string    `json:"comment"`
	Status        string    `json:"status"`
	DateOfBirth   time.Time `json:"date_of_birth"`
}

func (r CreatePatientRequest) ToDomainModel() (interface{}, error) {
	return domain.Patient{
		Name:          r.Name,
		Phone1:        r.Phone1,
		Phone2:        r.Phone2,
		Address:       r.Address,
		Email:         r.Email,
		Sex:           r.Sex,
		ImportantInfo: r.ImportantInfo,
		Comment:       r.Comment,
		Status:        r.Status,
		DateOfBirth:   r.DateOfBirth,
	}, nil
}

func (r UpdatePatientRequest) ToDomainModel() (interface{}, error) {
	return domain.Patient{
		Name:          r.Name,
		Phone1:        r.Phone1,
		Phone2:        r.Phone2,
		Address:       r.Address,
		Email:         r.Email,
		Sex:           r.Sex,
		ImportantInfo: r.ImportantInfo,
		Comment:       r.Comment,
		Status:        r.Status,
		DateOfBirth:   r.DateOfBirth,
	}, nil

}
