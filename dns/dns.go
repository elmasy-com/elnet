package dns

import (
	"fmt"
	"math/rand"

	"github.com/elmasy-com/identify"
	"github.com/miekg/dns"
)

// Default client configuration.
//
// Set in the init() function to read resolv.conf or use Cloudflare + Google.
var Conf *dns.ClientConfig

// Initialize Conf.
// If reading resolv.conf failed, use Cloudflare + Google
func init() {

	var err error

	Conf, err = dns.ClientConfigFromFile("/etc/resolv.conf")
	if err == nil {
		return
	}

	Conf = new(dns.ClientConfig)

	Conf.Servers = []string{"1.1.1.1", "1.0.0.1", "8.8.8.8", "8.8.4.4"}
	Conf.Search = make([]string, 0)
	Conf.Port = "53"
	Conf.Ndots = 1
	Conf.Timeout = 5
	Conf.Attempts = 2
}

// getServer used to randomize DNS servers.
func getServer() string {

	var r string

	switch len(Conf.Servers) {
	case 0:
		r = "1.1.1.1"
	case 1:
		r = Conf.Servers[0]
	default:
		r = Conf.Servers[rand.Intn(len(Conf.Servers))]
	}

	if identify.IsValidIPv6(r) {
		r = "[" + r + "]"
	}

	return r + ":53"
}

// Generic query for type t.
// Returns the Answer section.
// In case of error, the answer will be nil and return ErrX or any unknown error.
func Query(name string, t uint16) ([]dns.RR, error) {

	name = dns.Fqdn(name)

	msg := new(dns.Msg)
	msg.SetQuestion(name, t)

	in, err := dns.Exchange(msg, getServer())
	if err != nil {
		return nil, err
	}

	if in.Rcode == 0 {
		return in.Answer, nil
	}

	return nil, RcodeToError(in.Rcode)
}

// IsSet checks whether a record with type t is set for name.
// NXDOMAIN is not an error here, because it means "not found".
func IsSet(name string, t uint16) (bool, error) {

	name = dns.Fqdn(name)

	msg := new(dns.Msg)
	msg.SetQuestion(name, t)

	in, err := dns.Exchange(msg, getServer())
	if err != nil {
		return false, err
	}

	switch in.Rcode {
	case 0:
		return len(in.Answer) != 0, nil
	case 3:
		// NXDOMAIN means "not found",  not an error here
		return false, nil
	default:
		return false, RcodeToError(in.Rcode)
	}
}

// IsExists checks whether a record with type A, AAAA, TXT, CNAME, MX, NS or SRV is set for name.
// NXDOMAIN is not an error here, because it means "not found".
//
// If found a setted record, this function returns without trying for the other types.
func IsExists(name string) (bool, error) {

	// A
	setA, err := IsSetA(name)
	if err != nil {
		return false, fmt.Errorf("check A failed: %w", err)
	}
	if setA {
		return true, nil
	}

	// AAAA
	setAAAA, err := IsSetAAAA(name)
	if err != nil {
		return false, fmt.Errorf("check AAAA failed: %w", err)
	}
	if setAAAA {
		return true, nil
	}

	// TXT
	setTXT, err := IsSetTXT(name)
	if err != nil {
		return false, fmt.Errorf("check TXT failed: %w", err)
	}
	if setTXT {
		return true, nil
	}

	// CNAME
	setCNAME, err := IsSetCNAME(name)
	if err != nil {
		return false, fmt.Errorf("check CNAME failed: %w", err)
	}
	if setCNAME {
		return true, nil
	}

	// MX
	setMX, err := IsSetMX(name)
	if err != nil {
		return false, fmt.Errorf("check MX failed: %w", err)
	}
	if setMX {
		return true, nil
	}

	// NS
	setNS, err := IsSetNS(name)
	if err != nil {
		return false, fmt.Errorf("check NS failed: %w", err)
	}
	if setNS {
		return true, nil
	}

	// SRV
	setSRV, err := IsSetSRV(name)
	if err != nil {
		return false, fmt.Errorf("check SRV failed: %w", err)
	}

	return setSRV, nil
}
