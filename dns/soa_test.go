package dns

import "testing"

func TestQuerySOA(t *testing.T) {

	r, err := QuerySOA("elmasy.com")
	if err != nil {
		t.Fatalf("TestQuerySOA failed: %s\n", err)
	}

	t.Logf("%#v\n", r)
}

func TestQuerySOARetry(t *testing.T) {

	r, err := QuerySOARetry("elmasy.com", 3)
	if err != nil {
		t.Fatalf("TestQuerySOARetry failed: %s\n", err)
	}

	t.Logf("%#v\n", r)
}
