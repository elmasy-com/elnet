package dns

import (
	"errors"
	"fmt"

	"github.com/miekg/dns"
)

var TypeNS uint16 = 2

// QueryNS returns the answer as a string slice.
// Returns nil in case of error.
func QueryNS(name string) ([]string, error) {

	a, err := Query(name, TypeNS)
	if err != nil {
		return nil, err
	}

	r := make([]string, 0)

	for i := range a {

		switch v := a[i].(type) {
		case *dns.NS:
			r = append(r, v.Ns)
		case *dns.CNAME:
			// Ignore CNAME
			continue
		case *dns.DNAME:
			// Ignore DNAME
			continue
		default:
			return nil, fmt.Errorf("unknown type: %T", v)
		}
	}

	return r, nil
}

// QueryNSServer returns the answer as a string slice. Use server s to query.
// Returns nil in case of error.
func QueryNSServer(name string, s string) ([]string, error) {

	a, err := QueryServer(name, TypeNS, s)
	if err != nil {
		return nil, err
	}

	r := make([]string, 0)

	for i := range a {

		switch v := a[i].(type) {
		case *dns.NS:
			r = append(r, v.Ns)
		case *dns.CNAME:
			// Ignore CNAME
			continue
		case *dns.DNAME:
			// Ignore DNAME
			continue
		default:
			return nil, fmt.Errorf("unknown type: %T", v)
		}
	}

	return r, nil
}

// QueryNSRetry query for NS record and retry for MaxRetries times if an error occured.
//
// NXDOMAIN is not count as an error!
func QueryNSRetry(name string) ([]string, error) {

	var (
		r   []string = nil
		err error    = ErrInvalidMaxRetries
	)

	for i := -1; i < MaxRetries-1; i++ {

		r, err = QueryNSServer(name, GetServer(i))
		if err == nil || errors.Is(err, ErrName) {
			return r, err
		}
	}

	return nil, err
}

// IsSetNS checks whether an NS type record set for name.
// NXDOMAIN is not an error here, because it means "not found".
func IsSetNS(name string) (bool, error) {
	return IsSet(name, TypeNS)
}
