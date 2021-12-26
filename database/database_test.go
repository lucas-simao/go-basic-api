package database

import (
	"testing"

	"github.com/jmoiron/sqlx"
)

func TestConnection(t *testing.T) {
	var got *sqlx.DB

	want := Connect()

	if got == want {
		t.Error("error")
	}

}

func TestReadEnvs(t *testing.T) {
	got := readEnvs()

	notWant := DatabaseConfig{}

	if got != notWant {
		return
	} else {
		t.Errorf("got %q, wanted %q", got, notWant)
	}

}
