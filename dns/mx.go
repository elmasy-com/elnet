package domain

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
// The length of the returned slice can be 0 if no record matching for type MX, but record with other type exist.
// Returns nil in case of error.
func QueryMX(name string) ([]MX, error) {

	var (
		a   []dns.RR
		r   = make([]MX, 0)
		err error
	)

	a, err = Query(name, TypeMX)
	if err != nil {
		return r, err
	}

	for i := range a {

		switch v := a[i].(type) {
		case *dns.MX:
			r = append(r, MX{Preference: int(v.Preference), Exchange: v.Mx})
		default:
			return r, fmt.Errorf("unknown type: %T", v)
		}
	}

	return r, err
}

// IsSetMX checks whether an MX type record set for name.
// NXDOMAIN is not an error here, because it means "not found".
func IsSetMX(name string) (bool, error) {
	return IsSet(name, TypeMX)
}
