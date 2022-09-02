package domain

import (
	"bytes"
	"testing"
)

func TestIsValid(t *testing.T) {

	if !IsValid("elmasy.com") {
		t.Errorf("elmasy.com is invalid\n")
	}

	if IsValid("aaaaaa") {
		t.Errorf("aaaaaa is valid!\n")
	}
}

func TestGetTLD(t *testing.T) {

	tld := GetTLD("test.test.elmasy.com")

	if tld != "com" {
		t.Errorf("TLD not found, result: \"%s\"\n", tld)
	}
}

func TestGetTLDbytes(t *testing.T) {

	tld := GetTLDBytes([]byte("test.test.elmasy.com"))

	if !bytes.Equal(tld, []byte("com")) {
		t.Errorf("TLD bytes not found, result: \"%v\"\n", tld)
	}
}

func TestGetSub(t *testing.T) {

	sub := GetSub("test.test.test.elmasy.com")

	if sub != "test.test.test" {
		t.Errorf("subdomain not found, result: \"%s\"\n", sub)
	}

	sub = GetSub(".elmasy.com")

	if sub != "" {
		t.Errorf("subdomain not found, result: \"%s\"\n", sub)
	}
}

func TestGetSubBytes(t *testing.T) {

	sub := GetSubBytes([]byte("test.test.test.elmasy.com"))

	if !bytes.Equal(sub, []byte("test.test.test")) {
		t.Errorf("subdomain bytes not found, result: \"%v\"\n", sub)
	}

	sub = GetSubBytes([]byte(".elmasy.com"))

	if sub != nil {
		t.Errorf("subdomain not found, result: \"%s\"\n", sub)
	}
}

func TestGetDomain(t *testing.T) {

	sub := GetDomain("test.test.test.elmasy.com")

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
