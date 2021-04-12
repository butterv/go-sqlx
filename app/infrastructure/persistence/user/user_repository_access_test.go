package user_test

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"

	"github.com/butterv/go-sqlx/app/domain/model"
	"github.com/butterv/go-sqlx/app/infrastructure/persistence/test"
	"github.com/butterv/go-sqlx/app/infrastructure/persistence/user"
)

func TestNewUserRepositoryAccess(t *testing.T) {
	_, db := test.DbMock(t)
	defer db.Close()

	got := user.NewUserRepositoryAccess(db)
	if got == nil {
		t.Fatalf("user.NewUserRepositoryAccess(db) = nil; want not nil")
	}
}

func TestNewUserRepositoryAccess_ReturnsNil(t *testing.T) {
	var db *sqlx.DB

	got := user.NewUserRepositoryAccess(db)
	if got != nil {
		t.Fatalf("user.NewUserRepositoryAccess(db) != nil; want nil")
	}
}

func TestDbUserRepository_FindByID(t *testing.T) {
	now := time.Now()
	want := &model.User{
		ID:        "TEST_ID",
		Email:     "TEST_EMAIL",
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: nil,
	}

	wantQuery := "SELECT * FROM users WHERE id = ? AND deleted_at IS NULL"

	mock, db := test.DbMock(t)
	defer db.Close()

	id := model.UserID("TEST_ID")

	mock.ExpectQuery(regexp.QuoteMeta(wantQuery)).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "created_at", "updated_at", "deleted_at"}).
			AddRow(want.ID, want.Email, want.CreatedAt, want.UpdatedAt, want.DeletedAt))

	r := user.NewUserRepositoryAccess(db)

	got, err := r.FindByID(id)
	if err != nil {
		t.Fatalf("r.FindByID(%s) = _, %#v; want nil", id, err)
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("r.FindByID(%s) = %#v, _; want %v\ndiff = %s", id, got, want, diff)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("mock.ExpectationsWereMet() = %#v; want nil", err)
	}
}

func TestDbUserRepository_FindByID_NotFound(t *testing.T) {
	wantQuery := "SELECT * FROM users WHERE id = ? AND deleted_at IS NULL"

	mock, db := test.DbMock(t)
	defer db.Close()

	id := model.UserID("TEST_ID")

	mock.ExpectQuery(regexp.QuoteMeta(wantQuery)).
		WithArgs(id).
		WillReturnError(sql.ErrNoRows)

	r := user.NewUserRepositoryAccess(db)

	got, err := r.FindByID(id)
	if err != nil {
		t.Fatalf("r.FindByID(%s) = _, %#v; want nil", id, err)
	}
	if got != nil {
		t.Errorf("r.FindByID(%s) = %#v, _; want nil", id, got)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("mock.ExpectationsWereMet() = %#v; want nil", err)
	}
}

func TestDbUserRepository_FindByID_Error(t *testing.T) {
	wantErr := errors.New("an error occurred")
	wantQuery := "SELECT * FROM users WHERE id = ? AND deleted_at IS NULL"

	mock, db := test.DbMock(t)
	defer db.Close()

	id := model.UserID("TEST_ID")

	mock.ExpectQuery(regexp.QuoteMeta(wantQuery)).
		WithArgs(id).
		WillReturnError(wantErr)

	r := user.NewUserRepositoryAccess(db)

	got, err := r.FindByID(id)
	if err == nil {
		t.Fatalf("r.FindByID(%s) = _, nil; want %v", id, wantErr)
	}
	if !errors.Is(err, wantErr) {
		t.Errorf("r.FindByID(%s) = _, %#v; want %v", id, err, wantErr)
	}
	if got != nil {
		t.Errorf("r.FindByID(%s) = %#v, _; want nil", id, got)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("mock.ExpectationsWereMet() = %#v; want nil", err)
	}
}

func TestDbUserRepository_FindByIDs(t *testing.T) {
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

	wantQuery := "SELECT * FROM users WHERE id IN (?, ?, ?) AND deleted_at IS NULL"

	mock, db := test.DbMock(t)
	defer db.Close()

	ids := []model.UserID{"TEST_ID1", "TEST_ID2", "TEST_ID3"}

	rows := sqlmock.NewRows([]string{"id", "email", "created_at", "updated_at", "deleted_at"})
	for _, v := range want {
		rows.AddRow(v.ID, v.Email, v.CreatedAt, v.UpdatedAt, v.DeletedAt)
	}

	mock.ExpectQuery(regexp.QuoteMeta(wantQuery)).
		WithArgs(ids[0], ids[1], ids[2]).
		WillReturnRows(rows)

	r := user.NewUserRepositoryAccess(db)

	got, err := r.FindByIDs(ids)
	if err != nil {
		t.Fatalf("r.FindByIDs(%v) = _, %#v; want nil", ids, err)
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("r.FindByIDs(%v) = %#v, _; want %v\ndiff = %s", ids, got, want, diff)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("mock.ExpectationsWereMet() = %#v; want nil", err)
	}
}

func TestDbUserRepository_FindByIDs_InReturnsError(t *testing.T) {
	_, db := test.DbMock(t)
	defer db.Close()

	var ids []model.UserID

	r := user.NewUserRepositoryAccess(db)

	got, err := r.FindByIDs(ids)
	if err == nil {
		t.Fatalf("r.FindByIDs(%v) = _, nil; want not nil", ids)
	}
	if got != nil {
		t.Errorf("r.FindByIDs(%v) = %#v, _; want nil", ids, got)
	}
}

func TestDbUserRepository_FindByIDs_NotFound(t *testing.T) {
	wantQuery := "SELECT * FROM users WHERE id IN (?, ?, ?) AND deleted_at IS NULL"

	mock, db := test.DbMock(t)
	defer db.Close()

	ids := []model.UserID{"TEST_ID1", "TEST_ID2", "TEST_ID3"}

	mock.ExpectQuery(regexp.QuoteMeta(wantQuery)).
		WithArgs(ids[0], ids[1], ids[2]).
		WillReturnError(sql.ErrNoRows)

	r := user.NewUserRepositoryAccess(db)

	got, err := r.FindByIDs(ids)
	if err != nil {
		t.Fatalf("r.FindByIDs(%v) = _, %#v; want nil", ids, err)
	}
	if got != nil {
		t.Errorf("r.FindByIDs(%v) = %#v, _; want nil", ids, got)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("mock.ExpectationsWereMet() = %#v; want nil", err)
	}
}

func TestDbUserRepository_FindByIDs_Error(t *testing.T) {
	wantErr := errors.New("an error occurred")
	wantQuery := "SELECT * FROM users WHERE id IN (?, ?, ?) AND deleted_at IS NULL"

	mock, db := test.DbMock(t)
	defer db.Close()

	ids := []model.UserID{"TEST_ID1", "TEST_ID2", "TEST_ID3"}

	mock.ExpectQuery(regexp.QuoteMeta(wantQuery)).
		WithArgs(ids[0], ids[1], ids[2]).
		WillReturnError(wantErr)

	r := user.NewUserRepositoryAccess(db)

	got, err := r.FindByIDs(ids)
	if err == nil {
		t.Fatalf("r.FindByIDs(%v) = _, nil; want %v", ids, wantErr)
	}
	if !errors.Is(err, wantErr) {
		t.Errorf("r.FindByIDs(%v) = _, %#v; want %v", ids, err, wantErr)
	}
	if got != nil {
		t.Errorf("r.FindByIDs(%v) = %#v, _; want nil", ids, got)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("mock.ExpectationsWereMet() = %#v; want nil", err)
	}
}
