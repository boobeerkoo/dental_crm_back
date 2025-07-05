package domain

import "time"

type Patient struct {
	Id            uint64
	Name          string
	Phone1        string
	Phone2        string
	Address       string
	Email         string
	Sex           string
	ImportantInfo string
	Comment       string
	Status        string
	DateOfBirth   time.Time
	CreatedDate   time.Time
	UpdatedDate   time.Time
	DeletedDate   *time.Time
}

type Patients struct {
	Items []Patient
	Total uint64
	Pages uint
}
