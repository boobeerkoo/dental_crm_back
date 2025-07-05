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

type AppointmentController struct {
	appointmentService app.AppointmentService
}

func NewAppointmentController(as app.AppointmentService) AppointmentController {
	return AppointmentController{
		appointmentService: as,
	}
}

func (c AppointmentController) SaveAppointment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		appointment, err := requests.Bind(r, requests.CreateAppointmentRequest{}, domain.Appointment{})
		if err != nil {
			log.Printf("AppointmentController: %s", err)
			BadRequest(w, err)
			return
		}

		appointment, err = c.appointmentService.Save(appointment)
		if err != nil {
			log.Printf("AppointmentController: %s", err)
			InternalServerError(w, err)
			return
		}

		var appointmentDto resources.AppointmentDto
		Created(w, appointmentDto.DomainToDto(appointment))
	}
}

func (c AppointmentController) UpdateAppointment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(chi.URLParam(r, "aId"), 10, 64)
		if err != nil {
			log.Printf("AppointmentController: %s", err)
			BadRequest(w, err)
			return
		}

		appointment, err := requests.Bind(r, requests.UpdateAppointmentRequest{}, domain.Appointment{})
		if err != nil {
			log.Printf("AppointmentController: %s", err)
			BadRequest(w, err)
			return
		}

		appointment, err = c.appointmentService.Update(appointment, id)
		if err != nil {
			log.Printf("AppointmentController: %s", err)
			InternalServerError(w, err)
			return
		}

		var appointmentDto resources.AppointmentDto
		Success(w, appointmentDto.DomainToDto(appointment))
	}
}

func (c AppointmentController) ShowList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		appointments, err := c.appointmentService.ShowList()
		if err != nil {
			log.Printf("AppointmentController: %s", err)
			InternalServerError(w, err)
			return
		}

		var appointmentDto resources.AppointmentDto
		Success(w, appointmentDto.DomainToDtoCollection(appointments))
	}
}

func (c AppointmentController) FindById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(chi.URLParam(r, "aId"), 10, 64)
		if err != nil {
			log.Printf("AppointmentController: %s", err)
			BadRequest(w, err)
			return
		}

		appointment, err := c.appointmentService.FindById(id)
		if err != nil {
			log.Printf("AppointmentController: %s", err)
			InternalServerError(w, err)
			return
		}

		var appointmentDto resources.AppointmentDto
		Success(w, appointmentDto.DomainToDto(appointment))
	}
}

func (c AppointmentController) FindByPatientId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		patientId, err := strconv.ParseUint(chi.URLParam(r, "pId"), 10, 64)
		if err != nil {
			log.Printf("AppointmentController: %s", err)
			BadRequest(w, err)
			return
		}

		appointments, err := c.appointmentService.FindByPatientId(patientId)
		if err != nil {
			log.Printf("AppointmentController: %s", err)
			InternalServerError(w, err)
			return
		}

		var appointmentDto resources.AppointmentDto
		Success(w, appointmentDto.DomainToDtoCollection(appointments))
	}
}

func (c AppointmentController) DeleteAppointment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(chi.URLParam(r, "aId"), 10, 64)
		if err != nil {
			log.Printf("AppointmentController: %s", err)
			BadRequest(w, err)
			return
		}

		err = c.appointmentService.Delete(id)
		if err != nil {
			log.Printf("AppointmentController: %s", err)
			InternalServerError(w, err)
			return
		}

		Ok(w)
	}
}
