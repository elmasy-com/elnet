package dns

import (
	"fmt"

	"github.com/miekg/dns"
)

var TypeTXT uint16 = 16

// QueryTXT returns the answer as a string slice.
// Returns nil in case of error.
func QueryTXT(name string) ([]string, error) {

	a, err := Query(name, TypeTXT)
	if err != nil {
		return nil, err
	}

	r := make([]string, 0)

	for i := range a {

		switch v := a[i].(type) {
		case *dns.TXT:
			r = append(r, v.Txt...)
		default:
			return nil, fmt.Errorf("unknown type: %T", v)
		}
	}

	return r, nil
}

// QueryTXTRetry query for TXT record and retry for n times if an error occured.
func QueryTXTRetry(name string, n int) ([]string, error) {

	if n < 1 {
		return nil, fmt.Errorf("invalid number of retry: %d", n)
	}

	var (
		r   []string = nil
		err error
	)

	for i := 0; i < n; i++ {

		r, err = QueryTXT(name)
		if err == nil {
			return r, nil
		}
	}

	return nil, err
}

// IsSetTXT checks whether an TXT type record set for name.
// NXDOMAIN is not an error here, because it means "not found".
func IsSetTXT(name string) (bool, error) {
	return IsSet(name, TypeTXT)
}
