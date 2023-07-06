package dns

import (
	"fmt"

	"github.com/miekg/dns"
)

var TypeCNAME uint16 = 5

// QueryCNAME returns the answer as a string slice.
// The length of the returned slice can be 0 if no record matching for type CNAME, but record with other type exist.
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

// IsSetCNAME checks whether an CNAME type record set for name.
// NXDOMAIN is not an error here, because it means "not found".
func IsSetCNAME(name string) (bool, error) {
	return IsSet(name, TypeCNAME)
}
