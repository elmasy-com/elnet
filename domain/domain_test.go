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

	if IsValid(".") {
		t.Errorf("\".\" is valid!\n")
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

func TestIsValidDLS(t *testing.T) {

	if !IsValidSLD("elmasy") {
		t.Errorf("elmasy is invalid\n")
	}

	if IsValidSLD("elmasy.") {
		t.Errorf("elmasy. is valid\n")
	}

	if IsValidSLD(".elmasy") {
		t.Errorf(".elmasy is valid\n")
	}

	if IsValidSLD(".elmasy.com.") {
		t.Errorf(".elmasy.com. is valid\n")
	}

	if !IsValidSLD("aaaaaa") {
		t.Errorf("aaaaaa is invalid!\n")
	}

	if IsValidSLD("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.elmasy.com") {
		t.Errorf("long aaaaaa is valid\n")
	}

	if IsValidSLD("") {
		t.Errorf("empty is valid!\n")
	}

	if IsValidSLD(".") {
		t.Errorf("\".\" is valid!\n")
	}

	if IsValidSLD("a a") {
		t.Errorf("\"a a\" is valid!\n")
	}

	if IsValidSLD("a=a") {
		t.Errorf("\"a=a\" is valid!\n")
	}

	if !IsValidSLD("a-a") {
		t.Errorf("\"a_a\" is invalid!\n")
	}

	if IsValidSLD("-aa") {
		t.Errorf("\"_aa\" is valid!\n")
	}

	if IsValidSLD("aa-") {
		t.Errorf("\"aa_\" is valid!\n")
	}

	if IsValidSLD("a--a") {
		t.Errorf("\"a__a\" is valid!\n")
	}
}

func BenchmarkIsValidSLD(b *testing.B) {

	for i := 0; i < b.N; i++ {
		IsValidSLD("test.elmasy.com.")
	}
}
func TestGetParts(t *testing.T) {

	// 0. element = test domain
	// 1. element = tld
	// 2. element = domain
	// 3. element = sub
	cases := [][4]string{
		{"com", "com", "", ""},
		{"a.0emm.com", "com", "0emm", "a"},
		{"0emm.com", "com", "0emm", ""},
		{"amazon.co.uk", "co.uk", "amazon", ""},
		{"books.amazon.co.uk", "co.uk", "amazon", "books"},
		{"amazon.com", "com", "amazon", ""},
		{"example0.debian.net", "net", "debian", "example0"},
		{"example1.debian.org", "org", "debian", "example1"},
		{"golang.dev", "dev", "golang", ""},
		{"golang.net", "net", "golang", ""},
		{"play.golang.org", "org", "golang", "play"},
		{"gophers.in.space.museum", "space.museum", "in", "gophers"},
		{"b.c.d.0emm.com", "com", "0emm", "b.c.d"},
		{"there.is.no.such-tld", "such-tld", "no", "there.is"},
		{"foo.org", "org", "foo", ""},
		{"foo.co.uk", "co.uk", "foo", ""},
		{"foo.dyndns.org", "org", "dyndns", "foo"},
		{"www.foo.dyndns.org", "org", "dyndns", "www.foo"},
		{"foo.blogspot.co.uk", "co.uk", "blogspot", "foo"},
		{"www.foo.blogspot.co.uk", "co.uk", "blogspot", "www.foo"},
		{"test.com.test.com", "com", "test", "test.com"},
		{"s3.ca-central-1.amazonaws.com", "com", "amazonaws", "s3.ca-central-1"},
		{"www.test.r.appspot.com", "com", "appspot", "www.test.r"},
		{"test.blogspot.commmmm", "commmmm", "blogspot", "test"},
		{"test.blogspot.colu", "colu", "blogspot", "test"},
		{"test.blogspot.bgtrfesw.bgtrfesw.bgtrfesw", "bgtrfesw", "bgtrfesw", "test.blogspot.bgtrfesw"},
	}

	for i := range cases {

		parts := GetParts(cases[i][0])

		switch {

		case parts.TLD != cases[i][1]:
			t.Fatalf("FAIL: TLD failed with %s, want=%s get=%s\n", cases[i][0], cases[i][1], parts.TLD)
		case parts.Domain != cases[i][2]:
			t.Fatalf("FAIL: Domain failed with %s, want=%s get=%s\n", cases[i][0], cases[i][2], parts.Domain)
		case parts.Sub != cases[i][3]:
			t.Fatalf("FAIL: Sub failed with %s, want=%s get=%s\n", cases[i][0], cases[i][3], parts.Sub)
		}
	}
}

func BenchmarkGetParts(b *testing.B) {

	for i := 0; i < b.N; i++ {
		GetParts("test.s3.dualstack.ap-northeast-2.amazonaws.com.")
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
		{"co.uk", "co.uk"},
		{"co.uk.", "co.uk"},
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
		GetTLD("test.s3.dualstack.ap-northeast-2.amazonaws.com")
	}
}

func TestGetTLDIndex(t *testing.T) {

	// 1. element = test string
	// 2. element = wanted result
	cases := [][2]string{
		{"", ""},
		{".", ""},
		{"a.", "a."},
		{".a", ""},
		{"com.", "com."},
		{".com", ""},
		{"co.uk", "co.uk"},
		{"co.uk.", "co.uk."},
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
		{"test.com.", "com."},
		{"test.com.test.com.", "com."},
		{"s3.ca-central-1.amazonaws.com", "com"},
		{"www.test.r.appspot.com", "com"},
		{"test.blogspot.com", "com"}}

	for i := range cases {
		tld := GetTLDIndex(cases[i][0])

		if tld == -1 {
			if cases[i][1] != "" {
				t.Fatalf("Case: '%s', want: '%s', index: %d\n", cases[i][0], cases[i][1], tld)
			}
			continue
		}

		if cases[i][0][tld:] != cases[i][1] {
			t.Fatalf("Case: '%s', want: '%s', got: '%s', index: %d\n", cases[i][0], cases[i][1], cases[i][0][tld:], tld)
		}
	}
}

func BenchmarkGetTLDIndex(b *testing.B) {

	for i := 0; i < b.N; i++ {
		GetTLDIndex("test.s3.dualstack.ap-northeast-2.amazonaws.com")
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
		GetDomain("test.s3.dualstack.ap-northeast-2.amazonaws.com.")
	}
}

func TestGetDomainIndex(t *testing.T) {

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
		{"test.com.", "test.com."},
		{"test.com.test.com.", "test.com."},
		{"s3.ca-central-1.amazonaws.com", "amazonaws.com"},
		{"www.test.r.appspot.com", "appspot.com"},
		{"test.blogspot.com", "blogspot.com"},
	}

	for i := range cases {
		tld := GetDomainIndex(cases[i][0])

		if tld == -1 {
			if cases[i][1] != "" {
				t.Fatalf("Case: '%s', want: '%s', index: %d\n", cases[i][0], cases[i][1], tld)
			}
			continue
		}

		if cases[i][0][tld:] != cases[i][1] {
			t.Fatalf("Case: '%s', want: '%s', got: '%s', index: %d\n", cases[i][0], cases[i][1], cases[i][0][tld:], tld)
		}
	}
}

func BenchmarkGetDomainIndex(b *testing.B) {

	for i := 0; i < b.N; i++ {
		GetDomainIndex("test.s3.dualstack.ap-northeast-2.amazonaws.com.")
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
		GetSub("test.s3.dualstack.ap-northeast-2.amazonaws.com.")
	}
}
