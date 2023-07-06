package dns

import "testing"

func TestQueryCNAME(t *testing.T) {

	r, err := QueryCNAME("autodiscover.elmasy.com")
	if err != nil {
		t.Fatalf("TestQueryCNAME failed: %s\n", err)
	}

	t.Logf("%#v\n", r)
}

func TestQueryCNAMERetry(t *testing.T) {

	r, err := QueryCNAMERetry("elmasy.com", 3)
	if err != nil {
		t.Fatalf("TestQueryCNAMERetry failed: %s\n", err)
	}

	t.Logf("%#v\n", r)
}
