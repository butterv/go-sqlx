package user

import (
	"github.com/butterv/go-sqlx/app/domain/model"
	"github.com/butterv/go-sqlx/app/infrastructure/inmemory/test"
)

type userRepositoryAccess struct {
	s *test.Store
}

func NewUserRepositoryAccess(s *test.Store) *userRepositoryAccess {
	if s == nil {
		return nil
	}

	return &userRepositoryAccess{
		s: s,
	}
}

func (r *userRepositoryAccess) FindByID(id model.UserID) (*model.User, error) {
	for _, u := range r.s.Users {
		if u.ID == id && u.DeletedAt == nil {
			return u, nil
		}
	}

	return nil, nil
}

func (r *userRepositoryAccess) FindByIDs(ids []model.UserID) (model.Users, error) {
	var us model.Users

	for _, u := range r.s.Users {
		for _, id := range ids {
			if u.ID == id && u.DeletedAt == nil {
				us = append(us, u)
			}
		}
	}

	return us, nil
}
