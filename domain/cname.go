package domain

import (
	"fmt"

	"github.com/miekg/dns"
)

// QueryCNAME returns the answer as a string slice.
// The length of the returned slice can be 0 if no record matching for type CNAME, but record with other type exist.
// Returns nil in case of error.
func QueryCNAME(name string) ([]string, error) {

	var (
		a   []dns.RR
		r   = make([]string, 0)
		err error
	)

	a, err = Query(name, dns.TypeCNAME)
	if err != nil {
		return r, err
	}

	for i := range a {

		switch v := a[i].(type) {
		case *dns.CNAME:
			r = append(r, v.Target)
		default:
			return r, fmt.Errorf("unknown type: %T", v)
		}
	}

	return r, err
}

// IsSetCNAME checks whether an CNAME type record set for name.
// NXDOMAIN is not an error here, because it means "not found".
func IsSetCNAME(name string) (bool, error) {
	return IsSet(name, dns.TypeCNAME)
}
