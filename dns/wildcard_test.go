package dns

import "testing"

func TestIsWildcard(t *testing.T) {

	r, err := IsWildcard("www.example.com", TypeA)
	if err != nil {
		t.Fatalf("Fail: failed to check if www.example.com is a wildcard: %s\n", err)
	}

	if r {
		t.Fatalf("Fail: www.example.com is reported a wildcard domain\n")
	}

	r, err = IsWildcard("test.cyberdivine.ch", TypeA)
	if err != nil {
		t.Fatalf("Fail: failed to check if test.cyberdivine.ch is a wildcard: %s\n", err)
	}

	if !r {
		t.Fatalf("Fail: test.cyberdivine.ch is not reported a wildcard domain\n")
	}
}
