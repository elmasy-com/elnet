package dns

import (
	"errors"
	"fmt"
	"net"

	"github.com/miekg/dns"
)

var TypeA uint16 = 1

// QueryA returns a slice of net.IP.
// The answer slice will be nil in case of error.
//
// The other record types are ignored.
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
		case *dns.DNAME:
			// Ignore DNAME
			continue
		default:
			return nil, fmt.Errorf("unknown type: %T", v)
		}
	}

	return r, nil
}

// QueryAServer returns a slice of net.IP. Use server s to query.
// The answer slice will be nil in case of error.
//
// The other record types are ignored.
func QueryAServer(name string, s string) ([]net.IP, error) {

	a, err := QueryServer(name, TypeA, s)
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
		case *dns.DNAME:
			// Ignore DNAME
			continue
		default:
			return nil, fmt.Errorf("unknown type: %T", v)
		}
	}

	return r, nil
}

// QueryARetry query for A record and retry for MaxRetries times if an error occured.
//
// NXDOMAIN is not count as an error!
func QueryARetry(name string) ([]net.IP, error) {

	var (
		r   []net.IP = nil
		err error    = ErrInvalidMaxRetries
	)

	for i := -1; i < MaxRetries-1; i++ {

		r, err = QueryAServer(name, getServer(i))
		if err == nil || errors.Is(err, ErrName) {
			return r, err
		}
	}

	return nil, err
}

// QueryARetryStr query for A record and retry for MaxRetries times if an error occured and returns the result as a string slice.
func QueryARetryStr(name string) ([]string, error) {

	r, err := QueryARetry(name)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0, len(r))

	for i := range r {
		result = append(result, r[i].String())
	}

	return result, nil
}

// IsSetA checks whether an A type record set for name.
// NXDOMAIN is not an error here, because it means "not found".
func IsSetA(name string) (bool, error) {
	return IsSet(name, TypeA)
}
