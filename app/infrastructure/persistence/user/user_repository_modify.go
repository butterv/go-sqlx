package user

import (
	"github.com/jmoiron/sqlx"

	"github.com/butterv/go-sqlx/app/domain/model"
)

type userRepositoryModify struct {
	tx *sqlx.Tx
}

func NewUserRepositoryModify(tx *sqlx.Tx) *userRepositoryModify {
	if tx == nil {
		return nil
	}

	return &userRepositoryModify{
		tx: tx,
	}
}

func (r *userRepositoryModify) Create(id model.UserID, email string) error {
	u := &model.User{
		ID:    id,
		Email: email,
	}

	_, err := r.tx.NamedExec("INSERT INTO users (id, email) VALUES (:id, :email)", u)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepositoryModify) DeleteByID(id model.UserID) error {
	_, err := r.tx.Exec("UPDATE users SET deleted_at = NOW() WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
