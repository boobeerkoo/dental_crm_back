package domain

import "time"

type DentalFormula struct {
	Id          uint64
	PatientId   uint64
	DoctorId    uint64
	CreatedDate time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time
}

type DentalFormulas struct {
	Items []DentalFormula
	Total uint64
	Pages uint
}
