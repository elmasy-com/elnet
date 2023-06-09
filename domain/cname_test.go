package domain

import "testing"

func TestQueryCNAME(t *testing.T) {

	r, err := QueryCNAME("autodiscover.elmasy.com")
	if err != nil {
		t.Fatalf("TestQueryCNAME failed: %s\n", err)
	}

	t.Logf("%#v\n", r)
}
