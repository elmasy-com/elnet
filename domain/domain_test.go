package domain

import "testing"

func TestIsValid(t *testing.T) {

	if !IsValid("elmasy.com") {
		t.Errorf("elmasy.com is invalid\n")
	}

	if IsValid("aaaaaa") {
		t.Errorf("aaaaaa is valid!\n")
	}
}

func TestGetTLD(t *testing.T) {

	tld := GetTLD("elmasy.com")

	if tld != "com" {
		t.Errorf("TLD not found, result: \"%s\"\n", tld)
	}
}

func TestGetSub(t *testing.T) {

	sub := GetSub("test.elmasy.com")

	if sub != "test" {
		t.Errorf("subdomain not found, result: \"%s\"\n", sub)
	}

	sub = GetSub(".elmasy.com")

	if sub != "" {
		t.Errorf("subdomain not found, result: \"%s\"\n", sub)
	}
}

func TestGetDomain(t *testing.T) {

	sub := GetDomain("test.elmasy.com")

	if sub != "elmasy.com" {
		t.Errorf("subdomain not found, result: \"%s\"\n", sub)
	}
}

func TestIsWildcard(t *testing.T) {

	if !IsWildcard("*.elmasy.com") {
		t.Errorf("*.elmasy.com should be a wildcard\n")
	}

	if IsWildcard("test.elmasy.com") {
		t.Errorf("test.elmasy.com should NOT be a wildcard\n")
	}
}
