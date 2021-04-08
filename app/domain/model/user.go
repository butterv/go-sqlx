package model

import "time"

type UserID string

type User struct {
	ID    UserID `db:"id"`
	Email string `db:"email"`

	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
