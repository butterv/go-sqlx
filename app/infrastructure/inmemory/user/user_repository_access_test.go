package user_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/butterv/go-sqlx/app/domain/model"
	"github.com/butterv/go-sqlx/app/infrastructure/inmemory/test"
	"github.com/butterv/go-sqlx/app/infrastructure/inmemory/user"
)

func TestNewUserRepositoryAccess(t *testing.T) {
	s := test.NewStore()

	got := user.NewUserRepositoryAccess(s)
	if got == nil {
		t.Fatalf("user.NewUserRepositoryAccess(s) = nil; want not nil")
	}
}

func TestNewUserRepositoryAccess_ReturnsNil(t *testing.T) {
	var s *test.Store

	got := user.NewUserRepositoryAccess(s)
	if got != nil {
		t.Fatalf("user.NewUserRepositoryAccess(s) != nil; want nil")
	}
}

func TestUserRepository_FindByID(t *testing.T) {
	now := time.Now()
	want := &model.User{
		ID:        "TEST_ID",
		Email:     "TEST_EMAIL",
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: nil,
	}

	s := test.NewStore()
	s.AddUsers(want)
	r := user.NewUserRepositoryAccess(s)

	id := model.UserID("TEST_ID")
	got, err := r.FindByID(id)
	if err != nil {
		t.Fatalf("r.FindByID(%s) = _, %#v; want nil", id, err)
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("r.FindByID(%s) = %#v, _; want %v\ndiff = %s", id, got, want, diff)
	}
}

func TestUserRepository_FindByID_NotFound(t *testing.T) {
	s := test.NewStore()
	r := user.NewUserRepositoryAccess(s)

	id := model.UserID("TEST_ID")
	got, err := r.FindByID(id)
	if err != nil {
		t.Fatalf("r.FindByID(%s) = _, %#v; want nil", id, err)
	}
	if got != nil {
		t.Errorf("r.FindByID(%s) = %#v, _; want nil", id, got)
	}
}

func TestUserRepository_FindByIDs(t *testing.T) {
	now := time.Now()
	want := model.Users{
		{
			ID:        "TEST_ID1",
			Email:     "TEST_EMAIL1",
			CreatedAt: now,
			UpdatedAt: now,
			DeletedAt: nil,
		},
		{
			ID:        "TEST_ID2",
			Email:     "TEST_EMAIL2",
			CreatedAt: now,
			UpdatedAt: now,
			DeletedAt: nil,
		},
		{
			ID:        "TEST_ID3",
			Email:     "TEST_EMAIL3",
			CreatedAt: now,
			UpdatedAt: now,
			DeletedAt: nil,
		},
	}

	s := test.NewStore()
	s.AddUsers(want...)
	r := user.NewUserRepositoryAccess(s)

	ids := []model.UserID{"TEST_ID1", "TEST_ID2", "TEST_ID3"}
	got, err := r.FindByIDs(ids)
	if err != nil {
		t.Fatalf("r.FindByIDs(%v) = _, %#v; want nil", ids, err)
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("r.FindByIDs(%v) = %#v, _; want %v\ndiff = %s", ids, got, want, diff)
	}
}

func TestUserRepository_FindByIDs_NotFound(t *testing.T) {
	s := test.NewStore()
	r := user.NewUserRepositoryAccess(s)

	ids := []model.UserID{"TEST_ID1", "TEST_ID2", "TEST_ID3"}
	got, err := r.FindByIDs(ids)
	if err != nil {
		t.Fatalf("r.FindByIDs(%v) = _, %#v; want nil", ids, err)
	}
	if got != nil {
		t.Errorf("r.FindByIDs(%v) = %#v, _; want nil", ids, got)
	}
}
