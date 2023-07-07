package dns

import "testing"

func TestQueryAAAA(t *testing.T) {

	r, err := QueryAAAA("elmasy.com")
	if err != nil {
		t.Fatalf("TestQueryAAAA failed: %s\n", err)
	}

	for i := range r {
		t.Logf("elmasy.com AAAA -> %s\n", r[i])
	}
}

func TestQueryAAAARetry(t *testing.T) {

	r, err := QueryAAAARetry("elmasy.com")
	if err != nil {
		t.Fatalf("TestQueryAAAARetry failed: %s\n", err)
	}

	for i := range r {
		t.Logf("elmasy.com AAAA -> %s\n", r[i])
	}
}

func TestQueryAAAARetryStr(t *testing.T) {

	r, err := QueryAAAARetryStr("elmasy.com")
	if err != nil {
		t.Fatalf("TestQueryAAAARetry failed: %s\n", err)
	}

	for i := range r {
		t.Logf("elmasy.com AAAA -> %s\n", r[i])
	}
}

func TestIsSetAAAA(t *testing.T) {

	r, err := IsSetAAAA("elmasy.com")
	if err != nil {
		t.Fatalf("TestAIsSet failed: %s\n", err)
	}

	if r != true {
		t.Fatalf("TestAIsSet failed: elmasy.com is not set!\n")
	}
}
