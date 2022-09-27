package domain

import (
	"fmt"
	"strings"

	"golang.org/x/net/publicsuffix"
)

// IsValid checks if a ByteSeq is a presentation-format domain name
// (currently restricted to hostname-compatible "preferred name" LDH labels and
// SRV-like "underscore labels".
func IsValid(d string) bool {

	/*
	 * The base is a copy from https://github.com/golang/go/blob/3e387528e54971d6009fe8833dcab6fc08737e04/src/net/dnsclient.go#L78
	 */

	l := len(d)

	switch {
	case l == 0 || l > 254 || l == 254 && d[l-1] != '.':
		// See RFC 1035, RFC 3696.
		// Presentation format has dots before every label except the first, and the
		// terminal empty label is optional here because we assume fully-qualified
		// (absolute) input. We must therefore reserve space for the first and last
		// labels' length octets in wire format, where they are necessary and the
		// maximum total length is 255.
		// So our _effective_ maximum is 253, but 254 is not rejected if the last
		// character is a dot.
		return false
	case d[0] == '.':
		// Mising label, the domain name cant start with a dot.
		return false
	case !strings.Contains(d, "."):
		return false
	case strings.Contains(d, " "):
		return false
	case d == ".":
		// The root domain name is valid. See golang.org/issue/45715.
		return true
	}

	last := byte('.')
	nonNumeric := false // true once we've seen a letter or hyphen
	partlen := 0
	for i := 0; i < len(d); i++ {
		c := d[i]
		switch {
		default:
			return false
		case 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_':
			nonNumeric = true
			partlen++
		case '0' <= c && c <= '9':
			// fine
			partlen++
		case c == '-':
			// Byte before dash cannot be dot.
			if last == '.' {
				return false
			}
			partlen++
			nonNumeric = true
		case c == '.':
			// Byte before dot cannot be dot, dash.
			if last == '.' || last == '-' {
				return false
			}
			if partlen > 63 || partlen == 0 {
				return false
			}
			partlen = 0
		}
		last = c
	}
	if last == '-' || partlen > 63 {
		return false
	}

	return nonNumeric
}

// GetDomain returns the domain of d (eg.: sub.example.com -> example.com).
func GetDomain(d string) (string, error) {

	switch {
	case d == "":
		return "", fmt.Errorf("domain is empty")
	case d[len(d)-1] == '.':
		d = d[:len(d)-1]
	}

	return publicsuffix.EffectiveTLDPlusOne(d)
}

// MustGetDomain returns the domain of d (eg.: sub.example.com -> example.com).
// This function panics if failed to get domain.
// Created to use after IsValid().
func MustGetDomain(d string) string {

	v, err := GetDomain(d)
	if err != nil {
		panic(err)
	}

	return v
}

// GetSub returns the Subdomain of the given domain d.
// Return an empty string if no subdomain found, or error if d is an invalid domain.
func GetSub(d string) (string, error) {

	s, err := GetDomain(d)
	if err != nil {
		return "", err
	}

	// Not error, but no subdomain
	if s == d {
		return "", nil
	}

	return d[:strings.Index(d, s)-1], nil
}

// MustGetSub returns the Subdomain of the given domain d.
// This function panics if failed to get subdomain.
// Created to use after IsValid().
func MustGetSub(d string) string {

	s, err := GetDomain(d)
	if err != nil {
		panic(err)
	}

	if s == d {
		panic("no subdomain in " + d)
	}

	i := strings.Index(d, s)

	return d[:i]
}
