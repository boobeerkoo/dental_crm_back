package domain

import "time"

type Appointment struct {
	Id              uint64
	PatientId       uint64
	DoctorId        uint64
	AppointmentDate time.Time
	Duration        string
	Status          string
	Comment         string
	CreatedDate     time.Time
	UpdatedDate     time.Time
	DeletedDate     *time.Time
}

type Appointments struct {
	Items []Appointment
	Total uint64
	Pages uint
}
