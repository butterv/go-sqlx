package persistence_test

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"

	"github.com/butterv/go-sqlx/app/domain/model"
	"github.com/butterv/go-sqlx/app/infrastructure/persistence"
)

var usersColumns = []string{"id", "email", "created_at", "updated_at", "deleted_at"}

func TestUserRepository_FindByID(t *testing.T) {
	now := time.Now()
	want := &model.User{
		ID:        "TEST_ID",
		Email:     "TEST_EMAIL",
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: nil,
	}

	wantQuery := "SELECT * FROM users WHERE id = ?"

	mock, db := dbMock(t)
	defer db.Close()

	id := model.UserID("TEST_ID")

	mock.ExpectQuery(regexp.QuoteMeta(wantQuery)).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows(usersColumns).
			AddRow(want.ID, want.Email, want.CreatedAt, want.UpdatedAt, want.DeletedAt))

	r := persistence.NewUserRepository(db)

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

func TestUserRepository_FindByID_NotFound(t *testing.T) {
	wantQuery := "SELECT * FROM users WHERE id = ?"

	mock, db := dbMock(t)
	defer db.Close()

	id := model.UserID("TEST_ID")

	mock.ExpectQuery(regexp.QuoteMeta(wantQuery)).
		WithArgs(id).
		WillReturnError(sql.ErrNoRows)

	r := persistence.NewUserRepository(db)

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

func TestUserRepository_FindByID_Error(t *testing.T) {
	wantErr := errors.New("an error occurred")
	wantQuery := "SELECT * FROM users WHERE id = ?"

	mock, db := dbMock(t)
	defer db.Close()

	id := model.UserID("TEST_ID")

	mock.ExpectQuery(regexp.QuoteMeta(wantQuery)).
		WithArgs(id).
		WillReturnError(wantErr)

	r := persistence.NewUserRepository(db)

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

func TestUserRepository_FindByIDs(t *testing.T) {
	now := time.Now()
	want := []model.User{
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

	wantQuery := "SELECT * FROM users WHERE id IN (?, ?, ?)"

	mock, db := dbMock(t)
	defer db.Close()

	ids := []model.UserID{"TEST_ID1", "TEST_ID2", "TEST_ID3"}

	rows := sqlmock.NewRows(usersColumns)
	for _, v := range want {
		rows.AddRow(v.ID, v.Email, v.CreatedAt, v.UpdatedAt, v.DeletedAt)
	}

	mock.ExpectQuery(regexp.QuoteMeta(wantQuery)).
		WithArgs(ids[0], ids[1], ids[2]).
		WillReturnRows(rows)

	r := persistence.NewUserRepository(db)

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

func TestUserRepository_FindByIDs_InReturnsError(t *testing.T) {
	_, db := dbMock(t)
	defer db.Close()

	r := persistence.NewUserRepository(db)

	var ids []model.UserID

	got, err := r.FindByIDs(ids)
	if err == nil {
		t.Fatalf("r.FindByIDs(%v) = _, nil; want not nil", ids)
	}
	if got != nil {
		t.Errorf("r.FindByIDs(%v) = %#v, _; want nil", ids, got)
	}
}

func TestUserRepository_FindByIDs_NotFound(t *testing.T) {
	wantQuery := "SELECT * FROM users WHERE id IN (?, ?, ?)"

	mock, db := dbMock(t)
	defer db.Close()

	ids := []model.UserID{"TEST_ID1", "TEST_ID2", "TEST_ID3"}

	mock.ExpectQuery(regexp.QuoteMeta(wantQuery)).
		WithArgs(ids[0], ids[1], ids[2]).
		WillReturnError(sql.ErrNoRows)

	r := persistence.NewUserRepository(db)

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

func TestUserRepository_FindByIDs_Error(t *testing.T) {
	wantErr := errors.New("an error occurred")
	wantQuery := "SELECT * FROM users WHERE id IN (?, ?, ?)"

	mock, db := dbMock(t)
	defer db.Close()

	ids := []model.UserID{"TEST_ID1", "TEST_ID2", "TEST_ID3"}

	mock.ExpectQuery(regexp.QuoteMeta(wantQuery)).
		WithArgs(ids[0], ids[1], ids[2]).
		WillReturnError(wantErr)

	r := persistence.NewUserRepository(db)

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

func TestUserRepository_Create(t *testing.T) {
	wantQuery := "INSERT INTO users (id, email) VALUES (?, ?)"

	mock, db := dbMock(t)
	defer db.Close()

	id := model.UserID("TEST_ID")
	email := "TEST_EMAIL"

	mock.ExpectExec(regexp.QuoteMeta(wantQuery)).
		WithArgs(id, email).
		WillReturnResult(sqlmock.NewResult(1, 1))

	r := persistence.NewUserRepository(db)

	err := r.Create(id, email)
	if err != nil {
		t.Fatalf("r.Create(%s, %s) = %#v; want nil", id, email, err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("mock.ExpectationsWereMet() = %#v; want nil", err)
	}
}

func TestUserRepository_Create_Error(t *testing.T) {
	wantErr := errors.New("an error occurred")
	wantQuery := "INSERT INTO users (id, email) VALUES (?, ?)"

	mock, db := dbMock(t)
	defer db.Close()

	id := model.UserID("TEST_ID")
	email := "TEST_EMAIL"

	mock.ExpectExec(regexp.QuoteMeta(wantQuery)).
		WithArgs(id, email).
		WillReturnError(wantErr)

	r := persistence.NewUserRepository(db)

	err := r.Create(id, email)
	if err == nil {
		t.Fatalf("r.Create(%s, %s) = nil; want %v", id, email, wantErr)
	}
	if !errors.Is(err, wantErr) {
		t.Errorf("r.Create(%s, %s) = %#v; want %v", id, email, err, wantErr)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("mock.ExpectationsWereMet() = %#v; want nil", err)
	}
}

func TestUserRepository_DeleteByID(t *testing.T) {
	wantQuery := "UPDATE users SET deleted_at = NOW() WHERE id = ?"

	mock, db := dbMock(t)
	defer db.Close()

	id := model.UserID("TEST_ID")

	mock.ExpectExec(regexp.QuoteMeta(wantQuery)).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	r := persistence.NewUserRepository(db)

	err := r.DeleteByID(id)
	if err != nil {
		t.Fatalf("r.DeleteByID(%s) = %#v; want nil", id, err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("mock.ExpectationsWereMet() = %#v; want nil", err)
	}
}

func TestUserRepository_DeleteByID_Error(t *testing.T) {
	wantErr := errors.New("an error occurred")
	wantQuery := "UPDATE users SET deleted_at = NOW() WHERE id = ?"

	mock, db := dbMock(t)
	defer db.Close()

	id := model.UserID("TEST_ID")

	mock.ExpectExec(regexp.QuoteMeta(wantQuery)).
		WithArgs(id).
		WillReturnError(wantErr)

	r := persistence.NewUserRepository(db)

	err := r.DeleteByID(id)
	if err == nil {
		t.Fatalf("r.DeleteByID(%s) = nil; want %v", id, wantErr)
	}
	if !errors.Is(err, wantErr) {
		t.Errorf("r.DeleteByID(%s) = %#v; want %v", id, err, wantErr)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("mock.ExpectationsWereMet() = %#v; want nil", err)
	}
}
