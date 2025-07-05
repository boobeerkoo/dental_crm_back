package domain

import "time"

type Doctor struct {
	Id             uint64
	Name           string
	Phone1         string
	Phone2         string
	Specialization string
	Sex            string
	UserId         uint64

	CreatedDate time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time
}

type Doctors struct {
	Items []Doctor
	Total uint64
	Pages uint
}
