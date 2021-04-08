package repository

import "github.com/butterv/go-sqlx/app/domain/model"

type UserRepositoryAccess interface {
	FindByID(id model.UserID) (*model.User, error)
	FindByIDs(ids []model.UserID) ([]model.User, error)
}

type UserRepositoryModify interface {
	UserRepositoryAccess

	Create(id model.UserID, email string) error
	DeleteByID(id model.UserID) error
}
