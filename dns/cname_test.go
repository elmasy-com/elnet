package dns

import "testing"

func TestQueryCNAME(t *testing.T) {

	r, err := QueryCNAME("autodiscover.elmasy.com")
	if err != nil {
		t.Fatalf("TestQueryCNAME failed: %s\n", err)
	}

	for i := range r {
		t.Logf("autodiscover.elmasy.com CNAME -> %s\n", r[i])
	}
}

func TestQueryCNAMERetry(t *testing.T) {

	r, err := QueryCNAMERetry("autodiscover.elmasy.com")
	if err != nil {
		t.Fatalf("TestQueryCNAMERetry failed: %s\n", err)
	}

	for i := range r {
		t.Logf("autodiscover.elmasy.com CNAME -> %s\n", r[i])
	}
}
