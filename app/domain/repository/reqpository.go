package repository

import "context"

type Repository interface {
	NewConnection(ctx context.Context) (Connection, error)
}

type Connection interface {
	Close() error
	RunTransaction(f func(Transaction) error) error

	User() UserRepositoryAccess
}

type Transaction interface {
	User() UserRepositoryModify
}
