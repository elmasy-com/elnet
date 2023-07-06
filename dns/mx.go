package dns

import (
	"fmt"

	"github.com/miekg/dns"
)

var TypeMX uint16 = 15

type MX struct {
	Preference int    // Priority
	Exchange   string // Server's hostname
}

// QueryMX returns a slice of MX struct.
// Returns nil in case of error.
func QueryMX(name string) ([]MX, error) {

	a, err := Query(name, TypeMX)
	if err != nil {
		return nil, err
	}

	r := make([]MX, 0)

	for i := range a {

		switch v := a[i].(type) {
		case *dns.MX:
			r = append(r, MX{Preference: int(v.Preference), Exchange: v.Mx})
		default:
			return nil, fmt.Errorf("unknown type: %T", v)
		}
	}

	return r, nil
}

// QueryMXRetry query for MX record and retry for n times if an error occured.
func QueryMXRetry(name string, n int) ([]MX, error) {

	if n < 1 {
		return nil, fmt.Errorf("invalid number of retry: %d", n)
	}

	var (
		r   []MX = nil
		err error
	)

	for i := 0; i < n; i++ {

		r, err = QueryMX(name)
		if err == nil {
			return r, nil
		}
	}

	return nil, err
}

// IsSetMX checks whether an MX type record set for name.
// NXDOMAIN is not an error here, because it means "not found".
func IsSetMX(name string) (bool, error) {
	return IsSet(name, TypeMX)
}
