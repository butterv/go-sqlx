package persistence

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/butterv/go-sqlx/app/domain/model"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FindByID(id model.UserID) (*model.User, error) {
	var u model.User

	err := r.db.Get(&u, "SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &u, nil
}

func (r *userRepository) FindByIDs(ids []model.UserID) ([]model.User, error) {
	var us []model.User

	query, params, err := sqlx.In("SELECT * FROM users WHERE id IN (?)", ids)
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

func (r *userRepository) Create(id model.UserID, email string) error {
	u := &model.User{
		ID:    id,
		Email: email,
	}

	_, err := r.db.NamedExec("INSERT INTO users (id, email) VALUES (:id, :email)", u)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) DeleteByID(id model.UserID) error {
	_, err := r.db.Exec("UPDATE users SET deleted_at = NOW() WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
