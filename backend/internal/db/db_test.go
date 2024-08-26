package db

import (
	"testing"

	"github.com/pashagolub/pgxmock"
)

func TestInitDB(t *testing.T) {
	mock, err := pgxmock.NewPool(pgxmock.MonitorPingsOption(true))

	if err != nil {
		t.Fatalf("error creating the connection pool")
	}
	
	defer mock.Close()
	
	mock.ExpectPing()

	 if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	 }
}