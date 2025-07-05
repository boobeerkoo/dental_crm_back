package app

import (
	"github.com/GrassBusinessLabs/eduprog-go-back/internal/domain"
	"github.com/GrassBusinessLabs/eduprog-go-back/internal/infra/database"
	"log"
)

type DoctorService interface {
	Save(doctor domain.Doctor) (domain.Doctor, error)
	Update(doctor domain.Doctor, id uint64) (domain.Doctor, error)
	ShowList() ([]domain.Doctor, error)
	FindById(id uint64) (domain.Doctor, error)
	Delete(id uint64) error
}

type doctorService struct {
	doctorRepo database.DoctorRepository
}

func NewDoctorService(dr database.DoctorRepository) DoctorService {
	return doctorService{
		doctorRepo: dr,
	}
}

func (s doctorService) Save(doctor domain.Doctor) (domain.Doctor, error) {
	p, err := s.doctorRepo.Save(doctor)
	if err != nil {
		log.Printf("DoctorService: %s", err)
		return domain.Doctor{}, err
	}
	return p, nil
}

func (s doctorService) Update(doctor domain.Doctor, id uint64) (domain.Doctor, error) {
	d, err := s.doctorRepo.Update(doctor, id)
	if err != nil {
		log.Printf("DoctorService: %s", err)
		return domain.Doctor{}, err
	}
	return d, err
}

func (s doctorService) ShowList() ([]domain.Doctor, error) {
	d, err := s.doctorRepo.ShowList()
	if err != nil {
		log.Printf("DoctorService: %s", err)
		return []domain.Doctor{}, err
	}
	return d, nil
}

func (s doctorService) FindById(id uint64) (domain.Doctor, error) {
	d, err := s.doctorRepo.FindById(id)
	if err != nil {
		log.Printf("DoctorService: %s", err)
		return domain.Doctor{}, err
	}
	return d, nil
}

func (s doctorService) Delete(id uint64) error {
	err := s.doctorRepo.Delete(id)
	if err != nil {
		log.Printf("DoctorService: %s", err)
		return err
	}
	return nil
}
