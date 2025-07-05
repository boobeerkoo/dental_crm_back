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

type DoctorController struct {
	doctorService app.DoctorService
}

func NewDoctorController(ds app.DoctorService) DoctorController {
	return DoctorController{
		doctorService: ds,
	}
}

func (c DoctorController) SaveDoctor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		doctor, err := requests.Bind(r, requests.CreateDoctorRequest{}, domain.Doctor{})
		if err != nil {
			log.Printf("DoctorController: %s", err)
			BadRequest(w, err)
			return
		}

		doctor, err = c.doctorService.Save(doctor)
		if err != nil {
			log.Printf("DoctorController: %s", err)
			InternalServerError(w, err)
			return
		}

		var doctorDto resources.DoctorDto
		Created(w, doctorDto.DomainToDto(doctor))
	}
}

func (c DoctorController) UpdateDoctor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(chi.URLParam(r, "dId"), 10, 64)
		if err != nil {
			log.Printf("DoctorController: %s", err)
			BadRequest(w, err)
			return
		}

		doctor, err := requests.Bind(r, requests.UpdateDoctorRequest{}, domain.Doctor{})
		if err != nil {
			log.Printf("DoctorController: %s", err)
			BadRequest(w, err)
			return
		}

		doctor, err = c.doctorService.Update(doctor, id)
		if err != nil {
			log.Printf("DoctorController: %s", err)
			InternalServerError(w, err)
			return
		}

		var doctorDto resources.DoctorDto
		Success(w, doctorDto.DomainToDto(doctor))
	}
}

func (c DoctorController) ShowList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		doctors, err := c.doctorService.ShowList()
		if err != nil {
			log.Printf("DoctorController: %s", err)
			InternalServerError(w, err)
			return
		}

		var doctorDto resources.DoctorDto
		Success(w, doctorDto.DomainToDtoCollection(doctors))
	}
}

func (c DoctorController) FindById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(chi.URLParam(r, "dId"), 10, 64)
		if err != nil {
			log.Printf("DoctorController: %s", err)
			BadRequest(w, err)
			return
		}

		doctor, err := c.doctorService.FindById(id)
		if err != nil {
			log.Printf("DoctorController: %s", err)
			InternalServerError(w, err)
			return
		}

		var doctorDto resources.DoctorDto
		Success(w, doctorDto.DomainToDto(doctor))
	}
}

func (c DoctorController) DeleteDoctor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(chi.URLParam(r, "dId"), 10, 64)
		if err != nil {
			log.Printf("DoctorController: %s", err)
			BadRequest(w, err)
			return
		}

		err = c.doctorService.Delete(id)
		if err != nil {
			log.Printf("DoctorController: %s", err)
			InternalServerError(w, err)
			return
		}

		Ok(w)
	}
}
