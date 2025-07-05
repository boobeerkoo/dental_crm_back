package database

import (
	"github.com/GrassBusinessLabs/eduprog-go-back/internal/domain"
	"github.com/upper/db/v4"
	"time"
)

const DoctorsTableName = "doctors"

type doctor struct {
	Id             uint64     `db:"id,omitempty"`
	Name           string     `db:"name"`
	Phone1         string     `db:"phone1"`
	Phone2         string     `db:"phone2"`
	Sex            string     `db:"sex"`
	Specialization string     `db:"specialization"`
	UserId         uint64     `db:"user_id"`
	CreatedDate    time.Time  `db:"created_date,omitempty"`
	UpdatedDate    time.Time  `db:"updated_date,omitempty"`
	DeletedDate    *time.Time `db:"deleted_date,omitempty"`
}

type DoctorRepository interface {
	Save(doctor domain.Doctor) (domain.Doctor, error)
	Update(doctor domain.Doctor, id uint64) (domain.Doctor, error)
	ShowList() ([]domain.Doctor, error)
	FindById(id uint64) (domain.Doctor, error)
	Delete(id uint64) error
}

type doctorRepository struct {
	coll db.Collection
}

func NewDoctorRepository(dbSession db.Session) DoctorRepository {
	return doctorRepository{
		coll: dbSession.Collection(DoctorsTableName),
	}
}

func (r doctorRepository) Save(doctor domain.Doctor) (domain.Doctor, error) {
	p := r.mapDomainToModel(doctor)
	p.Id = 0
	p.CreatedDate, p.UpdatedDate = time.Now(), time.Now()
	err := r.coll.InsertReturning(&p)
	if err != nil {
		return domain.Doctor{}, err
	}
	return r.mapModelToDomain(p), nil
}

func (r doctorRepository) Update(doctor domain.Doctor, id uint64) (domain.Doctor, error) {
	e := r.mapDomainToModel(doctor)
	e.UpdatedDate = time.Now()
	err := r.coll.Find(db.Cond{"id": id}).Update(&e)
	if err != nil {
		return domain.Doctor{}, err
	}

	return r.mapModelToDomain(e), nil
}

func (r doctorRepository) ShowList() ([]domain.Doctor, error) {
	var doctors []doctor

	err := r.coll.Find(db.Cond{"deleted_date": nil}).All(&doctors)
	if err != nil {
		return []domain.Doctor{}, err
	}

	return r.mapModelToDomainCollection(doctors), nil
}

func (r doctorRepository) FindById(id uint64) (domain.Doctor, error) {
	var d doctor
	err := r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).One(&d)
	if err != nil {
		return domain.Doctor{}, err
	}

	return r.mapModelToDomain(d), nil
}

func (r doctorRepository) Delete(id uint64) error {
	return r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).Update(map[string]interface{}{"deleted_date": time.Now()})
}

func (r doctorRepository) mapDomainToModel(d domain.Doctor) doctor {
	return doctor{
		Id:             d.Id,
		Name:           d.Name,
		Phone1:         d.Phone1,
		Phone2:         d.Phone2,
		Sex:            d.Sex,
		Specialization: d.Specialization,
		UserId:         d.UserId,
		CreatedDate:    d.CreatedDate,
		UpdatedDate:    d.UpdatedDate,
		DeletedDate:    d.DeletedDate,
	}
}

func (r doctorRepository) mapModelToDomain(m doctor) domain.Doctor {
	return domain.Doctor{
		Id:             m.Id,
		Name:           m.Name,
		Phone1:         m.Phone1,
		Phone2:         m.Phone2,
		Sex:            m.Sex,
		Specialization: m.Specialization,
		UserId:         m.UserId,
		CreatedDate:    m.CreatedDate,
		UpdatedDate:    m.UpdatedDate,
		DeletedDate:    m.DeletedDate,
	}
}

func (r doctorRepository) mapModelToDomainCollection(m []doctor) []domain.Doctor {
	result := make([]domain.Doctor, len(m))

	for i := range m {
		result[i] = r.mapModelToDomain(m[i])
	}

	return result
}
