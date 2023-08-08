package validator

// Domain returns whether d is valid domain.
//
// This function returns false for "." (root domain).
func Domain(d string) bool {

	// Domain checks if a ByteSeq is a presentation-format domain name
	// (currently restricted to hostname-compatible "preferred name" LDH labels and
	// SRV-like "underscore labels".

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

// DomainPart returns whether d is valid domain part (eg.: subdomain part or Second Level Domain).
//
// Domain d can be up to 63 character long, can include a-z, A-Z, 0-9 and "-".
// The string must not starts and ends with a hyphen ("-") and two consecutive hyphen is not allowed.
func DomainPart(d string) bool {

	// Source: https://www.saveonhosting.com/scripts/index.php?rp=/knowledgebase/52/What-are-the-valid-characters-for-a-domain-name-and-how-long-can-a-domain-name-be.html

	l := len(d)

	switch {
	case l == 0 || l > 63:
		// See RFC 1035, RFC 3696.
		// Presentation format has dots before every label except the first, and the
		// terminal empty label is optional here because we assume fully-qualified
		// (absolute) input. We must therefore reserve space for the first and last
		// labels' length octets in wire format, where they are necessary and the
		// maximum total length is 255.
		// So our _effective_ maximum is 253, but 254 is not rejected if the last
		// character is a dot.
		return false
	case d[0] == '-' || d[l-1] == '-':
		// Cant starts and ends with "-"
		return false
	case d == ".":
		// The root domain name is technically valid. See golang.org/issue/45715.
		// But not for this package
		return false
	}

	// Indicates, that the last character was a hyphen
	lastHypen := false

	for i := 0; i < len(d); i++ {
		c := d[i]
		switch {
		default:
			return false
		case 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z':
			continue
		case '0' <= c && c <= '9':
			// fine
			continue
		case c == '-':
			// Two consecutive hyphen is not allowed (eg.: "a--a")
			if lastHypen {
				return false
			}

			lastHypen = true
		}
	}

	return true
}
