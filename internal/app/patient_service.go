package app

import (
	"github.com/GrassBusinessLabs/eduprog-go-back/internal/domain"
	"github.com/GrassBusinessLabs/eduprog-go-back/internal/infra/database"
	"log"
)

type PatientService interface {
	Save(patient domain.Patient) (domain.Patient, error)
	Update(patient domain.Patient, id uint64) (domain.Patient, error)
	ShowList() (domain.Patients, error)
	FindById(id uint64) (domain.Patient, error)
	Delete(id uint64) error
}

type patientService struct {
	patientRepo database.PatientRepository
}

func NewPatientService(pr database.PatientRepository) PatientService {
	return patientService{
		patientRepo: pr,
	}
}

func (s patientService) Save(patient domain.Patient) (domain.Patient, error) {
	p, err := s.patientRepo.Save(patient)
	if err != nil {
		log.Printf("PatientService: %s", err)
		return domain.Patient{}, err
	}
	return p, nil
}

func (s patientService) Update(patient domain.Patient, id uint64) (domain.Patient, error) {
	p, err := s.patientRepo.Update(patient, id)

	return p, err
}

func (s patientService) ShowList() (domain.Patients, error) {
	p, err := s.patientRepo.ShowList()
	if err != nil {
		log.Printf("PatientService: %s", err)
		return domain.Patients{}, err
	}
	return p, nil
}

func (s patientService) FindById(id uint64) (domain.Patient, error) {
	p, err := s.patientRepo.FindById(id)
	if err != nil {
		log.Printf("PatientService: %s", err)
		return domain.Patient{}, err
	}
	return p, nil
}

func (s patientService) Delete(id uint64) error {
	err := s.patientRepo.Delete(id)
	if err != nil {
		log.Printf("PatientService: %s", err)
		return err
	}
	return nil
}
