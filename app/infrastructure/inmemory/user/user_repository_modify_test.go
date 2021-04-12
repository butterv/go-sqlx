package user_test

import (
	"testing"
	"time"

	"github.com/butterv/go-sqlx/app/domain/model"
	"github.com/butterv/go-sqlx/app/infrastructure/inmemory/test"
	"github.com/butterv/go-sqlx/app/infrastructure/inmemory/user"
)

func TestNewUserRepositoryModify(t *testing.T) {
	s := test.NewStore()

	got := user.NewUserRepositoryModify(s)
	if got == nil {
		t.Fatalf("user.NewUserRepositoryModify(s) = nil; want not nil")
	}
}

func TestNewUserRepositoryModify_ReturnsNil(t *testing.T) {
	var s *test.Store

	got := user.NewUserRepositoryModify(s)
	if got != nil {
		t.Fatalf("user.NewUserRepositoryModify(s) != nil; want nil")
	}
}

func TestUserRepositoryModify_Create(t *testing.T) {
	s := test.NewStore()
	r := user.NewUserRepositoryModify(s)

	id := model.UserID("TEST_ID")
	email := "TEST_EMAIL"

	err := r.Create(id, email)
	if err != nil {
		t.Fatalf("r.Create(%s, %s) = _, %#v; want nil", id, email, err)
	}
	if u := s.Users.FindByID(id); u == nil {
		t.Errorf("s.Users.FindByID(%s) = nil; want not nil", id)
	}
}

func TestUserRepositoryModify_DeleteByID(t *testing.T) {
	now := time.Now()
	s := test.NewStore()
	s.AddUsers(&model.User{
		ID:        "TEST_ID",
		Email:     "TEST_EMAIL",
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: nil,
	})

	r := user.NewUserRepositoryModify(s)

	id := model.UserID("TEST_ID")

	err := r.DeleteByID(id)
	if err != nil {
		t.Fatalf("r.DeleteByID(%s) = _, %#v; want nil", id, err)
	}

	u := s.Users.FindByID(id)
	if u == nil {
		t.Fatalf("s.Users.FindByID(%s) = nil; want not nil", id)
	}
	if u.DeletedAt == nil {
		t.Errorf("u.DeletedAt is nil; want not nil")
	}
}
