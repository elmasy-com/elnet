package dns

import (
	"fmt"

	"github.com/miekg/dns"
)

var TypeCAA uint16 = 257

type CAA struct {
	Flag  uint8
	Tag   string
	Value string
}

func (c CAA) String() string {
	return fmt.Sprintf("%d %s %s", c.Flag, c.Tag, c.Value)
}

// QueryCAA returns a slice of net.IP.
// The answer slice will be nil in case of error.
func QueryCAA(name string) ([]CAA, error) {

	a, err := Query(name, TypeCAA)
	if err != nil {
		return nil, err
	}

	r := make([]CAA, 0)

	for i := range a {

		switch v := a[i].(type) {
		case *dns.CAA:
			r = append(r, CAA{Flag: v.Flag, Tag: v.Tag, Value: v.Value})
		case *dns.CNAME:
			// Ignore CNAME
			continue
		default:
			return nil, fmt.Errorf("unknown type: %T", v)
		}
	}

	return r, nil
}

// QueryCAARetry query for CAA record and retry for MaxRetries times if an error occured.
func QueryCAARetry(name string) ([]CAA, error) {

	var (
		r   []CAA = nil
		err error = ErrInvalidMaxRetries
	)

	for i := 0; i < MaxRetries; i++ {

		r, err = QueryCAA(name)
		if err == nil {
			return r, nil
		}
	}

	return nil, err
}

// QueryCAARetryStr query for CAA record and retry for MaxRetries times if an error occured and returns the result as a string slice.
func QueryCAARetryStr(name string) ([]string, error) {

	r, err := QueryCAARetry(name)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0, len(r))

	for i := range r {
		result = append(result, r[i].String())
	}

	return result, nil
}

// IsSetCAA checks whether an A type record set for name.
// NXDOMAIN is not an error here, because it means "not found".
func IsSetCAA(name string) (bool, error) {
	return IsSet(name, TypeCAA)
}
