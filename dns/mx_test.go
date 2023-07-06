package dns

import "testing"

func TestQueryMX(t *testing.T) {

	r, err := QueryMX("elmasy.com")
	if err != nil {
		t.Fatalf("TestQueryMX failed: %s\n", err)
	}

	t.Logf("%#v\n", r)
}

func TestQueryMXRetry(t *testing.T) {

	r, err := QueryMXRetry("elmasy.com", 3)
	if err != nil {
		t.Fatalf("TestQueryMXRetry failed: %s\n", err)
	}

	t.Logf("%#v\n", r)
}
