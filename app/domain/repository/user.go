package repository

import (
	"github.com/butterv/go-sqlx/app/domain/model"
)

type UserRepositoryAccess interface {
	FindByID(id model.UserID) (*model.User, error)
	FindByIDs(ids []model.UserID) (model.Users, error)
}

type UserRepositoryModify interface {
	Create(id model.UserID, email string) error
	DeleteByID(id model.UserID) error
}
