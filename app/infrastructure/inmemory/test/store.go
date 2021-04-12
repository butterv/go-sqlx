package test

import (
	"github.com/butterv/go-sqlx/app/domain/model"
)

type Store struct {
	Users model.Users
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) AddUsers(us ...*model.User) {
	s.Users = append(s.Users, us...)
}
