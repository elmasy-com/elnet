package dns

import (
	"fmt"

	mdns "github.com/miekg/dns"
)

type Record struct {
	Type  uint16
	Value string
}

// QueryAll query every known type and returns the records.
// This function checks whether name with the type is a wildcard, and if name is a wildcard, ommit from the retuned []Record.
func (s *Servers) QueryAll(name string) ([]Record, []error) {

	var (
		errs = make([]error, 0)
		rr   = make([]mdns.RR, 0)
	)

	/*
	 * A
	 */

	r, err := s.TryQuery(name, TypeA)
	if err != nil {
		errs = append(errs, fmt.Errorf("query A: %w", err))
	}

	// Checks whether name is a wildcard
	wc, err := s.IsWildcard(name, TypeA)
	if err != nil {
		errs = append(errs, fmt.Errorf("wildcard A: %w", err))
	}

	// If not a wildcard domain, append the result
	if !wc {
		rr = append(rr, r...)
	}

	/*
	 * AAAA
	 */

	r, err = s.TryQuery(name, TypeAAAA)
	if err != nil {
		errs = append(errs, fmt.Errorf("query AAAA: %w", err))
	}

	// Checks whether name is a wildcard
	wc, err = s.IsWildcard(name, TypeAAAA)
	if err != nil {
		errs = append(errs, fmt.Errorf("wildcard AAAA: %w", err))
	}

	// If not a wildcard domain, append the result
	if !wc {
		rr = append(rr, r...)
	}

	/*
	 * CAA
	 */

	r, err = s.TryQuery(name, TypeCAA)
	if err != nil {
		errs = append(errs, fmt.Errorf("query CAA: %w", err))
	}

	// Checks whether name is a wildcard
	wc, err = s.IsWildcard(name, TypeCAA)
	if err != nil {
		errs = append(errs, fmt.Errorf("wildcard CAA: %w", err))
	}

	// If not a wildcard domain, append the result
	if !wc {
		rr = append(rr, r...)
	}

	/*
	 * CNAME
	 */

	r, err = s.TryQuery(name, TypeCNAME)
	if err != nil {
		errs = append(errs, fmt.Errorf("query CNAME: %w", err))
	}

	// Checks whether name is a wildcard
	wc, err = s.IsWildcard(name, TypeCNAME)
	if err != nil {
		errs = append(errs, fmt.Errorf("wildcard CNAME: %w", err))
	}

	// If not a wildcard domain, append the result
	if !wc {
		rr = append(rr, r...)
	}

	/*
	 * DNAME
	 */

	r, err = s.TryQuery(name, TypeDNAME)
	if err != nil {
		errs = append(errs, fmt.Errorf("query DNAME: %w", err))
	}

	// Checks whether name is a wildcard
	wc, err = s.IsWildcard(name, TypeDNAME)
	if err != nil {
		errs = append(errs, fmt.Errorf("wildcard DNAME: %w", err))
	}

	// If not a wildcard domain, append the result
	if !wc {
		rr = append(rr, r...)
	}

	/*
	 * MX
	 */
	r, err = s.TryQuery(name, TypeMX)
	if err != nil {
		errs = append(errs, fmt.Errorf("query MX: %w", err))
	}

	// Checks whether name is a wildcard
	wc, err = s.IsWildcard(name, TypeMX)
	if err != nil {
		errs = append(errs, fmt.Errorf("wildcard MX: %w", err))
	}

	// If not a wildcard domain, append the result
	if !wc {
		rr = append(rr, r...)
	}

	/*
	 * NS
	 */

	r, err = s.TryQuery(name, TypeNS)
	if err != nil {
		errs = append(errs, fmt.Errorf("query NS: %w", err))
	}

	// Checks whether name is a wildcard
	wc, err = s.IsWildcard(name, TypeNS)
	if err != nil {
		errs = append(errs, fmt.Errorf("wildcard NS: %w", err))
	}

	// If not a wildcard domain, append the result
	if !wc {
		rr = append(rr, r...)
	}

	/*
	 * SOA
	 */

	r, err = s.TryQuery(name, TypeSOA)
	if err != nil {
		errs = append(errs, fmt.Errorf("query SOA: %w", err))
	}

	// Checks whether name is a wildcard
	wc, err = s.IsWildcard(name, TypeSOA)
	if err != nil {
		errs = append(errs, fmt.Errorf("wildcard SOA: %w", err))
	}

	// If not a wildcard domain, append the result
	if !wc {
		rr = append(rr, r...)
	}

	/*
	 * SRV
	 */

	r, err = s.TryQuery(name, TypeSRV)
	if err != nil {
		errs = append(errs, fmt.Errorf("query SRV: %w", err))
	}

	// Checks whether name is a wildcard
	wc, err = s.IsWildcard(name, TypeTXT)
	if err != nil {
		errs = append(errs, fmt.Errorf("wildcard TXT: %w", err))
	}

	// If not a wildcard domain, append the result
	if !wc {
		rr = append(rr, r...)
	}

	/*
	 * TXT
	 */

	r, err = s.TryQuery(name, TypeTXT)
	if err != nil {
		errs = append(errs, fmt.Errorf("query TXT: %w", err))
	}

	// Checks whether name is a wildcard
	wc, err = s.IsWildcard(name, TypeTXT)
	if err != nil {
		errs = append(errs, fmt.Errorf("wildcard TXT: %w", err))
	}

	// If not a wildcard domain, append the result
	if !wc {
		rr = append(rr, r...)
	}

	/*
	 * Read the answers
	 */
	rs := make([]Record, 0, len(rr))

	for i := range rr {

		switch v := rr[i].(type) {
		case *mdns.A:
			rs = append(rs, Record{Type: TypeA, Value: v.A.String()})
		case *mdns.AAAA:
			rs = append(rs, Record{Type: TypeAAAA, Value: v.AAAA.String()})
		case *mdns.CAA:
			rs = append(rs, Record{Type: TypeCAA, Value: fmt.Sprintf("%d %s %s", v.Flag, v.Tag, v.Value)})
		case *mdns.CNAME:
			rs = append(rs, Record{Type: TypeCNAME, Value: v.Target})
		case *mdns.DNAME:
			rs = append(rs, Record{Type: TypeDNAME, Value: v.Target})
		case *mdns.MX:
			rs = append(rs, Record{Type: TypeMX, Value: fmt.Sprintf("%d %s", v.Preference, v.Mx)})
		case *mdns.NS:
			rs = append(rs, Record{Type: TypeNS, Value: v.Ns})
		case *mdns.SOA:
			rs = append(rs, Record{Type: TypeSOA, Value: fmt.Sprintf("%s %s %d %d %d %d %d", v.Ns, v.Mbox, v.Serial, v.Refresh, v.Retry, v.Expire, v.Minttl)})
		case *mdns.SRV:
			rs = append(rs, Record{Type: TypeSRV, Value: fmt.Sprintf("%d %d %d %s", v.Priority, v.Weight, v.Port, v.Target)})
		case *mdns.TXT:
			for ii := range v.Txt {
				rs = append(rs, Record{Type: TypeTXT, Value: v.Txt[ii]})
			}
		default:
			errs = append(errs, fmt.Errorf("unknown type: %T", v))
		}
	}

	return rs, errs
}

// QueryAll query every known type and returns the records.
// This function checks whether name with the type is a wildcard.
func QueryAll(name string) ([]Record, []error) {

	return DefaultServers.QueryAll(name)
}
