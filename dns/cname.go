package dns

import (
	"fmt"

	"github.com/miekg/dns"
)

var TypeCNAME uint16 = 5

// QueryCNAME returns the answer as a string slice.
// Returns nil in case of error.
func QueryCNAME(name string) ([]string, error) {

	a, err := Query(name, TypeCNAME)
	if err != nil {
		return nil, err
	}

	r := make([]string, 0)

	for i := range a {

		switch v := a[i].(type) {
		case *dns.CNAME:
			r = append(r, v.Target)
		case *dns.DNAME:
			// Ignore DNAME
			continue
		default:
			return nil, fmt.Errorf("unknown type: %T", v)
		}
	}

	return r, nil
}

// QueryCNAMEServer returns the answer as a string slice. Use server s to query.
// Returns nil in case of error.
func QueryCNAMEServer(name string, s string) ([]string, error) {

	a, err := QueryServer(name, TypeCNAME, s)
	if err != nil {
		return nil, err
	}

	r := make([]string, 0)

	for i := range a {

		switch v := a[i].(type) {
		case *dns.CNAME:
			r = append(r, v.Target)
		case *dns.DNAME:
			// Ignore DNAME
			continue
		default:
			return nil, fmt.Errorf("unknown type: %T", v)
		}
	}

	return r, nil
}

// QueryCNAMERetry query for CNAME record and retry for MaxRetries times if an error occured.
func QueryCNAMERetry(name string) ([]string, error) {

	var (
		r   []string = nil
		err error    = ErrInvalidMaxRetries
	)

	for i := 0; i < MaxRetries; i++ {

		r, err = QueryCNAMEServer(name, getServer(i))
		if err == nil {
			return r, nil
		}
	}

	return nil, err
}

// IsSetCNAME checks whether an CNAME type record set for name.
// NXDOMAIN is not an error here, because it means "not found".
func IsSetCNAME(name string) (bool, error) {
	return IsSet(name, TypeCNAME)
}
