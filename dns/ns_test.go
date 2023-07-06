package dns

import "testing"

func TestQueryNS(t *testing.T) {

	r, err := QueryNS("elmasy.com")
	if err != nil {
		t.Fatalf("TestQueryNS failed: %s\n", err)
	}

	t.Logf("%#v\n", r)
}

func TestQueryNSRetry(t *testing.T) {

	r, err := QueryNSRetry("elmasy.com")
	if err != nil {
		t.Fatalf("TestQueryNSRetry failed: %s\n", err)
	}

	t.Logf("%#v\n", r)
}
