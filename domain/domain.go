package domain

import (
	"bytes"
	"errors"
	"fmt"
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

// GetDomain returns the domain (eg.: "example.com").
func (p *Parts) GetDomain() string {
	return p.Domain + "." + p.TLD
}

func (p Parts) String() string {

	v := ""

	if p.Sub != "" {
		v += p.Sub + "."
	}

	v += p.Domain + "." + p.TLD

	return v
}

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

// Clean removes the trailing dot
// returns a lower cased version of d.
func Clean(d string) string {

	// Remove the trailing dot.
	if d[len(d)-1] == '.' {
		d = d[:len(d)-1]
	}

	return strings.ToLower(d)
}

// Returns the dot indexes.
// If no dot found, returns nil
func getDotsIndex(d string) []int {

	var indexes []int = nil

	for i := 0; i < len(d); i++ {
		if d[i] == '.' {
			indexes = append(indexes, i)
		}
	}

	return indexes
}

// Returns the dot indexes from right to front.
// If no dot found, returns nil.
func getDotsIndexRev(d string) []int {

	var indexes []int = nil

	for i := len(d) - 1; i >= 0; i-- {
		if d[i] == '.' {
			indexes = append(indexes, i)
		}
	}

	return indexes
}

// getPartsIndex returns the dot index before the TLD and aftre the subdomain,
// allows an effective dissect of the FQDN.
// If d is a TLD, then domain will be -1.
func getPartsIndex(d string) (tld int, domain int) {

	// Default
	tld = 0
	domain = -1

	di := getDotsIndexRev(d)

	// Iterate in reverse order over the dot indexes
	for i := range di {

		// dot index + 1 to skip dot in the string

		tldStr, icann := publicsuffix.PublicSuffix(d[di[i]+1:])

		switch {

		case icann && bytes.Equal([]byte(tldStr), []byte(d[di[i]+1:])):
			// ICANN managed TLD found, continue to the next
			tld = di[i]
			if i == 0 {
				domain = 0
			}
		case icann && !bytes.Equal([]byte(tldStr), []byte(d[di[i]+1:])):
			// The returned TLD is ICANN managed and differs from d[di[i]+1:]
			// This means, that d[di[i]+1:] was a domain, and PublicSuffix() founded the TLD
			return di[i-1], di[i]

		case !icann && bytes.Equal([]byte(tldStr), []byte(d[di[i]+1:])) && i == 0:
			// Non existent TLD found in the first round (eg.: there.is.no.SUCH-TLD)

			if len(di) == 1 {
				// Only one dot present -> domain.INVALID-TLD
				//                        0     di[0]
				return di[0], 0
			}

			// Multiple dot presents -> sub.domain.INVALID-TLD
			//                           di[1]  di[0]
			return di[0], di[1]

		case !icann && bytes.Equal([]byte(tldStr), []byte(d[di[i]+1:])) && i > 0:

			// ICANN managed TLD with privately managed  domain
			//  -> example.debian.net
			//         di[i] di[i-1]
			return di[i-1], di[i]
		}

	}

	return tld, domain
}

// GetParts validate d with IsValid, Clean() and returns the parts of d.
// If d is invalid, returns ErrInvalidDomain.
// Will panic if failed to get parts after validation.
func GetParts(d string) (*Parts, error) {

	if !IsValid(d) {
		return nil, ErrInvalidDomain
	}

	d = Clean(d)

	tld, dom := getPartsIndex(d)
	if tld < 1 || dom < 0 {
		panic(fmt.Sprintf("getPartsIndex() failed after validation: %s", d))
	}

	p := new(Parts)

	p.TLD = d[tld+1:]

	if dom == 0 {
		p.Domain = d[0:tld]
	} else {
		p.Domain = d[dom+1 : tld]
		p.Sub = d[:dom]
	}

	return p, nil
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
