package dns

import (
	"errors"
	"fmt"

	"github.com/miekg/dns"
)

var TypeDNAME uint16 = 39

// QueryDNAME returns the target string.
//
// The other records are ignored.
func QueryDNAME(name string) (string, error) {

	a, err := Query(name, TypeDNAME)
	if err != nil {
		return "", err
	}

	for i := range a {

		switch v := a[i].(type) {
		case *dns.DNAME:
			return v.Target, nil
		case *dns.CNAME:
			// Ignore CNAME
			continue
		default:
			return "", fmt.Errorf("unknown type: %T", v)
		}
	}

	return "", nil
}

// QueryDNAMEServer returns the target string. USe server s to query.
//
// The other records are ignored.
func QueryDNAMEServer(name string, s string) (string, error) {

	a, err := QueryServer(name, TypeDNAME, s)
	if err != nil {
		return "", err
	}

	for i := range a {

		switch v := a[i].(type) {
		case *dns.DNAME:
			return v.Target, nil
		case *dns.CNAME:
			// Ignore CNAME
			continue
		default:
			return "", fmt.Errorf("unknown type: %T", v)
		}
	}

	return "", nil
}

// QueryDNAMERetry query for DNAME record and retry for MaxRetries times if an error occured.
//
// NXDOMAIN is not count as an error!
func QueryDNAMERetry(name string) (string, error) {

	var (
		r   string
		err error = ErrInvalidMaxRetries
	)

	for i := -1; i < MaxRetries-1; i++ {

		r, err = QueryDNAMEServer(name, getServer(i))
		if err == nil || errors.Is(err, ErrName) {
			return r, err
		}
	}

	return "", err
}

// IsSetA checks whether an A type record set for name.
// NXDOMAIN is not an error here, because it means "not found".
func IsSetDNAME(name string) (bool, error) {
	return IsSet(name, TypeDNAME)
}
