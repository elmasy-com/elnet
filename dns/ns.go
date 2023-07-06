package dns

import (
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
		default:
			return nil, fmt.Errorf("unknown type: %T", v)
		}
	}

	return r, nil
}

// QueryNSRetry query for NS record and retry for n times if an error occured.
func QueryNSRetry(name string, n int) ([]string, error) {

	if n < 1 {
		return nil, fmt.Errorf("invalid number of retry: %d", n)
	}

	var (
		r   []string = nil
		err error
	)

	for i := 0; i < n; i++ {

		r, err = QueryNS(name)
		if err == nil {
			return r, nil
		}
	}

	return nil, err
}

// IsSetNS checks whether an NS type record set for name.
// NXDOMAIN is not an error here, because it means "not found".
func IsSetNS(name string) (bool, error) {
	return IsSet(name, TypeNS)
}
