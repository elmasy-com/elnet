package validator

import "testing"

func TestDomain(t *testing.T) {

	// Generate a string with 400 x 'a'
	longStringFunc := func() string {
		s := make([]byte, 400)
		for i := 0; i < 400; i++ {
			s = append(s, 'a')
		}
		return string(s)
	}

	cases := []struct {
		Domain string
		Result bool
	}{
		{Domain: "elmasy.com", Result: true},
		{Domain: "elmasy.com.", Result: true},
		{Domain: ".elmasy.com", Result: false},
		{Domain: ".elmasy.com.", Result: false},
		{Domain: "aaaaaa", Result: false},
		{Domain: longStringFunc(), Result: false},
		{Domain: "", Result: false},
		{Domain: ".", Result: false},
		{Domain: "a a", Result: false},
		{Domain: "a=a", Result: false},
		{Domain: " ", Result: false},
		{Domain: "*", Result: false},
	}

	for i := range cases {
		if r := Domain(cases[i].Domain); r != cases[i].Result {
			t.Fatalf("FAIL: %s is %v", cases[i].Domain, r)
		}
	}
}

func BenchmarkDomain(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Domain("test.elmasy.com.")
	}
}

func TestDomainPart(t *testing.T) {

	// Generate a string with 400 x 'a'
	longStringFunc := func() string {
		s := make([]byte, 64)
		for i := 0; i < 64; i++ {
			s = append(s, 'a')
		}
		return string(s)
	}

	cases := []struct {
		Domain string
		Result bool
	}{
		{Domain: "localhost", Result: true},
		{Domain: "elmasy.", Result: false},
		{Domain: ".elmasy", Result: false},
		{Domain: ".elmasy.com.", Result: false},
		{Domain: "aaaaaa", Result: true},
		{Domain: longStringFunc(), Result: false},
		{Domain: "", Result: false},
		{Domain: ".", Result: false},
		{Domain: "a a", Result: false},
		{Domain: "a=a", Result: false},
		{Domain: "a-a", Result: true},
		{Domain: "-aa", Result: false},
		{Domain: "aa-", Result: false},
		{Domain: "a--a", Result: false},
		{Domain: " ", Result: false},
		{Domain: "*", Result: false},
	}

	for i := range cases {
		if r := DomainPart(cases[i].Domain); r != cases[i].Result {
			t.Fatalf("FAIL: %s is %v", cases[i].Domain, r)
		}
	}
}

func BenchmarkDomainPart(b *testing.B) {

	for i := 0; i < b.N; i++ {
		DomainPart("elmasy")
	}
}
