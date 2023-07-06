package dns

import "testing"

func TestQueryMX(t *testing.T) {

	r, err := QueryMX("elmasy.com")
	if err != nil {
		t.Fatalf("TestQueryMX failed: %s\n", err)
	}

	for i := range r {
		t.Logf("elmasy.com MX -> %s\n", r[i])
	}
}

func TestQueryMXRetry(t *testing.T) {

	r, err := QueryMXRetry("elmasy.com")
	if err != nil {
		t.Fatalf("TestQueryMXRetry failed: %s\n", err)
	}

	for i := range r {
		t.Logf("elmasy.com MX -> %s\n", r[i])
	}
}
