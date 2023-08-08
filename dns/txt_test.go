package dns

import "testing"

func TestQueryTXT(t *testing.T) {

	r, err := QueryTXT("elmasy.com")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	for i := range r {
		t.Logf("elmasy.com TXT -> %s\n", r[i])
	}
}

func TestTryQueryTXT(t *testing.T) {

	r, err := TryQueryTXT("elmasy.com")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	for i := range r {
		t.Logf("elmasy.com TXT -> %s\n", r[i])
	}
}

func TestIsSetTXT(t *testing.T) {

	r, err := IsSetTXT("elmasy.com")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	if !r {
		t.Fatalf("FAIL: TXT is not set for elmasy.com\n")
	}
}
