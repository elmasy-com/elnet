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

	if !IsValid(".") {
		t.Errorf("\".\" is invalid!\n")
	}

	if IsValid("a a") {
		t.Errorf("\"a a\" is valid!\n")
	}

	if IsValid("a=a") {
		t.Errorf("\"a=a\" is valid!\n")
	}
}

func BenchmarkIsValid(b *testing.B) {

	for i := 0; i < b.N; i++ {
		IsValid("test.elmasy.com.")
	}
}

func TestGetDotsIndex(t *testing.T) {

	is := getDotsIndex("Hello. 世界")
	if len(is) != 1 || is[0] != 5 {
		t.Fatalf("FAIL: \"Hello. 世界\", want [5], got %v\n", is)
	}

	is = getDotsIndex("Hello. 世界. world")
	if len(is) != 2 || is[0] != 5 || is[1] != 13 {
		t.Fatalf("FAIL: \"Hello. 世界. world\", want [5,13], got %v\n", is)
	}

	is = getDotsIndex("dot.dot.dot.dot")
	if len(is) != 3 || is[0] != 3 || is[1] != 7 || is[2] != 11 {
		t.Fatalf("FAIL: \"dot.dot.dot.dot\", want [3,7,11], got %#v\n", is)
	}

	is = getDotsIndex("...")
	if len(is) != 3 || is[0] != 0 || is[1] != 1 || is[2] != 2 {
		t.Fatalf("FAIL: \"...\", want [0,1,2], got %#v\n", is)
	}

	is = getDotsIndex("dot")
	if is != nil {
		t.Fatalf("FAIL: \"dot\", want nil, got %#v\n", is)
	}
}

func TestGetTLD(t *testing.T) {

	// 1. element = test string
	// 2. element = wanted result
	cases := [][2]string{
		{"", ""},
		{".", ""},
		{"a.", "a"},
		{".a", ""},
		{"com.", "com"},
		{".com", ""},
		{"cromulent", "cromulent"},
		{"a.0emm.com", "com"},
		{"0emm.com", "com"},
		{"amazon.co.uk", "co.uk"},
		{"books.amazon.co.uk", "co.uk"},
		{"amazon.com", "com"},
		{"example0.debian.net", "net"},
		{"example1.debian.org", "org"},
		{"golang.dev", "dev"},
		{"golang.net", "net"},
		{"play.golang.org", "org"},
		{"gophers.in.space.museum", "space.museum"},
		{"b.c.d.0emm.com", "com"},
		{"there.is.no.such-tld", "such-tld"},
		{"foo.org", "org"},
		{"foo.co.uk", "co.uk"},
		{"foo.dyndns.org", "org"},
		{"www.foo.dyndns.org", "org"},
		{"foo.blogspot.co.uk", "co.uk"},
		{"www.foo.blogspot.co.uk", "co.uk"},
		{"test.com.test.com", "com"},
		{"test.com.", "com"},
		{"test.com.test.com.", "com"},
		{"s3.ca-central-1.amazonaws.com", "com"},
		{"www.test.r.appspot.com", "com"},
		{"test.blogspot.com", "com"},
	}

	for i := range cases {
		tld := GetTLD(cases[i][0])
		if tld != cases[i][1] {
			t.Fatalf("Case: '%s', want: '%s', got: '%s'\n", cases[i][0], cases[i][1], tld)
		}
	}
}

func BenchmarkGetTLD(b *testing.B) {

	for i := 0; i < b.N; i++ {
		GetTLD("test.example.co.uk.")
	}
}

func TestGetDomain(t *testing.T) {

	// 1. element = test string
	// 2. element = wanted result
	cases := [][2]string{
		{"", ""},
		{".", ""},
		{".cromulent", ""},
		{"a.0emm.com", "0emm.com"}, // a.0emm.com is a TLD as per publicsuffix
		{"0emm.com", "0emm.com"},   // 0emm.com is not a TLD, only *.0emm.com
		{"amazon.co.uk", "amazon.co.uk"},
		{"books.amazon.co.uk", "amazon.co.uk"},
		{"amazon.com", "amazon.com"},
		{"example0.debian.net", "debian.net"},
		{"example1.debian.org", "debian.org"},
		{"golang.dev", "golang.dev"},
		{"golang.net", "golang.net"},
		{"play.golang.org", "golang.org"},
		{"gophers.in.space.museum", "in.space.museum"},
		{"b.c.d.0emm.com", "0emm.com"},
		{"there.is.no.such-tld", "no.such-tld"},
		{"foo.org", "foo.org"},
		{"foo.co.uk", "foo.co.uk"},
		{"foo.dyndns.org", "dyndns.org"},
		{"www.foo.dyndns.org", "dyndns.org"},
		{"foo.blogspot.co.uk", "blogspot.co.uk"},
		{"www.foo.blogspot.co.uk", "blogspot.co.uk"},
		{"test.com.test.com", "test.com"},
		{"test.com.", "test.com"},
		{"test.com.test.com.", "test.com"},
		{"s3.ca-central-1.amazonaws.com", "amazonaws.com"},
		{"www.test.r.appspot.com", "appspot.com"},
		{"test.blogspot.com", "blogspot.com"},
	}

	for i := range cases {
		tld := GetDomain(cases[i][0])
		if tld != cases[i][1] {
			t.Fatalf("Case: %s, want: %s, got: %s\n", cases[i][0], cases[i][1], tld)
		}
	}
}

func BenchmarkGetDomain(b *testing.B) {

	for i := 0; i < b.N; i++ {
		GetDomain("test.example.co.uk.")
	}
}

func TestGetSub(t *testing.T) {

	// 1. element = test string
	// 2. element = wanted result
	cases := [][2]string{
		{"", ""},
		{".", ""},
		{".cromulent", ""},
		{"a.0emm.com", "a"}, // a.0emm.com is a TLD as per publicsuffix
		{"0emm.com", ""},    // 0emm.com is not a TLD, only *.0emm.com
		{"amazon.co.uk", ""},
		{"books.amazon.co.uk", "books"},
		{"amazon.com", ""},
		{"example0.debian.net", "example0"},
		{"example1.debian.org", "example1"},
		{"golang.dev", ""},
		{"golang.net", ""},
		{"play.golang.org", "play"},
		{"gophers.in.space.museum", "gophers"},
		{"b.c.d.0emm.com", "b.c.d"},
		{"there.is.no.such-tld", "there.is"},
		{"foo.org", ""},
		{"foo.co.uk", ""},
		{"foo.dyndns.org", "foo"},
		{"www.foo.dyndns.org", "www.foo"},
		{"foo.blogspot.co.uk", "foo"},
		{"www.foo.blogspot.co.uk", "www.foo"},
		{"test.com.test.com", "test.com"},
		{"test.com.", ""},
		{"test.com.test.com.", "test.com"},
		{"s3.ca-central-1.amazonaws.com", "s3.ca-central-1"},
		{"www.test.r.appspot.com", "www.test.r"},
		{"test.blogspot.com", "test"},
	}

	for i := range cases {
		tld := GetSub(cases[i][0])
		if tld != cases[i][1] {
			t.Fatalf("Case: %s, want: %s, got: %s\n", cases[i][0], cases[i][1], tld)
		}
	}
}

func BenchmarkGetSub(b *testing.B) {

	for i := 0; i < b.N; i++ {
		GetSub("test.example.co.uk.")
	}
}
