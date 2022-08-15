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
}