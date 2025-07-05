package app

import (
	"github.com/GrassBusinessLabs/eduprog-go-back/internal/domain"
	"github.com/GrassBusinessLabs/eduprog-go-back/internal/infra/database"
	"log"
)

type TreatmentService interface {
	Save(treatment domain.Treatment) (domain.Treatment, error)
	Update(treatment domain.Treatment, id uint64) (domain.Treatment, error)
	ShowList() ([]domain.Treatment, error)
	FindById(id uint64) (domain.Treatment, error)
	FindByPatientId(patientId uint64) ([]domain.Treatment, error)
	Delete(id uint64) error
}
type treatmentService struct {
	treatmentRepo database.TreatmentRepository
}

func NewTreatmentService(tr database.TreatmentRepository) TreatmentService {
	return treatmentService{
		treatmentRepo: tr,
	}
}

func (s treatmentService) Save(treatment domain.Treatment) (domain.Treatment, error) {
	a, err := s.treatmentRepo.Save(treatment)
	if err != nil {
		log.Printf("TreatmentService: %s", err)
		return domain.Treatment{}, err
	}
	return a, nil
}

func (s treatmentService) Update(treatment domain.Treatment, id uint64) (domain.Treatment, error) {
	a, err := s.treatmentRepo.Update(treatment, id)
	if err != nil {
		log.Printf("TreatmentService: %s", err)
		return domain.Treatment{}, err
	}
	return a, nil
}

func (s treatmentService) ShowList() ([]domain.Treatment, error) {
	a, err := s.treatmentRepo.ShowList()
	if err != nil {
		log.Printf("TreatmentService: %s", err)
		return []domain.Treatment{}, err
	}
	return a, nil
}

func (s treatmentService) FindById(id uint64) (domain.Treatment, error) {
	a, err := s.treatmentRepo.FindById(id)
	if err != nil {
		log.Printf("TreatmentService: %s", err)
		return domain.Treatment{}, err
	}
	return a, nil
}

func (s treatmentService) FindByPatientId(patientId uint64) ([]domain.Treatment, error) {
	a, err := s.treatmentRepo.FindByPatientId(patientId)
	if err != nil {
		log.Printf("TreatmentService: %s", err)
		return []domain.Treatment{}, err
	}
	return a, nil
}

func (s treatmentService) Delete(id uint64) error {
	err := s.treatmentRepo.Delete(id)
	if err != nil {
		log.Printf("TreatmentService: %s", err)
		return err
	}
	return nil
}
