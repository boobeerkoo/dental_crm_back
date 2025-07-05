package database

import (
	"github.com/GrassBusinessLabs/eduprog-go-back/internal/domain"
	"github.com/upper/db/v4"
	"time"
)

const PatientsTableName = "patients"

type patient struct {
	Id            uint64     `db:"id,omitempty"`
	Name          string     `db:"name"`
	Phone1        string     `db:"phone1"`
	Phone2        string     `db:"phone2"`
	Address       string     `db:"address"`
	Email         string     `db:"email"`
	Sex           string     `db:"sex"`
	ImportantInfo string     `db:"important_info"`
	Comment       string     `db:"comment"`
	Status        string     `db:"status"`
	DateOfBirth   time.Time  `db:"date_of_birth"`
	CreatedDate   time.Time  `db:"created_date,omitempty"`
	UpdatedDate   time.Time  `db:"updated_date,omitempty"`
	DeletedDate   *time.Time `db:"deleted_date,omitempty"`
}

type PatientRepository interface {
	Save(patient domain.Patient) (domain.Patient, error)
	Update(patient domain.Patient, id uint64) (domain.Patient, error)
	ShowList() (domain.Patients, error)
	FindById(id uint64) (domain.Patient, error)
	Delete(id uint64) error
}

type patientRepository struct {
	coll db.Collection
}

func NewPatientRepository(dbSession db.Session) PatientRepository {
	return patientRepository{
		coll: dbSession.Collection(PatientsTableName),
	}
}

func (r patientRepository) Save(patient domain.Patient) (domain.Patient, error) {
	p := r.mapDomainToModel(patient)
	p.Id = 0
	p.CreatedDate, p.UpdatedDate = time.Now(), time.Now()
	err := r.coll.InsertReturning(&p)
	if err != nil {
		return domain.Patient{}, err
	}
	return r.mapModelToDomain(p), nil
}

func (r patientRepository) Update(patient domain.Patient, id uint64) (domain.Patient, error) {
	e := r.mapDomainToModel(patient)
	e.UpdatedDate = time.Now()
	err := r.coll.Find(db.Cond{"id": id}).Update(&e)
	if err != nil {
		return domain.Patient{}, err
	}

	return r.mapModelToDomain(e), nil
}

func (r patientRepository) ShowList() (domain.Patients, error) {
	var patientsSlice []patient
	var patients domain.Patients

	err := r.coll.Find(db.Cond{"deleted_date": nil}).All(&patientsSlice)
	if err != nil {
		return domain.Patients{}, err
	}

	for i := range patientsSlice {
		patients.Items = append(patients.Items, r.mapModelToDomain(patientsSlice[i]))
	}
	patients.Total = uint64(len(patientsSlice))

	return patients, nil
}

func (r patientRepository) FindById(id uint64) (domain.Patient, error) {
	var p patient
	err := r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).One(&p)
	if err != nil {
		return domain.Patient{}, err
	}

	return r.mapModelToDomain(p), nil
}

func (r patientRepository) Delete(id uint64) error {
	return r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).Update(map[string]interface{}{"deleted_date": time.Now()})
}

func (r patientRepository) mapDomainToModel(d domain.Patient) patient {
	return patient{
		Id:            d.Id,
		Name:          d.Name,
		Phone1:        d.Phone1,
		Phone2:        d.Phone2,
		Address:       d.Address,
		Email:         d.Email,
		Sex:           d.Sex,
		ImportantInfo: d.ImportantInfo,
		Comment:       d.Comment,
		Status:        d.Status,
		DateOfBirth:   d.DateOfBirth,
		CreatedDate:   d.CreatedDate,
		UpdatedDate:   d.UpdatedDate,
		DeletedDate:   d.DeletedDate,
	}
}

func (r patientRepository) mapModelToDomain(m patient) domain.Patient {
	return domain.Patient{
		Id:            m.Id,
		Name:          m.Name,
		Phone1:        m.Phone1,
		Phone2:        m.Phone2,
		Address:       m.Address,
		Email:         m.Email,
		Sex:           m.Sex,
		ImportantInfo: m.ImportantInfo,
		Comment:       m.Comment,
		Status:        m.Status,
		DateOfBirth:   m.DateOfBirth,
		CreatedDate:   m.CreatedDate,
		UpdatedDate:   m.UpdatedDate,
		DeletedDate:   m.DeletedDate,
	}
}
