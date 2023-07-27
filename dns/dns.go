package dns

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/elmasy-com/identify"
	"github.com/miekg/dns"
)

// Default client configuration.
//
// Set in the init() function to read resolv.conf or use Cloudflare + Google + Quad9.
var Conf *dns.ClientConfig

var Client *dns.Client

// MaxRetries is the default number of retries of failed queries.
// Must be greater than 1, else functions will fail with ErrInvalidMaxRetries.
var MaxRetries int = 5

// Initialize Conf.
// If reading resolv.conf failed, use Cloudflare + Google
func init() {

	var err error

	UpdateClient("udp", 2*time.Second)

	Conf, err = dns.ClientConfigFromFile("/etc/resolv.conf")
	if err == nil {
		return
	}

	UpdateConf([]string{"1.1.1.1", "1.0.0.1", "8.8.8.8", "8.8.4.4", "9.9.9.10", "149.112.112.10"}, "53")
}

// getServer returns a DNS server to use.
// If index is between the servers range, returns the selected server.
// Else, returns a random one.
//
// Set index to -1 to get a random one.
func getServer(index int) string {

	var r string

	switch l := len(Conf.Servers); l {
	case 0:
		// No server is configured
		r = "1.1.1.1"
	case 1:
		// Only ne server
		r = Conf.Servers[0]
	default:

		if index >= 0 && index < l {
			// If index is between the servers range, return the selected server
			r = Conf.Servers[index]
		} else {
			// If not in range, return a random server
			r = Conf.Servers[rand.Intn(len(Conf.Servers))]
		}
	}

	if identify.IsValidIPv6(r) {
		r = "[" + r + "]"
	}

	return r + ":" + Conf.Port
}

// Create a new ClientConfig for Conf.
func UpdateConf(servers []string, port string) {

	Conf = new(dns.ClientConfig)

	Conf.Servers = servers
	Conf.Search = make([]string, 0)
	Conf.Port = port
	Conf.Ndots = 1
	Conf.Timeout = 5
	Conf.Attempts = 2
}

// Create a new Client for default client.
// net must be "udp", "tcp" or "tcp-tls".
func UpdateClient(net string, timeout time.Duration) error {

	if net != "udp" && net != "tcp" && net != "tcp-tls" {
		return fmt.Errorf("invalid net: %s", net)
	}

	Client = &dns.Client{Net: net, Timeout: timeout}

	return nil
}

// Generic query for type t to server s.
// Returns the Answer section.
// In case of error, the answer will be nil and return ErrX or any unknown error.
func QueryServer(name string, t uint16, s string) ([]dns.RR, error) {

	name = dns.Fqdn(name)

	msg := new(dns.Msg)
	msg.SetQuestion(name, t)

	in, _, err := Client.Exchange(msg, s)
	if err != nil {
		return nil, err
	}

	if in.Rcode == 0 {
		return in.Answer, nil
	}

	return nil, RcodeToError(in.Rcode)
}

// Generic query for type t.
// Returns the Answer section.
// In case of error, the answer will be nil and return ErrX or any unknown error.
func Query(name string, t uint16) ([]dns.RR, error) {

	return QueryServer(name, t, getServer(-1))
}

// IsSet checks whether a record with type t is set for name.
// This function retries the query in case of error up to MaxRetries times.
// NXDOMAIN is not an error here, because it means "not found".
func IsSet(name string, t uint16) (bool, error) {

	err := ErrInvalidMaxRetries
	var in *dns.Msg

	name = dns.Fqdn(name)

	for i := 0; i < MaxRetries; i++ {

		msg := new(dns.Msg)
		msg.SetQuestion(name, t)

		in, _, err = Client.Exchange(msg, getServer(i))
		if err == nil {
			break
		}
	}

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

// IsExists checks whether a record with type A, AAAA, TXT, CNAME, MX, NS, CAA or SRV is set for name.
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

	// CAA
	setCAA, err := IsSetCAA(name)
	if err != nil {
		return false, fmt.Errorf("chack CAA failed: %w", err)
	}
	if setCAA {
		return true, nil
	}

	// SRV
	setSRV, err := IsSetSRV(name)
	if err != nil {
		return false, fmt.Errorf("check SRV failed: %w", err)
	}

	return setSRV, nil
}
