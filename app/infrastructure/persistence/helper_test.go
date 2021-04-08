package persistence_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func dbMock(t *testing.T) (sqlmock.Sqlmock, *sqlx.DB) {
	t.Helper()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New() = _, _, %#v; want nil", err)
	}

	sqlxDB := sqlx.NewDb(db, "mysql")

	return mock, sqlxDB
}
