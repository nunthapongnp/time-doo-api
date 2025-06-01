package user

import (
	"time-doo-api/internal/domain"
	ctx "time-doo-api/pkg/context"
)

type repository struct {
	db *ctx.AppDbContext
}

func NewUserRepository(db *ctx.AppDbContext) UserRepository {
	return &repository{db: db}
}

func (r *repository) Add(user *domain.User) error {
	return r.db.Create(user)
}

func (r *repository) FindByID(id int64) (*domain.User, error) {
	var u domain.User
	err := r.db.First(&u, id)
	return &u, err
}

func (r *repository) FindByEmail(email string) (*domain.User, error) {
	var u domain.User
	err := r.db.First(&u, "email = ?", email)
	return &u, err
}

func (r *repository) GetAll() ([]*domain.User, error) {
	var users []*domain.User
	err := r.db.Find(&users)
	return users, err
}

func (r *repository) Edit(user *domain.User) error {
	return r.db.Update(user, user)
}

func (r *repository) Remove(id int64) error {
	return r.db.Delete(&domain.User{ID: id})
}
