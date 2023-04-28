package domain

import (
	"errors"
	"strings"

	"golang.org/x/net/publicsuffix"
)

type Parts struct {
	TLD    string // Top level domain (eg.: "com"). Cant be empty.
	Domain string // Domain part (eg.: example"). Cant be empty.
	Sub    string // Subdomain part (eg.: "www"). Can be empty.
}

var (
	ErrInvalidDomain = errors.New("invalid domain")
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
	case d == ".":
		// The root domain name is technically valid. See golang.org/issue/45715.
		// But not for this package
		return false
	case d[0] == '.':
		// Mising label, the domain name cant start with a dot.
		return false
	}

	containsDot := false
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
			containsDot = true
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

	if !containsDot {
		return false
	}

	return nonNumeric
}

// Clean removes the trailing dot and returns a lower cased version of d.
func Clean(d string) string {

	// Remove the trailing dot.
	if d[len(d)-1] == '.' {
		d = d[:len(d)-1]
	}

	return strings.ToLower(d)
}

// GetParts returns the parts of d.
//
// If d is just a TLD, returns a struct with empty Domain (eg.: "com" -> &Result{Sub: "", Domain: "", TLD: "com"}).
//
// If d does not contains subdomain, than Sub will be empty (eg.: "example.com" -> &Result{Sub: "", Domain: "example", TLD: "com"}).
//
// Returns nil if d is empty, a dot (".") or starts with a dot (eg.: ".example.com").
//
// NOTE: This function does not validate and Clean() the given domain d. It is recommended to use IsValid() and Clean() before this function.
func GetParts(d string) *Parts {

	if d == "" || d == "." || d[0] == '.' {
		return nil
	}

	if d[len(d)-1] == '.' {
		d = d[:len(d)-1]
	}

	tldIndex := GetTLDIndex(d)
	if tldIndex <= 0 {
		return &Parts{TLD: d}
	}

	tldIndex -= 1
	domIndex := 0

	for i := 0; i < tldIndex; i++ {
		if d[i] == '.' {
			domIndex = i
		}
	}

	if domIndex == 0 {
		return &Parts{Sub: "", Domain: d[0:tldIndex], TLD: d[tldIndex+1:]}
	}

	return &Parts{Sub: d[0:domIndex], Domain: d[domIndex+1 : tldIndex], TLD: d[tldIndex+1:]}
}

// GetTLD returns the Top Level Domain of d (eg.: sub.exmaple.com -> com).
//
// Returns an empty string ("") if d is empty, a dot (".") or starts with a dot (eg.: ".example.com").
func GetTLD(d string) string {

	if d == "" || d == "." || d[0] == '.' {
		return ""
	}

	if d[len(d)-1] == '.' {
		d = d[:len(d)-1]
	}

	tld := d
	icann := false

	for {
		tld, icann = publicsuffix.PublicSuffix(tld)
		// Break if ICANN managed
		if icann {
			break
		}

		dot := strings.IndexByte(tld, '.')

		// No dot means the TLD (eg.: "com")
		if dot == -1 {
			break
		}

		// Get the next part of the domain
		tld = tld[dot+1:]
	}

	return tld
}

// GetTLDIndex returns the index of the Top Level Domain in d (eg.: sub.example.com -> 12).
//
// Returns 0 if d is a TLD.
//
// Returns -1 if d is empty, a dot (".") or starts with a dot (eg.: ".example.com").
func GetTLDIndex(d string) int {

	if d == "" || d == "." || d[0] == '.' {
		return -1
	}

	if d[len(d)-1] == '.' {
		d = d[:len(d)-1]
	}

	//return strings.LastIndex(d, GetTLD(d))
	return len(d) - len(GetTLD(d))
}

// GetDomain returns the domain of d (eg.: sub.example.com -> example.com).
//
// Returns an empty string ("") if d is empty, a dot ("."), starts with a dot (eg.: ".example.com") or d is just a TLD.
func GetDomain(d string) string {

	if d == "" || d == "." || d[0] == '.' {
		return ""
	}

	if d[len(d)-1] == '.' {
		d = d[:len(d)-1]
	}

	tld := GetTLD(d)

	if tld == "" || tld == d {
		return ""
	}

	// Get the index of the TLD -1 to skip the dot
	i := len(d) - len(tld) - 1

	return d[1+strings.LastIndex(d[:i], "."):]
}

// GetDomainIndex returns the index of the domain of d (eg.: sub.example.com -> 4).
//
// Returns -1 if d is empty, a dot ("."), starts with a dot (eg.: ".example.com") or d is just a TLD.
func GetDomainIndex(d string) int {

	if d == "" || d == "." || d[0] == '.' {
		return -1
	}

	if d[len(d)-1] == '.' {
		d = d[:len(d)-1]
	}

	dom := GetDomain(d)
	if dom == "" {
		return -1
	}

	return len(d) - len(dom)
}

// GetSub returns the Subdomain of the given domain d (eg.: eg.: sub.example.com -> example.com).
// If d is invalid or cant get the subdomain returns an empty string ("").
func GetSub(d string) string {

	s := GetDomain(d)
	if s == "" {
		return ""
	}

	// Not error, but no subdomain
	i := strings.LastIndex(d, s)
	if i <= 0 {
		return ""
	}

	return d[:i-1]
}
