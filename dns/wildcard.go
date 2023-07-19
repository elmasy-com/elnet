package dns

import (
	"strings"

	"github.com/g0rbe/slitu"
)

var charSet = []byte("abcdefghijklmnopqrstuvwxyz0123456789")

// This is a special case, where the total length of the domain is 253 and the size of first part is only one char (eg.: "a.a...").
// There is no room to fuzz the first part, have to check every possible characters to make sure its a wildcard domain.
func wildcardBruteforceOneChar(parts []string, t uint16) (bool, error) {

	for i := range charSet {

		parts[0] = string(charSet[i])
		v := strings.Join(parts, ".")

		r, err := IsSet(v, t)

		if err != nil {
			return false, err
		}

		// Return on the first false
		if !r {
			return false, nil
		}
	}

	return true, nil
}

// IsWildcard check if name is a wildcard domain.
//
// NOTE: Use IsValid() and Clean() before this function!
func IsWildcard(name string, t uint16) (bool, error) {

	if !HasSub(name) {
		// Domain without subdomain cant be a wildcard
		return false, nil
	}

	parts := strings.Split(name, ".")

	// partSize is the possible max size of first part to fuzz
	partSize := 253 - len(name) + len(parts[0])

	if partSize == 1 {
		return wildcardBruteforceOneChar(parts, t)
	}

	// Limit the part size to 63
	if partSize > 63 {
		partSize = 63
	}

	maxCheck := 0

	// The total number of checks based on the max length.
	switch {
	case partSize > 31:
		maxCheck = 3
	case partSize > 15:
		maxCheck = 5
	case partSize > 8:
		maxCheck = 10
	default:
		maxCheck = 15
	}

	for i := 0; i < maxCheck; i++ {

		parts[0] = slitu.RandomString(charSet, partSize)
		v := strings.Join(parts, ".")

		r, err := IsSet(v, t)
		if err != nil {
			return false, err
		}

		// Return on the first false
		if !r {
			return false, nil
		}
	}

	return true, nil
}
