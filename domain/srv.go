package domain

import (
	"fmt"

	"github.com/miekg/dns"
)

type SRV struct {
	Priority int
	Weight   int
	Port     int
	Target   string
}

// QuerySRV returns the answer as a slice os SRV.
// The length of the returned slice can be 0 if no record matching for type SRV, but record with other type exist.
// Returns nil in case of error.
func QuerySRV(name string) ([]SRV, error) {

	var (
		a   []dns.RR
		r   = make([]SRV, 0)
		err error
	)

	a, err = Query(name, dns.TypeSRV)
	if err != nil {
		return r, err
	}

	for i := range a {

		switch v := a[i].(type) {
		case *dns.SRV:
			r = append(r, SRV{Priority: int(v.Priority), Weight: int(v.Weight), Port: int(v.Port), Target: v.Target})
		default:
			return r, fmt.Errorf("unknown type: %T", v)
		}
	}

	return r, err
}

// IsSetSRV checks whether an SRV type record set for name.
// NXDOMAIN is not an error here, because it means "not found".
func IsSetSRV(name string) (bool, error) {
	return IsSet(name, dns.TypeSRV)
}