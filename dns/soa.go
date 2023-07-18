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

func (s SOA) String() string {
	return fmt.Sprintf("%s %s %d %d %d %d %d", s.Mname, s.Rname, s.Serial, s.Refresh, s.Retry, s.Expire, s.MinTTL)
}

// QuerySOA returns the answer as a SOA struct.
// The returned *SOA **can be nil** if error is nil.
// Returns nil in case of error.
func QuerySOA(name string) (*SOA, error) {

	a, err := Query(name, TypeSOA)
	if err != nil {
		return nil, err
	}

	if len(a) == 0 {
		return nil, nil
	}
	if len(a) > 1 {
		return nil, fmt.Errorf("invalid answer: %#v", a)
	}

	v, ok := a[0].(*dns.SOA)

	if !ok {
		return nil, fmt.Errorf("invalid answer: %#v", a)
	}

	return &SOA{Mname: v.Ns, Rname: v.Mbox, Serial: int(v.Serial), Refresh: int(v.Refresh), Retry: int(v.Retry), Expire: int(v.Expire), MinTTL: int(v.Minttl)}, err
}

// QuerySOAServer returns the answer as a SOA struct.
// The returned *SOA **can be nil** if error is nil.
// Returns nil in case of error.
func QuerySOAServer(name string, s string) (*SOA, error) {

	a, err := QueryServer(name, TypeSOA, s)
	if err != nil {
		return nil, err
	}

	if len(a) == 0 {
		return nil, nil
	}
	if len(a) > 1 {
		return nil, fmt.Errorf("invalid answer: %#v", a)
	}

	v, ok := a[0].(*dns.SOA)

	if !ok {
		return nil, fmt.Errorf("invalid answer: %#v", a)
	}

	return &SOA{Mname: v.Ns, Rname: v.Mbox, Serial: int(v.Serial), Refresh: int(v.Refresh), Retry: int(v.Retry), Expire: int(v.Expire), MinTTL: int(v.Minttl)}, err
}

// QuerySOARetry query for SOA record and retry for MaxRetries times if an error occured.
func QuerySOARetry(name string) (*SOA, error) {

	var (
		r   *SOA  = nil
		err error = ErrInvalidMaxRetries
	)

	for i := 0; i < MaxRetries; i++ {

		r, err = QuerySOAServer(name, getServer(i))
		if err == nil {
			return r, nil
		}
	}

	return nil, err
}

// QuerySOARetry query for SOA record and retry for MaxRetries times if an error occured.
func QuerySOARetryStr(name string) (string, error) {

	r, err := QuerySOARetry(name)
	if err != nil {
		return "", err
	}

	if r == nil {
		return "", nil
	}

	return r.String(), nil
}

// TODO: Decide if the domain is registered based on the SOA record/root server
