package user_test

import (
	"errors"
	"regexp"
	"testing"

	"github.com/jmoiron/sqlx"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/butterv/go-sqlx/app/domain/model"
	"github.com/butterv/go-sqlx/app/infrastructure/persistence/test"
	"github.com/butterv/go-sqlx/app/infrastructure/persistence/user"
)

func TestNewUserRepositoryModify(t *testing.T) {
	_, tx := test.TxMock(t)

	got := user.NewUserRepositoryModify(tx)
	if got == nil {
		t.Fatalf("user.NewUserRepositoryModify(tx) = nil; want not nil")
	}
}

func TestNewUserRepositoryModify_ReturnsNil(t *testing.T) {
	var tx *sqlx.Tx

	got := user.NewUserRepositoryModify(tx)
	if got != nil {
		t.Fatalf("user.NewUserRepositoryModify(tx) != nil; want nil")
	}
}

func TestTxUserRepository_Create(t *testing.T) {
	wantQuery := "INSERT INTO users (id, email) VALUES (?, ?)"

	mock, tx := test.TxMock(t)

	id := model.UserID("TEST_ID")
	email := "TEST_EMAIL"

	mock.ExpectExec(regexp.QuoteMeta(wantQuery)).
		WithArgs(id, email).
		WillReturnResult(sqlmock.NewResult(1, 1))

	r := user.NewUserRepositoryModify(tx)

	err := r.Create(id, email)
	if err != nil {
		t.Fatalf("r.Create(%s, %s) = %#v; want nil", id, email, err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("mock.ExpectationsWereMet() = %#v; want nil", err)
	}
}

func TestTxUserRepository_Create_Error(t *testing.T) {
	wantErr := errors.New("an error occurred")
	wantQuery := "INSERT INTO users (id, email) VALUES (?, ?)"

	mock, tx := test.TxMock(t)

	id := model.UserID("TEST_ID")
	email := "TEST_EMAIL"

	mock.ExpectExec(regexp.QuoteMeta(wantQuery)).
		WithArgs(id, email).
		WillReturnError(wantErr)

	r := user.NewUserRepositoryModify(tx)

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

func TestTxUserRepository_DeleteByID(t *testing.T) {
	wantQuery := "UPDATE users SET deleted_at = NOW() WHERE id = ?"

	mock, tx := test.TxMock(t)

	id := model.UserID("TEST_ID")

	mock.ExpectExec(regexp.QuoteMeta(wantQuery)).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	r := user.NewUserRepositoryModify(tx)

	err := r.DeleteByID(id)
	if err != nil {
		t.Fatalf("r.DeleteByID(%s) = %#v; want nil", id, err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("mock.ExpectationsWereMet() = %#v; want nil", err)
	}
}

func TestTxUserRepository_DeleteByID_Error(t *testing.T) {
	wantErr := errors.New("an error occurred")
	wantQuery := "UPDATE users SET deleted_at = NOW() WHERE id = ?"

	mock, tx := test.TxMock(t)

	id := model.UserID("TEST_ID")

	mock.ExpectExec(regexp.QuoteMeta(wantQuery)).
		WithArgs(id).
		WillReturnError(wantErr)

	r := user.NewUserRepositoryModify(tx)

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
