package test_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/butterv/go-sqlx/app/domain/model"
	"github.com/butterv/go-sqlx/app/infrastructure/inmemory/test"
)

func TestNewStore(t *testing.T) {
	want := &test.Store{}

	got := test.NewStore()
	if diff := cmp.Diff(got, want, cmp.AllowUnexported(test.Store{})); diff != "" {
		t.Errorf("NewStore() = %#v; want %v\ndiff = %s", got, want, diff)
	}
}

func TestStore_AddUsers(t *testing.T) {
	now := time.Now()
	want := &test.Store{
		Users: model.Users{
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
		},
	}

	s := test.NewStore()
	s.AddUsers(want.Users...)
	if diff := cmp.Diff(s, want, cmp.AllowUnexported(test.Store{})); diff != "" {
		t.Errorf("NewStore() = %#v; want %v\ndiff = %s", s, want, diff)
	}
}
