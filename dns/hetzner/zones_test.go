package hetzner

import (
	"errors"
	"os"
	"testing"
	"time"
)

func TestGetAllZOnes(t *testing.T) {

	c := NewClientWithTimeout(os.Getenv("HETZNER_DNS_KEY"), 5*time.Second)

	zs, err := c.GetAllZones()
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	t.Logf("%#v\n", zs)
}

func TestGetAllZOnesNoAPIKey(t *testing.T) {

	c := NewClientWithTimeout("", 5*time.Second)

	zs, err := c.GetAllZones()

	if err == nil {
		t.Fatalf("FAIL: error is nil, returned: %#v\n", zs)
	}

	if !errors.Is(err, ErrNoAPIKey) {
		t.Fatalf("FAIL: error got: %s, want: %s\n", err, ErrNoAPIKey)
	}
}

func TestGetAllZOnesInvalidAPIKey(t *testing.T) {

	c := NewClientWithTimeout("invalid", 5*time.Second)

	zs, err := c.GetAllZones()

	if err == nil {
		t.Fatalf("FAIL: error is nil, returned: %#v\n", zs)
	}

	if !errors.Is(err, ErrInvalidAPIKey) {
		t.Fatalf("FAIL: error got: %s, want: %s\n", err, ErrInvalidAPIKey)
	}
}
