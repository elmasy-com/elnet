package dns

import (
	"errors"
	"testing"
)

func TestQueryA(t *testing.T) {

	r, err := QueryA("elmasy.com")
	if err != nil {
		t.Fatalf("TestQueryA failed: %s\n", err)
	}

	t.Logf("%#v\n", r)
}

func TestQueryAInvalid(t *testing.T) {

	r, err := QueryA("invalid.elmasy.com")
	if err == nil {
		t.Fatalf("TestQueryA failed: error must be NXDOMAIN\n")
	}
	if !errors.Is(err, ErrName) {
		t.Fatalf("Failed: invalid error: %s\n", err)
	}

	if len(r) != 0 {
		t.Fatalf("Failed: invalid result length: %d\n", len(r))
	}
}

func TestQueryALenZero(t *testing.T) {

	r, err := QueryA("_dmarc.elmasy.com")
	if err != nil {
		t.Fatalf("Failed: %s\n", err)
	}

	if len(r) != 0 {
		t.Fatalf("Failed: invalid result length: %d\n", len(r))
	}
}

func TestIsSetA(t *testing.T) {

	r, err := IsSetA("elmasy.com")
	if err != nil {
		t.Fatalf("TestIsSetA failed: %s\n", err)
	}

	if r != true {
		t.Fatalf("TestIsSetA failed: elmasy.com is not set!\n")
	}

	t.Logf("%#v\n", r)
}
