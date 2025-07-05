package domain

import "time"

type Tooth struct {
	Id              uint64
	DentalFormulaId uint64
	ToothNumber     uint64
	ToothName       string
	ToothType       string
	Damage          string
	Parodont        string
	Endo            string
	Constructions   string
	CreatedDate     time.Time
	UpdatedDate     time.Time
	DeletedDate     *time.Time
}

type Teeth struct {
	Items []Tooth
	Total uint64
	Pages uint
}
