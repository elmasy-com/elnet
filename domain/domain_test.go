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

func TestGetTLD(t *testing.T) {

	tld := GetTLD("test.test.elmasy.com")

	if string(tld) != "com" {
		t.Errorf("TLD not found, result: \"%s\"\n", tld)
	}

	tld = GetTLD("test.test.elmasy.test")

	if string(tld) != "test" {
		t.Errorf("TLD not found, result: \"%s\"\n", tld)
	}

	tld = GetTLD("test.test.elmasy.")

	if tld != nil {
		t.Errorf("TLD found for empty, result: \"%s\"\n", tld)
	}
}

func TestGetSub(t *testing.T) {

	sub := GetSub("test.test.test.elmasy.com")

	if string(sub) != "test.test.test" {
		t.Errorf("subdomain not found, result: \"%s\"\n", sub)
	}

	sub = GetSub(".elmasy.com")

	if sub != nil {
		t.Errorf("subdomain not found, result: \"%s\"\n", sub)
	}
}

func TestGetDomain(t *testing.T) {

	sub := GetDomain("test.test.test.elmasy.com")

	if string(sub) != "elmasy.com" {
		t.Errorf("subdomain not found, result: \"%s\"\n", sub)
	}

	sub = GetDomain("test.test.test.elmasy.")

	if sub != nil {
		t.Errorf("subdomain found for empty tld, result: \"%s\"\n", sub)
	}

	sub = GetDomain("test.test.test..com")

	if sub != nil {
		t.Errorf("subdomain found for empty domain, result: \"%s\"\n", sub)
	}

	sub = GetDomain("test.test.test..")

	if sub != nil {
		t.Errorf("subdomain found for full empty, result: \"%s\"\n", sub)
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
