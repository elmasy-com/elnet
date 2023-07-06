package dns

import (
	"fmt"
	"net"

	"github.com/miekg/dns"
)

var TypeAAAA uint16 = 28

// QueryAAAA returns a slice of net.IP.
// Returns nil in case of error.
//
// The CNAME record is ignored.
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
		default:
			return nil, fmt.Errorf("unknown type: %T", v)
		}
	}

	return r, nil
}

// QueryAAAARetry query for AAAA record and retry for MaxRetries times if an error occured.
func QueryAAAARetry(name string) ([]net.IP, error) {

	var (
		r   []net.IP = nil
		err error    = ErrInvalidMaxRetries
	)

	for i := 0; i < MaxRetries; i++ {

		r, err = QueryAAAA(name)
		if err == nil {
			return r, nil
		}
	}

	return nil, err
}

// IsSetAAAA checks whether an AAAA type record set for name.
// NXDOMAIN is not an error here, because it means "not found".
func IsSetAAAA(name string) (bool, error) {
	return IsSet(name, TypeAAAA)
}
