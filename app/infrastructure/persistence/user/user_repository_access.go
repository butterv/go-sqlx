package user

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/butterv/go-sqlx/app/domain/model"
)

type userRepositoryAccess struct {
	db *sqlx.DB
}

func NewUserRepositoryAccess(db *sqlx.DB) *userRepositoryAccess {
	if db == nil {
		return nil
	}

	return &userRepositoryAccess{
		db: db,
	}
}

func (r *userRepositoryAccess) FindByID(id model.UserID) (*model.User, error) {
	var u model.User

	err := r.db.Get(&u, "SELECT * FROM users WHERE id = ? AND deleted_at IS NULL", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &u, nil
}

func (r *userRepositoryAccess) FindByIDs(ids []model.UserID) (model.Users, error) {
	var us model.Users

	query, params, err := sqlx.In("SELECT * FROM users WHERE id IN (?) AND deleted_at IS NULL", ids)
	if err != nil {
		return nil, err
	}

	err = r.db.Select(&us, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return us, nil
}
