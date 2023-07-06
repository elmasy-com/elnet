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

	for i := range r {
		t.Logf("elmasy.com A -> %s\n", r[i])
	}
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

func TestQueryARetry(t *testing.T) {

	r, err := QueryARetry("elmasy.com")
	if err != nil {
		t.Fatalf("TestQueryARetry failed: %s\n", err)
	}

	for i := range r {
		t.Logf("elmasy.com A -> %s\n", r[i])
	}
}

func TestQueryARetryInvalidMaxRetries(t *testing.T) {

	MaxRetries = 0

	_, err := QueryARetry("elmasy.com")
	if err == nil {
		t.Fatalf("TestQueryARetry failed: err is nil\n")
	}

	if !errors.Is(err, ErrInvalidMaxRetries) {
		t.Fatalf("TestQueryARetry failed: %s\n", err)
	}

	MaxRetries = 5
}

func TestIsSetA(t *testing.T) {

	r, err := IsSetA("elmasy.com")
	if err != nil {
		t.Fatalf("TestIsSetA failed: %s\n", err)
	}

	if r != true {
		t.Fatalf("TestIsSetA failed: elmasy.com is not set!\n")
	}
}
