package domain

import (
	"bytes"
)

type ByteSeq interface {
	[]byte | string
}

// IsValid checks if a ByteSeq is a presentation-format domain name
// (currently restricted to hostname-compatible "preferred name" LDH labels and
// SRV-like "underscore labels".
func IsValid[T ByteSeq](d T) bool {

	/*
		The base is a copy from // A copy from https://github.com/golang/go/blob/3e387528e54971d6009fe8833dcab6fc08737e04/src/net/dnsclient.go#L78
	*/

	var (
		s []byte
		l int
	)

	switch v := any(d).(type) {
	case []byte:
		s = v
	case string:
		s = []byte(v)
	default:
		panic("Invalid type")
	}

	l = len(s)

	switch {
	case l == 0 || l > 254 || l == 254 && s[l-1] != '.':
		// See RFC 1035, RFC 3696.
		// Presentation format has dots before every label except the first, and the
		// terminal empty label is optional here because we assume fully-qualified
		// (absolute) input. We must therefore reserve space for the first and last
		// labels' length octets in wire format, where they are necessary and the
		// maximum total length is 255.
		// So our _effective_ maximum is 253, but 254 is not rejected if the last
		// character is a dot.
		return false
	case !bytes.Contains(s, []byte{byte('.')}):
		return false
	case bytes.Contains(s, []byte{' '}):
		return false
	case bytes.Equal(s, []byte{'.'}):
		// The root domain name is valid. See golang.org/issue/45715.
		return true
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

// GetTLD returns the Top-Level Domain of the given domain d.
// Returns nil if tld not found.
func GetTLD[T ByteSeq](d T) []byte {

	var b []byte

	switch v := any(d).(type) {
	case string:
		b = []byte(v)
	case []byte:
		b = v
	default:
		panic("Invalid type")
	}

	parts := bytes.Split(b, []byte{'.'})

	if len(parts) < 2 {
		return nil
	}

	if len(parts[len(parts)-1]) == 0 {
		return nil
	}

	return parts[len(parts)-1]
}

// GetSub returns the Subdomain (Third Level Domain) of the given domain d.
// Returns nil if the subdomain not found.
func GetSub[T ByteSeq](d T) []byte {

	var b []byte

	switch v := any(d).(type) {
	case string:
		b = []byte(v)
	case []byte:
		b = v
	default:
		panic("Invalid type")
	}

	parts := bytes.Split(b, []byte{'.'})

	if len(parts) < 3 {
		return nil
	}

	// Remove the domain and the tld and pass to Join()
	return bytes.Join(parts[:len(parts)-2], []byte("."))
}

// GetDomain returns the domain of d (eg.: sub.example.com -> example.com).
// Returns nil if the domain was nor found.
func GetDomain[T ByteSeq](d T) []byte {

	var b []byte

	switch v := any(d).(type) {
	case string:
		b = []byte(v)
	case []byte:
		b = v
	default:
		panic("Invalid type")
	}

	parts := bytes.Split(b, []byte{'.'})

	l := len(parts)

	switch {
	case l < 2:
		return nil
	case len(parts[l-1]) == 0:
		return nil
	case len(parts[l-2]) == 0:
		return nil

	}

	return bytes.Join(parts[len(parts)-2:], []byte{'.'})
}

// IsWildcard returns whether the given domain d is a wildcard domain.
func IsWildcard[T ByteSeq](d T) bool {

	s := GetSub(d)

	return bytes.Equal(s, []byte{'*'})
}
