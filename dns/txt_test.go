package domain

import "testing"

func TestQueryTXT(t *testing.T) {

	r, err := QueryTXT("elmasy.com")
	if err != nil {
		t.Fatalf("TestQueryTXT failed: %s\n", err)
	}

	t.Logf("%#v\n", r)
}
