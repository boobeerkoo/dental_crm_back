package controllers

import (
	"github.com/GrassBusinessLabs/eduprog-go-back/internal/app"
	"github.com/GrassBusinessLabs/eduprog-go-back/internal/domain"
	"github.com/GrassBusinessLabs/eduprog-go-back/internal/infra/http/requests"
	"github.com/GrassBusinessLabs/eduprog-go-back/internal/infra/http/resources"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

type TreatmentController struct {
	treatmentService app.TreatmentService
}

func NewTreatmentController(ts app.TreatmentService) TreatmentController {
	return TreatmentController{
		treatmentService: ts,
	}
}

func (c TreatmentController) SaveTreatment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		treatment, err := requests.Bind(r, requests.CreateTreatmentRequest{}, domain.Treatment{})
		if err != nil {
			log.Printf("TreatmentController: %s", err)
			BadRequest(w, err)
			return
		}

		treatment, err = c.treatmentService.Save(treatment)
		if err != nil {
			log.Printf("TreatmentController: %s", err)
			InternalServerError(w, err)
			return
		}

		var treatmentDto resources.TreatmentDto
		Created(w, treatmentDto.DomainToDto(treatment))
	}
}

func (c TreatmentController) UpdateTreatment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(chi.URLParam(r, "tId"), 10, 64)
		if err != nil {
			log.Printf("TreatmentController: %s", err)
			BadRequest(w, err)
			return
		}

		treatment, err := requests.Bind(r, requests.UpdateTreatmentRequest{}, domain.Treatment{})
		if err != nil {
			log.Printf("TreatmentController: %s", err)
			BadRequest(w, err)
			return
		}

		treatment, err = c.treatmentService.Update(treatment, id)
		if err != nil {
			log.Printf("TreatmentController: %s", err)
			InternalServerError(w, err)
			return
		}

		var treatmentDto resources.TreatmentDto
		Success(w, treatmentDto.DomainToDto(treatment))
	}
}

func (c TreatmentController) ShowList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		treatments, err := c.treatmentService.ShowList()
		if err != nil {
			log.Printf("TreatmentController: %s", err)
			InternalServerError(w, err)
			return
		}

		var treatmentDto resources.TreatmentDto
		Success(w, treatmentDto.DomainToDtoCollection(treatments))
	}
}

func (c TreatmentController) FindById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(chi.URLParam(r, "tId"), 10, 64)
		if err != nil {
			log.Printf("TreatmentController: %s", err)
			BadRequest(w, err)
			return
		}

		treatment, err := c.treatmentService.FindById(id)
		if err != nil {
			log.Printf("TreatmentController: %s", err)
			InternalServerError(w, err)
			return
		}

		var treatmentsDto resources.TreatmentDto
		Success(w, treatmentsDto.DomainToDto(treatment))
	}
}

func (c TreatmentController) FindByPatientId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		patientId, err := strconv.ParseUint(chi.URLParam(r, "pId"), 10, 64)
		if err != nil {
			log.Printf("TreatmentController: %s", err)
			BadRequest(w, err)
			return
		}

		treatments, err := c.treatmentService.FindByPatientId(patientId)
		if err != nil {
			log.Printf("TreatmentController: %s", err)
			InternalServerError(w, err)
			return
		}

		var treatmentDto resources.TreatmentDto
		Success(w, treatmentDto.DomainToDtoCollection(treatments))
	}
}

func (c TreatmentController) DeleteTreatment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(chi.URLParam(r, "tId"), 10, 64)
		if err != nil {
			log.Printf("TreatmentController: %s", err)
			BadRequest(w, err)
			return
		}

		err = c.treatmentService.Delete(id)
		if err != nil {
			log.Printf("TreatmentController: %s", err)
			InternalServerError(w, err)
			return
		}

		Ok(w)
	}
}
