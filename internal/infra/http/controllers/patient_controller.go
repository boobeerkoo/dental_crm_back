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

type PatientController struct {
	patientService app.PatientService
}

func NewPatientController(ps app.PatientService) PatientController {
	return PatientController{
		patientService: ps,
	}
}

func (c PatientController) SavePatient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		patient, err := requests.Bind(r, requests.CreatePatientRequest{}, domain.Patient{})
		if err != nil {
			log.Printf("PatientController: %s", err)
			BadRequest(w, err)
			return
		}

		patient, err = c.patientService.Save(patient)
		if err != nil {
			log.Printf("PatientController: %s", err)
			InternalServerError(w, err)
			return
		}

		var patientDto resources.PatientDto
		Created(w, patientDto.DomainToDto(patient))
	}
}

func (c PatientController) UpdatePatient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(chi.URLParam(r, "pId"), 10, 64)
		if err != nil {
			log.Printf("PatientController: %s", err)
			BadRequest(w, err)
			return
		}

		patient, err := requests.Bind(r, requests.UpdatePatientRequest{}, domain.Patient{})
		if err != nil {
			log.Printf("PatientController: %s", err)
			BadRequest(w, err)
			return
		}

		patient, err = c.patientService.Update(patient, id)
		if err != nil {
			log.Printf("PatientController: %s", err)
			InternalServerError(w, err)
			return
		}

		var patientDto resources.PatientDto
		Success(w, patientDto.DomainToDto(patient))
	}
}

func (c PatientController) ShowList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		patients, err := c.patientService.ShowList()
		if err != nil {
			log.Printf("PatientController: %s", err)
			InternalServerError(w, err)
			return
		}

		var patientsDto resources.PatientsDto
		Success(w, patientsDto.DomainToDtoCollection(patients))
	}
}

func (c PatientController) FindById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(chi.URLParam(r, "pId"), 10, 64)
		if err != nil {
			log.Printf("PatientController: %s", err)
			BadRequest(w, err)
			return
		}

		patient, err := c.patientService.FindById(id)
		if err != nil {
			log.Printf("PatientController: %s", err)
			InternalServerError(w, err)
			return
		}

		var patientDto resources.PatientDto
		Success(w, patientDto.DomainToDto(patient))
	}
}

func (c PatientController) DeletePatient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(chi.URLParam(r, "pId"), 10, 64)
		if err != nil {
			log.Printf("PatientController: %s", err)
			BadRequest(w, err)
			return
		}

		err = c.patientService.Delete(id)
		if err != nil {
			log.Printf("PatientController: %s", err)
			InternalServerError(w, err)
			return
		}

		Ok(w)
	}
}
