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

func (m MX) String() string {
	return fmt.Sprintf("%d %s", m.Preference, m.Exchange)
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

// QueryMXServer returns a slice of MX struct. Use server s to query.
// Returns nil in case of error.
func QueryMXServer(name string, s string) ([]MX, error) {

	a, err := QueryServer(name, TypeMX, s)
	if err != nil {
		return nil, err
	}

	r := make([]MX, 0)

	for i := range a {

		switch v := a[i].(type) {
		case *dns.MX:
			r = append(r, MX{Preference: int(v.Preference), Exchange: v.Mx})
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

// QueryMXRetry query for MX record and retry for MaxRetries times if an error occured.
func QueryMXRetry(name string) ([]MX, error) {

	var (
		r   []MX  = nil
		err error = ErrInvalidMaxRetries
	)

	for i := 0; i < MaxRetries; i++ {

		r, err = QueryMXServer(name, getServer(i))
		if err == nil {
			return r, nil
		}
	}

	return nil, err
}

// QueryMXRetryStr query for MX record and retry for MaxRetries times if an error occured and returns the result as a string slice.
func QueryMXRetryStr(name string) ([]string, error) {

	r, err := QueryMXRetry(name)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0, len(r))

	for i := range r {
		result = append(result, r[i].String())
	}

	return result, nil
}

// IsSetMX checks whether an MX type record set for name.
// NXDOMAIN is not an error here, because it means "not found".
func IsSetMX(name string) (bool, error) {
	return IsSet(name, TypeMX)
}
