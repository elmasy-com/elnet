package domain

import "testing"

func TestQueryAAAA(t *testing.T) {

	r, err := QueryAAAA("elmasy.com")
	if err != nil {
		t.Fatalf("TestQueryAAAA failed: %s\n", err)
	}

	t.Logf("%#v\n", r)
}

func TestIsSetAAAA(t *testing.T) {

	r, err := IsSetAAAA("elmasy.com")
	if err != nil {
		t.Fatalf("TestAIsSet failed: %s\n", err)
	}

	if r != true {
		t.Fatalf("TestAIsSet failed: elmasy.com is not set!\n")
	}

	t.Logf("%#v\n", r)
}
