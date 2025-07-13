package database

import (
	"github.com/GrassBusinessLabs/eduprog-go-back/internal/domain"
	"github.com/upper/db/v4"
	"time"
)

const UsersTableName = "users"

type user struct {
	Id          uint64     `db:"id,omitempty"`
	Name        string     `db:"name"`
	Email       string     `db:"email"`
	Password    string     `db:"password"`
	CreatedDate time.Time  `db:"created_date,omitempty"`
	UpdatedDate time.Time  `db:"updated_date,omitempty"`
	DeletedDate *time.Time `db:"deleted_date,omitempty"`
}

type UserRepository interface {
	FindByEmail(email string) (domain.User, error)
	Save(user domain.User) (domain.User, error)
	FindById(id uint64) (domain.User, error)
	ShowList() ([]domain.User, error)
	Update(user domain.User) (domain.User, error)
	Delete(id uint64) error
}

type userRepository struct {
	coll db.Collection
}

func NewUserRepository(dbSession db.Session) UserRepository {
	return userRepository{
		coll: dbSession.Collection(UsersTableName),
	}
}

func (r userRepository) FindByEmail(email string) (domain.User, error) {
	var u user
	err := r.coll.Find(db.Cond{"email": email, "deleted_date": nil}).One(&u)
	if err != nil {
		return domain.User{}, err
	}

	return r.mapModelToDomain(u), nil
}

func (r userRepository) Save(user domain.User) (domain.User, error) {
	u := r.mapDomainToModel(user)
	u.CreatedDate, u.UpdatedDate = time.Now(), time.Now()
	err := r.coll.InsertReturning(&u)
	if err != nil {
		return domain.User{}, err
	}
	return r.mapModelToDomain(u), nil
}

func (r userRepository) FindById(id uint64) (domain.User, error) {
	var u user
	err := r.coll.Find(db.Cond{"id": id}).One(&u)
	if err != nil {
		return domain.User{}, err
	}
	return r.mapModelToDomain(u), nil
}

func (r userRepository) Update(user domain.User) (domain.User, error) {
	u := r.mapDomainToModel(user)
	u.UpdatedDate = time.Now()
	err := r.coll.Find(db.Cond{"id": u.Id}).Update(&u)
	if err != nil {
		return domain.User{}, err
	}
	return r.mapModelToDomain(u), nil
}

func (r userRepository) ShowList() ([]domain.User, error) {
	var users []user

	err := r.coll.Find(db.Cond{"deleted_date": nil}).All(&users)
	if err != nil {
		return []domain.User{}, err
	}

	return r.mapModelToDomainCollection(users), nil
}
func (r userRepository) Delete(id uint64) error {
	return r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).Update(map[string]interface{}{"deleted_date": time.Now()})
}

func (r userRepository) mapDomainToModel(d domain.User) user {
	return user{
		Id:          d.Id,
		Name:        d.Name,
		Email:       d.Email,
		Password:    d.Password,
		CreatedDate: d.CreatedDate,
		UpdatedDate: d.UpdatedDate,
		DeletedDate: d.DeletedDate,
	}
}

func (r userRepository) mapModelToDomain(m user) domain.User {
	return domain.User{
		Id:          m.Id,
		Name:        m.Name,
		Email:       m.Email,
		Password:    m.Password,
		CreatedDate: m.CreatedDate,
		UpdatedDate: m.UpdatedDate,
		DeletedDate: m.DeletedDate,
	}
}

func (r userRepository) mapModelToDomainCollection(m []user) []domain.User {
	result := make([]domain.User, len(m))

	for i := range m {
		result[i] = r.mapModelToDomain(m[i])
	}

	return result
}
