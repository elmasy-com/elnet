package dns

import (
	"errors"
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
// The returned *SOA **can be nil**.
// Returns nil in case of error.
func QuerySOA(name string) (*SOA, error) {

	a, err := Query(name, TypeSOA)
	if err != nil {
		return nil, err
	}

	for i := range a {

		switch v := a[i].(type) {
		case *dns.SOA:
			return &SOA{Mname: v.Ns, Rname: v.Mbox, Serial: int(v.Serial), Refresh: int(v.Refresh), Retry: int(v.Retry), Expire: int(v.Expire), MinTTL: int(v.Minttl)}, nil
		case *dns.CNAME:
			// Ignore CNAME
			continue
		case *dns.DNAME:
			// Ignore DNAME
			continue
		default:
			return nil, fmt.Errorf("unknown type: %T", v)
		}
	}

	return nil, nil
}

// QuerySOAServer returns the answer as a SOA struct.
// The returned *SOA **can be nil**..
// Returns nil in case of error.
func QuerySOAServer(name string, s string) (*SOA, error) {

	a, err := QueryServer(name, TypeSOA, s)
	if err != nil {
		return nil, err
	}

	for i := range a {

		switch v := a[i].(type) {
		case *dns.SOA:
			return &SOA{Mname: v.Ns, Rname: v.Mbox, Serial: int(v.Serial), Refresh: int(v.Refresh), Retry: int(v.Retry), Expire: int(v.Expire), MinTTL: int(v.Minttl)}, nil
		case *dns.CNAME:
			// Ignore CNAME
			continue
		case *dns.DNAME:
			// Ignore DNAME
			continue
		default:
			return nil, fmt.Errorf("unknown type: %T", v)
		}
	}

	return nil, nil
}

// QuerySOARetry query for SOA record and retry for MaxRetries times if an error occured.
//
// NXDOMAIN is not count as an error!
func QuerySOARetry(name string) (*SOA, error) {

	var (
		r   *SOA  = nil
		err error = ErrInvalidMaxRetries
	)

	for i := -1; i < MaxRetries-1; i++ {

		r, err = QuerySOAServer(name, getServer(i))
		if err == nil || errors.Is(err, ErrName) {
			return r, err
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
