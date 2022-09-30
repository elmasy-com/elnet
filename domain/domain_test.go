package domain

import (
	"testing"
)

func TestIsValid(t *testing.T) {

	if !IsValid("elmasy.com") {
		t.Errorf("elmasy.com is invalid\n")
	}

	if !IsValid("elmasy.com.") {
		t.Errorf("elmasy.com. is invalid\n")
	}

	if IsValid(".elmasy.com") {
		t.Errorf(".elmasy.com is valid\n")
	}

	if IsValid(".elmasy.com.") {
		t.Errorf(".elmasy.com. is valid\n")
	}

	if IsValid("aaaaaa") {
		t.Errorf("aaaaaa is valid!\n")
	}

	if IsValid("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.elmasy.com") {
		t.Errorf("long aaaaaa is valid\n")
	}

	if IsValid("") {
		t.Errorf("empty is valid!\n")
	}
}

func BenchmarkIsValid(b *testing.B) {

	for i := 0; i < b.N; i++ {
		IsValid("test.elmasy.com.")
	}
}

func TestGetDomain(t *testing.T) {

	// 1. element = test string
	// 2. element = wanted result
	cases := [][2]string{
		{"", "error"},
		{".", "error"},
		{".cromulent", "error"},
		{"a.0emm.com", "error"},  // a.0emm.com is a TLD as per publicsuffix
		{"0emm.com", "0emm.com"}, // 0emm.com is not a TLD, only *.0emm.com
		{"amazon.co.uk", "amazon.co.uk"},
		{"books.amazon.co.uk", "amazon.co.uk"},
		{"amazon.com", "amazon.com"},
		{"example0.debian.net", "example0.debian.net"},
		{"example1.debian.org", "debian.org"},
		{"golang.dev", "golang.dev"},
		{"golang.net", "golang.net"},
		{"play.golang.org", "golang.org"},
		{"gophers.in.space.museum", "in.space.museum"},
		{"b.c.d.0emm.com", "c.d.0emm.com"},
		{"there.is.no.such-tld", "no.such-tld"},
		{"foo.org", "foo.org"},
		{"foo.co.uk", "foo.co.uk"},
		{"foo.dyndns.org", "foo.dyndns.org"},
		{"www.foo.dyndns.org", "foo.dyndns.org"},
		{"foo.blogspot.co.uk", "foo.blogspot.co.uk"},
		{"www.foo.blogspot.co.uk", "foo.blogspot.co.uk"},
		{"test.com.test.com", "test.com"},
		{"test.com.", "test.com"},
		{"test.com.test.com.", "test.com"},
	}

	for i := range cases {
		tld, err := GetDomain(cases[i][0])
		if err != nil {
			if cases[i][1] == "error" {
				continue
			}
			t.Fatalf("Case: %s, want: %s, got: %s\n", cases[i][0], cases[i][1], err)
		}
		if tld != cases[i][1] {
			t.Fatalf("Case: %s, want: %s, got: %s\n", cases[i][0], cases[i][1], tld)
		}
	}
}

func BenchmarkGetDomain(b *testing.B) {

	for i := 0; i < b.N; i++ {
		GetDomain("test.elmasy.com")
	}
}

func TestGetSub(t *testing.T) {

	// 1. element = test string
	// 2. element = wanted result
	cases := [][2]string{
		{"", "error"},
		{".", "error"},
		{".cromulent", "error"},
		{"a.0emm.com", "error"}, // a.0emm.com is a TLD as per publicsuffix
		{"0emm.com", ""},        // 0emm.com is not a TLD, only *.0emm.com
		{"amazon.co.uk", ""},
		{"books.amazon.co.uk", "books"},
		{"amazon.com", ""},
		{"example0.debian.net", ""},
		{"example1.debian.org", "example1"},
		{"golang.dev", ""},
		{"golang.net", ""},
		{"play.golang.org", "play"},
		{"gophers.in.space.museum", "gophers"},
		{"b.c.d.0emm.com", "b"},
		{"there.is.no.such-tld", "there.is"},
		{"foo.org", ""},
		{"foo.co.uk", ""},
		{"foo.dyndns.org", ""},
		{"www.foo.dyndns.org", "www"},
		{"foo.blogspot.co.uk", ""},
		{"www.foo.blogspot.co.uk", "www"},
		{"test.com.test.com", "test.com"},
		{"test.com.", ""},
		{"test.com.test.com.", "test.com"},
	}

	for i := range cases {
		tld, err := GetSub(cases[i][0])
		if err != nil {
			if cases[i][1] == "error" {
				continue
			}
			t.Fatalf("Case: %s, want: %s, got: %s\n", cases[i][0], cases[i][1], err)
		}
		if tld != cases[i][1] {
			t.Fatalf("Case: %s, want: %s, got: %s\n", cases[i][0], cases[i][1], tld)
		}
	}
}

func BenchmarkGetSub(b *testing.B) {

	for i := 0; i < b.N; i++ {
		GetSub("test.elmasy.co.uk.")
	}
}
