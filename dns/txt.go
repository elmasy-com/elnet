package domain

import (
	"fmt"

	"github.com/miekg/dns"
)

var TypeTXT uint16 = 16

// QueryTXT returns the answer as a string slice.
// The length of the returned slice can be 0 if no record matching for type TXT, but record with other type exist.
// Returns nil in case of error.
func QueryTXT(name string) ([]string, error) {

	var (
		a   []dns.RR
		r   = make([]string, 0)
		err error
	)

	a, err = Query(name, TypeTXT)
	if err != nil {
		return r, err
	}

	for i := range a {

		switch v := a[i].(type) {
		case *dns.TXT:
			r = append(r, v.Txt...)
		default:
			return r, fmt.Errorf("unknown type: %T", v)
		}
	}

	return r, err
}

// IsSetTXT checks whether an TXT type record set for name.
// NXDOMAIN is not an error here, because it means "not found".
func IsSetTXT(name string) (bool, error) {
	return IsSet(name, TypeTXT)
}
