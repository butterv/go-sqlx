package user

import (
	"time"

	"github.com/butterv/go-sqlx/app/domain/model"
	"github.com/butterv/go-sqlx/app/infrastructure/inmemory/test"
)

type userRepositoryModify struct {
	s *test.Store
}

func NewUserRepositoryModify(s *test.Store) *userRepositoryModify {
	if s == nil {
		return nil
	}

	return &userRepositoryModify{
		s: s,
	}
}

func (r *userRepositoryModify) Create(id model.UserID, email string) error {
	now := time.Now()
	u := &model.User{
		ID:        id,
		Email:     email,
		CreatedAt: now,
		UpdatedAt: now,
	}

	r.s.Users = append(r.s.Users, u)

	return nil
}

func (r *userRepositoryModify) DeleteByID(id model.UserID) error {
	for _, u := range r.s.Users {
		if u.ID == id {
			now := time.Now()
			u.DeletedAt = &now
		}
	}

	return nil
}
