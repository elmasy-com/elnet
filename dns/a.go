package dns

import (
	"fmt"
	"net"

	"github.com/miekg/dns"
)

var TypeA uint16 = 1

// QueryA returns a slice of net.IP.
// The answer slice will be nil in case of error.
//
// The CNAME record is ignored.
func QueryA(name string) ([]net.IP, error) {

	a, err := Query(name, TypeA)
	if err != nil {
		return nil, err
	}

	r := make([]net.IP, 0)

	for i := range a {

		switch v := a[i].(type) {
		case *dns.A:
			r = append(r, v.A)
		case *dns.CNAME:
			// Ignore CNAME
			continue
		default:
			return nil, fmt.Errorf("unknown type: %T", v)
		}
	}

	return r, nil
}

// QueryARetry query for A record and retry for n times if an error occured.
func QueryARetry(name string, n int) ([]net.IP, error) {

	if n < 1 {
		return nil, fmt.Errorf("invalid number of retry: %d", n)
	}

	var (
		r   []net.IP = nil
		err error
	)

	for i := 0; i < n; i++ {

		r, err = QueryA(name)
		if err == nil {
			return r, nil
		}
	}

	return nil, err
}

// IsSetA checks whether an A type record set for name.
// NXDOMAIN is not an error here, because it means "not found".
func IsSetA(name string) (bool, error) {
	return IsSet(name, TypeA)
}
