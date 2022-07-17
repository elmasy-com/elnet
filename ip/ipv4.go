package ip

import (
	"math/rand"
	"net"
	"strings"
)

// ReservedIPv4 is a collection of reserved IPv4 addresses.
// Source: https://www.iana.org/assignments/iana-ipv4-special-registry/iana-ipv4-special-registry.xhtml
var ReservedIPv4 = []net.IPNet{
	{IP: net.IPv4(0, 0, 0, 0), Mask: net.IPv4Mask(255, 0, 0, 0)},               // 0.0.0.0/8, "This" network
	{IP: net.IPv4(10, 0, 0, 0), Mask: net.IPv4Mask(255, 0, 0, 0)},              // 10.0.0.0/8, Class A private network
	{IP: net.IPv4(100, 64, 0, 0), Mask: net.IPv4Mask(255, 192, 0, 0)},          // 100.64.0.0/10, Carrier-grade NAT
	{IP: net.IPv4(127, 0, 0, 0), Mask: net.IPv4Mask(255, 0, 0, 0)},             // 127.0.0.0/8, Loopback
	{IP: net.IPv4(169, 254, 0, 0), Mask: net.IPv4Mask(255, 255, 0, 0)},         // 169.254.0.0/16, Link local
	{IP: net.IPv4(172, 16, 0, 0), Mask: net.IPv4Mask(255, 240, 0, 0)},          // 172.16.0.0/12, Class B private network
	{IP: net.IPv4(192, 0, 0, 0), Mask: net.IPv4Mask(255, 255, 255, 0)},         // 192.0.0.0/24, IETF protocol assignments
	{IP: net.IPv4(192, 0, 2, 0), Mask: net.IPv4Mask(255, 255, 255, 0)},         // 192.0.2.0/24, TEST-NET-1
	{IP: net.IPv4(192, 88, 99, 0), Mask: net.IPv4Mask(255, 255, 255, 0)},       // 192.88.99.0/24, Reserved, formerly IPv6 to IPv4
	{IP: net.IPv4(192, 168, 0, 0), Mask: net.IPv4Mask(255, 255, 0, 0)},         // 192.168.0.0/24, Class C private network
	{IP: net.IPv4(198, 18, 0, 0), Mask: net.IPv4Mask(255, 254, 0, 0)},          // 198.18.0.0/15, Benchmarking
	{IP: net.IPv4(198, 51, 100, 0), Mask: net.IPv4Mask(255, 255, 255, 0)},      // 198.51.100.0/24, TEST-NET-2
	{IP: net.IPv4(203, 0, 113, 0), Mask: net.IPv4Mask(255, 255, 255, 0)},       // 203.0.113.0/24, TEST-NET-3
	{IP: net.IPv4(224, 0, 0, 0), Mask: net.IPv4Mask(240, 0, 0, 0)},             // 224.0.0.0/4, Multicast
	{IP: net.IPv4(233, 252, 0, 0), Mask: net.IPv4Mask(255, 255, 255, 0)},       // 233.252.0.0/24 , MCAST-TEST-NET
	{IP: net.IPv4(240, 0, 0, 0), Mask: net.IPv4Mask(240, 0, 0, 0)},             // 240.0.0.0/4, Reserved for future use
	{IP: net.IPv4(255, 255, 255, 255), Mask: net.IPv4Mask(255, 255, 255, 255)}, // 255.255.255.255/32, Broadcast
}

// IsReserved4 checks if the given IPv4 address is reserved.
func IsReserved4(ip net.IP) bool {
	for i := range ReservedIPv4 {
		if ReservedIPv4[i].Contains(ip) {
			return true
		}
	}
	return false
}

// IsValid4 checks whether ip is valid IPv4 address.
func IsValid4[T IPTypes](ip T) bool {

	// Check the string. IPv4 address should contains '.'.
	if v, ok := any(ip).(string); ok {
		if !strings.Contains(v, ".") {
			return false
		}
	}
	if v, ok := any(ip).(*string); ok {
		if !strings.Contains(*v, ".") {
			return false
		}
	}

	i := convertToIP(ip)
	if i == nil {
		return false
	}

	return i.To4() != nil
}

// GetRandom4 is return a random IPv4 address.
// The returned IP *can be* a reserved address.
func GetRandom4() net.IP {

	bytes := make([]byte, 4)

	rand.Read(bytes)

	return net.IPv4(bytes[0], bytes[1], bytes[2], bytes[3])
}

// GetPublic4 is return a *non reserved* IPv4 address.
func GetPublic4() net.IP {

	for {
		ip := GetRandom4()

		if !IsReserved4(ip) {
			return ip
		}
	}
}
