package dns

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

var TypeSRV uint16 = 33

// QuerySRV returns the answer as a slice os SRV.
// Returns nil in case of error.
func QuerySRV(name string) ([]SRV, error) {

	a, err := Query(name, TypeSRV)
	if err != nil {
		return nil, err
	}

	r := make([]SRV, 0)

	for i := range a {

		switch v := a[i].(type) {
		case *dns.SRV:
			r = append(r, SRV{Priority: int(v.Priority), Weight: int(v.Weight), Port: int(v.Port), Target: v.Target})
		default:
			return nil, fmt.Errorf("unknown type: %T", v)
		}
	}

	return r, err
}

// QuerySRVRetry query for SRV record and retry for MaxRetries times if an error occured.
func QuerySRVRetry(name string) ([]SRV, error) {

	var (
		r   []SRV = nil
		err error = ErrInvalidMaxRetries
	)

	for i := 0; i < MaxRetries; i++ {

		r, err = QuerySRV(name)
		if err == nil {
			return r, nil
		}
	}

	return nil, err
}

// IsSetSRV checks whether an SRV type record set for name.
// NXDOMAIN is not an error here, because it means "not found".
func IsSetSRV(name string) (bool, error) {
	return IsSet(name, TypeSRV)
}
