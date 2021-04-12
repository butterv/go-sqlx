package inmemory

import (
	"context"

	"github.com/butterv/go-sqlx/app/domain/repository"
	"github.com/butterv/go-sqlx/app/infrastructure/inmemory/test"
	"github.com/butterv/go-sqlx/app/infrastructure/inmemory/user"
)

type inmemoryRepository struct {
	s *test.Store
}

type inmemoryConnection struct {
	s *test.Store
}

type inmemoryTransaction struct {
	s *test.Store
}

func New(s *test.Store) repository.Repository {
	return &inmemoryRepository{
		s: s,
	}
}

func (r *inmemoryRepository) NewConnection(ctx context.Context) repository.Connection {
	return &inmemoryConnection{
		s: r.s,
	}
}

func (c *inmemoryConnection) Close() error {
	return nil
}

func (c *inmemoryConnection) RunTransaction(f func(repository.Transaction) error) error {
	err := f(&inmemoryTransaction{
		s: c.s,
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *inmemoryConnection) User() repository.UserRepositoryAccess {
	return user.NewUserRepositoryAccess(c.s)
}

func (c *inmemoryTransaction) User() repository.UserRepositoryModify {
	return user.NewUserRepositoryModify(c.s)
}
