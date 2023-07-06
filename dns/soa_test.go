package domain

import "testing"

func TestQuerySOA(t *testing.T) {

	r, err := QuerySOA("elmasy.com")
	if err != nil {
		t.Fatalf("TestQuerySOA failed: %s\n", err)
	}

	t.Logf("%#v\n", r)
}
