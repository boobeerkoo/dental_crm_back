package domain

import "time"

type Treatment struct {
	Id          uint64
	PatientId   uint64
	DoctorId    uint64
	Type        string
	Comment     string
	CreatedDate time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time
}

type Treatments struct {
	Items []Treatment
	Total uint64
	Pages uint
}
