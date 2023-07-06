package dns

import "testing"

func TestQueryTXT(t *testing.T) {

	r, err := QueryTXT("elmasy.com")
	if err != nil {
		t.Fatalf("TestQueryTXT failed: %s\n", err)
	}

	t.Logf("%#v\n", r)
}

func TestQueryTXTRetry(t *testing.T) {

	r, err := QueryTXTRetry("elmasy.com", 3)
	if err != nil {
		t.Fatalf("TestQueryTXTRetry failed: %s\n", err)
	}

	t.Logf("%#v\n", r)
}
