package domain

import (
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
	case d == ".":
		// The root domain name is valid. See golang.org/issue/45715.
		return true
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

func getDotsIndex(d string) []int {

	var indexes []int = nil

	for i := 0; i < len(d); i++ {
		if d[i] == '.' {
			indexes = append(indexes, i)
		}
	}

	return indexes
}

// GetTLD returns the TLD of d (eg.: sub.exmaple.com -> com).
// If d is invalid or cant get the TLD returns an empty string ("").
func GetTLD(d string) string {

	if d == "" || d == "." || d[0] == '.' {
		return ""
	}

	if d[len(d)-1] == '.' {
		d = d[:len(d)-1]
	}

	di := getDotsIndex(d)
	lenDi := len(di)

	// Store the result.
	// If di == nil, the the result will be d (the for loop will be skipped).
	result := d

	// Iterate in reverse order over the dot indexes
	for i := lenDi - 1; i >= 0; i-- {

		// dot index + 1 to skip dot in the string
		v := di[i] + 1

		tld, icann := publicsuffix.PublicSuffix(d[v:])

		// The returned TLD is ICANN managed and differs from d[v:]
		// This means, that d[v:] was a domain, and PublicSuffix() founded the TLD
		if icann && tld != d[v:] {
			break
		}

		// We got an unmanaged TLD
		if !icann {
			// If this was the first iteration, save the resulted TLD, because result contains d.
			if i == lenDi-1 {
				result = tld
			}
			break
		}

		result = tld
	}

	return result
}

// GetDomain returns the domain of d (eg.: sub.example.com -> example.com).
// If d is invalid or cant get the domain returns an empty string ("").
func GetDomain(d string) string {

	if d == "" || d == "." || d[0] == '.' {
		return ""
	}

	if d[len(d)-1] == '.' {
		d = d[:len(d)-1]
	}

	tld := GetTLD(d)
	switch tld {
	case "":
		// Cannot get TLD, d is invalid
		return ""
	case d:
		// s is a TLD
		return ""
	}

	if d[len(d)-1] == '.' {
		d = d[:len(d)-1]
	}

	// Get the index of the TLD -1 to skip the dot
	i := len(d) - len(tld) - 1

	return d[1+strings.LastIndex(d[:i], "."):]
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
