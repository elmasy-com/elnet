package dns

import (
	"fmt"
	"net"

	"github.com/miekg/dns"
)

var TypeA uint16 = 1

// QueryA returns a slice of net.IP.
// The length of the returned slice can be 0 if no record matching for type A, but record with other type exist.
// The answer slice will be nil in case of error.
func QueryA(name string) ([]net.IP, error) {

	var (
		a   []dns.RR
		r   = make([]net.IP, 0)
		err error
	)

	a, err = Query(name, TypeA)
	if err != nil {
		return r, err
	}

	for i := range a {

		switch v := a[i].(type) {
		case *dns.A:
			r = append(r, v.A)
		case *dns.CNAME:
			// Ignore CNAME
			continue
		default:
			return r, fmt.Errorf("unknown type: %T", v)
		}
	}

	return r, err
}

// IsSetA checks whether an A type record set for name.
// NXDOMAIN is not an error here, because it means "not found".
func IsSetA(name string) (bool, error) {
	return IsSet(name, TypeA)
}
