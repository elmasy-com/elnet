package dns

import "testing"

func TestQueryMX(t *testing.T) {

	r, err := QueryMX("elmasy.com")
	if err != nil {
		t.Fatalf("TestQueryMX failed: %s\n", err)
	}

	t.Logf("%#v\n", r)
}
