package domain

import (
	"testing"
)

func TestIsValid(t *testing.T) {

	if !IsValid("elmasy.com") {
		t.Errorf("elmasy.com is invalid\n")
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

func TestGetTLD(t *testing.T) {

	// 1. element = test string
	// 2. element = wanted result
	cases := [][2]string{
		{"", ""},
		{".", ""},
		{"com", "com"},
		{"com.", "com"},
		{".com", ""},
		{".com.", ""},
		{"elmasy.com", "com"},
		{"elmasy.com.", "com"},
		{".elmasy.com", "com"},
		{".elmasy.com.", "com"},
		{"test.test.elmasy.com", "com"},
		{"test.test.elmasy.com.", "com"},
	}

	for i := range cases {
		tld := GetTLD(cases[i][0])
		if tld != cases[i][1] {
			t.Errorf("Case: %s, want: %s, got: %s\n", cases[i][0], cases[i][1], tld)
		}
	}
}

func BenchmarkGetTLD(b *testing.B) {

	for i := 0; i < b.N; i++ {
		GetTLD("test.elmasy.com.")
	}
}

func TestGetDomain(t *testing.T) {

	// 1. element = test string
	// 2. element = wanted result
	cases := [][2]string{
		{"", ""},
		{".", ""},
		{"com", ""},
		{"com.", ""},
		{".com", ""},
		{".com.", ""},
		{"elmasy.com", "elmasy.com"},
		{"elmasy.com.", "elmasy.com"},
		{".elmasy.com", ""},
		{".elmasy.com.", ""},
		{"test.test.elmasy.com", "elmasy.com"},
		{"test.test.elmasy.com.", "elmasy.com"},
	}

	for i := range cases {
		tld := GetDomain(cases[i][0])
		if tld != cases[i][1] {
			t.Errorf("Case: %s, want: %s, got: %s\n", cases[i][0], cases[i][1], tld)
		}
	}
}

func BenchmarkGetDomain(b *testing.B) {

	for i := 0; i < b.N; i++ {
		GetDomain("test.elmasy.com.")
	}
}

func TestGetSub(t *testing.T) {

	// 1. element = test string
	// 2. element = wanted result
	cases := [][2]string{
		{"", ""},
		{".", ""},
		{"com", ""},
		{"com.", ""},
		{".com", ""},
		{".com.", ""},
		{"elmasy.com", ""},
		{"elmasy.com.", ""},
		{".elmasy.com", ""},
		{".elmasy.com.", ""},
		{"test.test.elmasy.com", "test.test"},
		{"test.test.elmasy.com.", "test.test"},
	}

	for i := range cases {
		tld := GetSub(cases[i][0])
		if tld != cases[i][1] {
			t.Errorf("Case: %s, want: %s, got: %s\n", cases[i][0], cases[i][1], tld)
		}
	}
}

func BenchmarkGetSub(b *testing.B) {

	for i := 0; i < b.N; i++ {
		GetSub("test.elmasy.com.")
	}
}

func TestIsWildcard(t *testing.T) {

	if !IsWildcard("*.elmasy.com") {
		t.Errorf("*.elmasy.com should be a wildcard\n")
	}

	if IsWildcard("test.elmasy.com") {
		t.Errorf("test.elmasy.com should NOT be a wildcard\n")
	}

	if IsWildcard(".elmasy.com") {
		t.Errorf(".elmasy.com should NOT be a wildcard\n")
	}
}

func BenchmarkIsWildCard(b *testing.B) {

	for i := 0; i < b.N; i++ {
		IsWildcard("test.elmasy.com.")
	}
}
