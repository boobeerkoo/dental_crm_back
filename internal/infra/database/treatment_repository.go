package database

import (
	"github.com/GrassBusinessLabs/eduprog-go-back/internal/domain"
	"github.com/upper/db/v4"
	"time"
)

const TreatmentsTableName = "treatments"

type treatment struct {
	Id          uint64     `db:"id,omitempty"`
	PatientId   uint64     `db:"patient_id"`
	DoctorId    uint64     `db:"doctor_id"`
	Type        string     `db:"type"`
	Comment     string     `db:"comment"`
	CreatedDate time.Time  `db:"created_date,omitempty"`
	UpdatedDate time.Time  `db:"updated_date,omitempty"`
	DeletedDate *time.Time `db:"deleted_date,omitempty"`
}

type TreatmentRepository interface {
	Save(treatment domain.Treatment) (domain.Treatment, error)
	Update(treatment domain.Treatment, id uint64) (domain.Treatment, error)
	ShowList() ([]domain.Treatment, error)
	FindById(id uint64) (domain.Treatment, error)
	FindByPatientId(patientId uint64) ([]domain.Treatment, error)
	Delete(id uint64) error
}

type treatmentRepository struct {
	coll db.Collection
}

func NewTreatmentRepository(dbSession db.Session) TreatmentRepository {
	return treatmentRepository{
		coll: dbSession.Collection(TreatmentsTableName),
	}
}

func (r treatmentRepository) Save(treatment domain.Treatment) (domain.Treatment, error) {
	t := r.mapDomainToModel(treatment)
	t.Id = 0
	t.CreatedDate, t.UpdatedDate = time.Now(), time.Now()
	err := r.coll.InsertReturning(&t)
	if err != nil {
		return domain.Treatment{}, err
	}
	return r.mapModelToDomain(t), nil
}

func (r treatmentRepository) Update(treatment domain.Treatment, id uint64) (domain.Treatment, error) {
	t := r.mapDomainToModel(treatment)
	t.UpdatedDate = time.Now()
	err := r.coll.Find(db.Cond{"id": id}).Update(&t)
	if err != nil {
		return domain.Treatment{}, err
	}

	return r.mapModelToDomain(t), nil
}

func (r treatmentRepository) ShowList() ([]domain.Treatment, error) {
	var treatments []treatment

	err := r.coll.Find(db.Cond{"deleted_date": nil}).All(&treatments)
	if err != nil {
		return []domain.Treatment{}, err
	}

	return r.mapModelToDomainCollection(treatments), nil
}

func (r treatmentRepository) FindById(id uint64) (domain.Treatment, error) {
	var t treatment
	err := r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).One(&t)
	if err != nil {
		return domain.Treatment{}, err
	}

	return r.mapModelToDomain(t), nil
}

func (r treatmentRepository) FindByPatientId(patientId uint64) ([]domain.Treatment, error) {
	var treatments []treatment

	err := r.coll.Find(db.Cond{"patient_id": patientId, "deleted_date": nil}).All(&treatments)
	if err != nil {
		return []domain.Treatment{}, err
	}

	return r.mapModelToDomainCollection(treatments), nil
}

func (r treatmentRepository) Delete(id uint64) error {
	return r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).Update(map[string]interface{}{"deleted_date": time.Now()})
}

func (r treatmentRepository) mapDomainToModel(d domain.Treatment) treatment {
	return treatment{
		Id:          d.Id,
		PatientId:   d.PatientId,
		DoctorId:    d.DoctorId,
		Type:        d.Type,
		Comment:     d.Comment,
		CreatedDate: d.CreatedDate,
		UpdatedDate: d.UpdatedDate,
		DeletedDate: d.DeletedDate,
	}
}

func (r treatmentRepository) mapModelToDomain(m treatment) domain.Treatment {
	return domain.Treatment{
		Id:          m.Id,
		PatientId:   m.PatientId,
		DoctorId:    m.DoctorId,
		Type:        m.Type,
		Comment:     m.Comment,
		CreatedDate: m.CreatedDate,
		UpdatedDate: m.UpdatedDate,
		DeletedDate: m.DeletedDate,
	}
}

func (r treatmentRepository) mapModelToDomainCollection(m []treatment) []domain.Treatment {
	result := make([]domain.Treatment, len(m))

	for i := range m {
		result[i] = r.mapModelToDomain(m[i])
	}

	return result
}
