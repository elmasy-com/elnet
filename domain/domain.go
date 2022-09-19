package domain

import (
	"strings"
)

// IsValid checks if a ByteSeq is a presentation-format domain name
// (currently restricted to hostname-compatible "preferred name" LDH labels and
// SRV-like "underscore labels".
func IsValid(d string) bool {

	/*
		The base is a copy from // A copy from https://github.com/golang/go/blob/3e387528e54971d6009fe8833dcab6fc08737e04/src/net/dnsclient.go#L78
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

// IsReservedSecondLevel resturns whether d is a reserved second level domain (eg.: co.uk).
func IsReservedSecondLevel(tld string) bool {

	// co.X
	if tld == "co.uk" || tld == "co.jp" || tld == "co.kr" || tld == "co.th" || tld == "co.za" ||

		// com.X
		tld == "com.br" || tld == "com.my" || tld == "com.tr" || tld == "com.pl" || tld == "com.tw" || tld == "com.ng" || tld == "com.au" ||

		// org.X
		tld == "org.uk" {
		return true
	}

	return false
}

// GetDomain returns the domain of d (eg.: sub.example.com -> example.com).
// Returns an empty string if domain not found.
// This function returns a slice of d, does not allocate a new string.
func GetDomain(d string) string {

	i := len(d) // The index of the occurence of  the dot

	// Need at least 3 character to get a domain (eg.: a.a)
	if i < 3 {
		return ""
	}

	// Domain names with a dot at the end are valid.
	// Remove the last dot from the string.
	if d[i-1] == '.' {
		d = d[:i-1]
		i--
	}

	// Get the first dot, the TLD
	i = strings.LastIndexByte(d[:i], '.')
	if i == -1 {
		// No dot in d, so cant get the domain name: d is invalid.
		return ""
	}

	// Nothing before the dot: d is invalid (eg.: .com)
	if d[:i] == "" {
		return ""
	}

	// Get the second dot, the domain
	i = strings.LastIndexByte(d[:i], '.')
	if i == -1 {
		// The second dot not found, so d has only one dot and two fields (eg.: elmasy.com)
		// d can be a reserved second level domain (eg.: co.uk) or a valid domain name.
		if IsReservedSecondLevel(d) {
			return ""
		}

		return d
	}

	// Nothing before the second dot: d is invalid (eg.: .elmasy.com)
	if d[:i] == "" {
		return ""
	}

	// Check reserved second level domain with the second dot's index.
	if IsReservedSecondLevel(d[i+1:]) {
		// d is including a reserved second level domain, so we need the third dot.
		// Get the third dot, the domain before a reserved second level domain (eg.: elmasy.co.uk)
		i = strings.LastIndexByte(d[:i], '.')
		if i == -1 {
			// The third dot not found, so d is a domain (eg.: elmasy.co.uk)
			return d
		}

		// Nothing before the third dot, d is invalid
		if d[:i] == "" {
			return ""
		}
	}

	// The subdomain starts with a dot: d is invalid
	if d[0] == '.' {
		return ""
	}

	return d[i+1:]
}

// GetSub returns the Subdomain (Third Level Domain) of the given domain d.
// Returns an empty string if subdomain not found.
// This function returns a slice of d, does not allocate a new string.
func GetSub(d string) string {

	i := len(d) // The index of the occurence of  the dot

	// Need at least 5 character to get a subdomain (eg.: a.a.a)
	if i < 5 {
		return ""
	}

	// Domain names with a dot at the end are valid.
	// Remove the last dot from the string.
	if d[i-1] == '.' {
		d = d[:i-1]
		i--
	}

	// Get the first dot, the TLD
	i = strings.LastIndexByte(d[:i], '.')
	if i == -1 {
		// No dot in d, so cant get the subdomain: d is invalid.
		return ""
	}

	// Nothing before the dot: d is invalid (eg.: .com)
	if d[:i] == "" {
		return ""
	}

	// Get the second dot, the domain
	i = strings.LastIndexByte(d[:i], '.')
	if i == -1 {
		// No second dot in d, so no subdomain: d is invalid.
		return ""
	}

	// Nothing before the second dot: d is invalid (eg.: .elmasy.com)
	if d[:i] == "" {
		return ""
	}

	// Check reserved second level domain with the second dot's index.
	if IsReservedSecondLevel(d[i+1:]) {
		// d is include a reserved second level domain, so we need the third dot.
		// Get the third dot, to get the domain before a reserved second level domain (eg.: elmasy.co.uk)
		i = strings.LastIndexByte(d[:i], '.')
		if i == -1 {
			// The third dot not found, so no subdomain: d is invalid
			return ""
		}
	}

	// The subdomain starts with a dot: d is invalid
	if d[0] == '.' {
		return ""
	}

	return d[:i]
}

// IsWildcard returns whether the given domain d is a wildcard domain.
func IsWildcard(d string) bool {

	s := GetSub(d)

	return s == "*"
}
