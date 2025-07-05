package database

import (
	"github.com/GrassBusinessLabs/eduprog-go-back/internal/domain"
	"github.com/upper/db/v4"
	"time"
)

const AppointmentsTableName = "appointments"

type appointment struct {
	Id              uint64     `db:"id,omitempty"`
	PatientId       uint64     `db:"patient_id"`
	DoctorId        uint64     `db:"doctor_id"`
	AppointmentDate time.Time  `db:"appointment_date"`
	Duration        string     `db:"duration"`
	Status          string     `db:"status"`
	Comment         string     `db:"comment"`
	CreatedDate     time.Time  `db:"created_date,omitempty"`
	UpdatedDate     time.Time  `db:"updated_date,omitempty"`
	DeletedDate     *time.Time `db:"deleted_date,omitempty"`
}

type AppointmentRepository interface {
	Save(appointment domain.Appointment) (domain.Appointment, error)
	Update(appointment domain.Appointment, id uint64) (domain.Appointment, error)
	ShowList() ([]domain.Appointment, error)
	FindById(id uint64) (domain.Appointment, error)
	FindByPatientId(patientId uint64) ([]domain.Appointment, error)
	Delete(id uint64) error
}

type appointmentRepository struct {
	coll db.Collection
}

func NewAppointmentRepository(dbSession db.Session) AppointmentRepository {
	return appointmentRepository{
		coll: dbSession.Collection(AppointmentsTableName),
	}
}

func (r appointmentRepository) Save(appointment domain.Appointment) (domain.Appointment, error) {
	a := r.mapDomainToModel(appointment)
	a.Id = 0
	a.CreatedDate, a.UpdatedDate = time.Now(), time.Now()
	err := r.coll.InsertReturning(&a)
	if err != nil {
		return domain.Appointment{}, err
	}
	return r.mapModelToDomain(a), nil
}

func (r appointmentRepository) Update(appointment domain.Appointment, id uint64) (domain.Appointment, error) {
	a := r.mapDomainToModel(appointment)
	a.UpdatedDate = time.Now()
	err := r.coll.Find(db.Cond{"id": id}).Update(&a)
	if err != nil {
		return domain.Appointment{}, err
	}

	return r.mapModelToDomain(a), nil
}

func (r appointmentRepository) ShowList() ([]domain.Appointment, error) {
	var appointments []appointment

	err := r.coll.Find(db.Cond{"deleted_date": nil}).All(&appointments)
	if err != nil {
		return []domain.Appointment{}, err
	}

	return r.mapModelToDomainCollection(appointments), nil
}

func (r appointmentRepository) FindById(id uint64) (domain.Appointment, error) {
	var a appointment
	err := r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).One(&a)
	if err != nil {
		return domain.Appointment{}, err
	}

	return r.mapModelToDomain(a), nil
}

func (r appointmentRepository) FindByPatientId(patientId uint64) ([]domain.Appointment, error) {
	var appointments []appointment

	err := r.coll.Find(db.Cond{"patient_id": patientId, "deleted_date": nil}).All(&appointments)
	if err != nil {
		return []domain.Appointment{}, err
	}

	return r.mapModelToDomainCollection(appointments), nil
}

func (r appointmentRepository) Delete(id uint64) error {
	return r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).Update(map[string]interface{}{"deleted_date": time.Now()})
}

func (r appointmentRepository) mapDomainToModel(d domain.Appointment) appointment {
	return appointment{
		Id:              d.Id,
		PatientId:       d.PatientId,
		DoctorId:        d.DoctorId,
		AppointmentDate: d.AppointmentDate,
		Duration:        d.Duration,
		Status:          d.Status,
		Comment:         d.Comment,
		CreatedDate:     d.CreatedDate,
		UpdatedDate:     d.UpdatedDate,
		DeletedDate:     d.DeletedDate,
	}
}

func (r appointmentRepository) mapModelToDomain(m appointment) domain.Appointment {
	return domain.Appointment{
		Id:              m.Id,
		PatientId:       m.PatientId,
		DoctorId:        m.DoctorId,
		AppointmentDate: m.AppointmentDate,
		Duration:        m.Duration,
		Status:          m.Status,
		Comment:         m.Comment,
		CreatedDate:     m.CreatedDate,
		UpdatedDate:     m.UpdatedDate,
		DeletedDate:     m.DeletedDate,
	}
}

func (r appointmentRepository) mapModelToDomainCollection(m []appointment) []domain.Appointment {
	result := make([]domain.Appointment, len(m))

	for i := range m {
		result[i] = r.mapModelToDomain(m[i])
	}

	return result
}
