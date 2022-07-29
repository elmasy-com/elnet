package domain

import "strings"

// IsValid checks if a string is a presentation-format domain name
// (currently restricted to hostname-compatible "preferred name" LDH labels and
// SRV-like "underscore labels".
func IsValid(s string) bool {

	/*
		A copy from // A copy from https://github.com/golang/go/blob/3e387528e54971d6009fe8833dcab6fc08737e04/src/net/dnsclient.go#L78
	*/

	// The root domain name is valid. See golang.org/issue/45715.
	if s == "." {
		return true
	}

	// See RFC 1035, RFC 3696.
	// Presentation format has dots before every label except the first, and the
	// terminal empty label is optional here because we assume fully-qualified
	// (absolute) input. We must therefore reserve space for the first and last
	// labels' length octets in wire format, where they are necessary and the
	// maximum total length is 255.
	// So our _effective_ maximum is 253, but 254 is not rejected if the last
	// character is a dot.
	l := len(s)
	if l == 0 || l > 254 || l == 254 && s[l-1] != '.' {
		return false
	}

	last := byte('.')
	nonNumeric := false // true once we've seen a letter or hyphen
	partlen := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
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

// GetTLD returns the Top-Level Domain of the given domain s.
// Returns an empty string ("") is s is not a valid domain or tld not found.
func GetTLD(s string) string {

	if !IsValid(s) {
		return ""
	}

	parts := strings.Split(s, ".")

	if len(parts) < 2 {
		return ""
	}

	return parts[1]
}

// GetSub returns the Subdomain (Third Level Domain) of the given domain s.
// Returns an empty string ("") is s is not a valid domain or the subdomain not found.
func GetSub(s string) string {

	if !IsValid(s) {
		return ""
	}

	parts := strings.Split(s, ".")

	if len(parts) < 3 {
		return ""
	}

	return parts[0]
}
