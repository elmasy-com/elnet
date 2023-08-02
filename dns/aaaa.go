package dns

import (
	"errors"
	"fmt"
	"net"

	"github.com/miekg/dns"
)

var TypeAAAA uint16 = 28

// QueryAAAA returns a slice of net.IP.
// Returns nil in case of error.
//
// The other record types are ignored.
func QueryAAAA(name string) ([]net.IP, error) {

	a, err := Query(name, TypeAAAA)
	if err != nil {
		return nil, err
	}

	r := make([]net.IP, 0)

	for i := range a {

		switch v := a[i].(type) {
		case *dns.AAAA:
			r = append(r, v.AAAA)
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

// QueryAAAAServer returns a slice of net.IP. Use server s to query.
// Returns nil in case of error.
//
// The other record types are ignored.
func QueryAAAAServer(name string, s string) ([]net.IP, error) {

	a, err := QueryServer(name, TypeAAAA, s)
	if err != nil {
		return nil, err
	}

	r := make([]net.IP, 0)

	for i := range a {

		switch v := a[i].(type) {
		case *dns.AAAA:
			r = append(r, v.AAAA)
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

// QueryAAAARetry query for AAAA record and retry for MaxRetries times if an error occured.
//
// NXDOMAIN is not count as an error!
func QueryAAAARetry(name string) ([]net.IP, error) {

	var (
		r   []net.IP = nil
		err error    = ErrInvalidMaxRetries
	)

	for i := -1; i < MaxRetries-1; i++ {

		r, err = QueryAAAAServer(name, getServer(i))
		if err == nil || errors.Is(err, ErrName) {
			return r, err
		}
	}

	return nil, err
}

// QueryAAAARetryStr query for AAAA record and retry for MaxRetries times if an error occured and returns the result as a string slice.
func QueryAAAARetryStr(name string) ([]string, error) {

	r, err := QueryAAAARetry(name)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0, len(r))

	for i := range r {
		result = append(result, r[i].String())
	}

	return result, nil
}

// IsSetAAAA checks whether an AAAA type record set for name.
// NXDOMAIN is not an error here, because it means "not found".
func IsSetAAAA(name string) (bool, error) {
	return IsSet(name, TypeAAAA)
}
