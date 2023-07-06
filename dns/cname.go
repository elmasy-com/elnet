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
		default:
			return nil, fmt.Errorf("unknown type: %T", v)
		}
	}

	return r, nil
}

// QueryCNAMERetry query for CNAME record and retry for n times if an error occured.
func QueryCNAMERetry(name string, n int) ([]string, error) {

	if n < 1 {
		return nil, fmt.Errorf("invalid number of retry: %d", n)
	}

	var (
		r   []string = nil
		err error
	)

	for i := 0; i < n; i++ {

		r, err = QueryCNAME(name)
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
