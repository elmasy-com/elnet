package dns

import "testing"

func TestQueryTXT(t *testing.T) {

	r, err := QueryTXT("elmasy.com")
	if err != nil {
		t.Fatalf("TestQueryTXT failed: %s\n", err)
	}

	for i := range r {
		t.Logf("elmasy.com TXT -> %s\n", r[i])
	}
}

func TestQueryTXTRetry(t *testing.T) {

	r, err := QueryTXTRetry("elmasy.com")
	if err != nil {
		t.Fatalf("TestQueryTXTRetry failed: %s\n", err)
	}

	for i := range r {
		t.Logf("elmasy.com TXT -> %s\n", r[i])
	}
}
