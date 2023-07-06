package domain

import (
	"fmt"

	"github.com/miekg/dns"
)

var TypeNS uint16 = 2

// QueryNS returns the answer as a string slice.
// The length of the returned slice can be 0 if no record matching for type MX, but record with other type exist.
// Returns nil in case of error.
func QueryNS(name string) ([]string, error) {

	var (
		a   []dns.RR
		r   = make([]string, 0)
		err error
	)

	a, err = Query(name, TypeNS)
	if err != nil {
		return r, err
	}

	for i := range a {

		switch v := a[i].(type) {
		case *dns.NS:
			r = append(r, v.Ns)
		default:
			return r, fmt.Errorf("unknown type: %T", v)
		}
	}

	return r, err
}

// IsSetNS checks whether an NS type record set for name.
// NXDOMAIN is not an error here, because it means "not found".
func IsSetNS(name string) (bool, error) {
	return IsSet(name, TypeNS)
}
