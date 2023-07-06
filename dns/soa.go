package dns

import (
	"fmt"

	"github.com/miekg/dns"
)

var TypeSOA uint16 = 6

// See more: https://www.rfc-editor.org/rfc/rfc1035.html#section-3.3.13
type SOA struct {
	Mname   string
	Rname   string
	Serial  int
	Refresh int
	Retry   int
	Expire  int
	MinTTL  int
}

// QuerySOA returns the answer as a SOA struct.
// The returned *SOA **cant be nil** if error is nil.
// Returns nil in case of error.
func QuerySOA(name string) (*SOA, error) {

	a, err := Query(name, TypeSOA)
	if err != nil {
		return nil, err
	}

	if len(a) != 1 {
		return nil, fmt.Errorf("invalid answer: %#v", a)
	}

	v, ok := a[0].(*dns.SOA)

	if !ok {
		return nil, fmt.Errorf("invalid answer: %#v", a)
	}

	return &SOA{Mname: v.Ns, Rname: v.Mbox, Serial: int(v.Serial), Refresh: int(v.Refresh), Retry: int(v.Retry), Expire: int(v.Expire), MinTTL: int(v.Minttl)}, err
}

// QuerySOARetry query for SOA record and retry for n times if an error occured.
func QuerySOARetry(name string, n int) (*SOA, error) {

	if n < 1 {
		return nil, fmt.Errorf("invalid number of retry: %d", n)
	}

	var (
		r   *SOA = nil
		err error
	)

	for i := 0; i < n; i++ {

		r, err = QuerySOA(name)
		if err == nil {
			return r, nil
		}
	}

	return nil, err
}

// TODO: Decide if the domain is registered based on the SOA record/root server
