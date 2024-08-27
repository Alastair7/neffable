package db

import (
	"errors"
	"testing"
)

type dbProviderMock struct {
	Success bool
}

func (dbMock *dbProviderMock) InitDB() error {
	if dbMock.Success == false {
		return errors.New("Error calling ping")
	}
	return nil
}

func (dbMock *dbProviderMock) SetSuccess(value bool) {
	dbMock.Success = value
} 
func TestInitDB(t *testing.T) {
	t.Run("not calls ping when error occurs", func(t *testing.T) {
		mock := &dbProviderMock{}
		mock.SetSuccess(false)
		err := InitDB(mock)
		if err == nil {
			t.Error("want error but instead got no errors")
		}
	t.Run("calls ping when not error occurs", func(t *testing.T) {
		mock := &dbProviderMock{}
		mock.SetSuccess(true)

		err := InitDB(mock)

		if err != nil {
			t.Error("want success but instead got error")
		}
	})
	})
}