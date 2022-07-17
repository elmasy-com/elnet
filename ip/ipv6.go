package ip

import (
	"math/rand"
	"net"
	"strings"
)

// ReservedIPv6 is a collection of reserved IPv6 addresses.
var ReservedIPv6 = []net.IPNet{
	{IP: net.IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Mask: net.IPMask{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255}}, // ::/128, Unspecified Address
	{IP: net.IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Mask: net.IPMask{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255}}, // ::1/128, Loopback Address
	{IP: net.IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 0, 0, 0, 0}, Mask: net.IPMask{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 0, 0, 0, 0}},     // ::ffff:0:0/96, IPv4-mapped addresses
	{IP: net.IP{0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 0, 0, 0, 0, 0, 0}, Mask: net.IPMask{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 0, 0, 0, 0}},     // ::ffff:0:0:0/96, IPv4 translated addresses
	{IP: net.IP{0, 100, 255, 155, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Mask: net.IPMask{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 0, 0, 0, 0}},   // 64:ff9b::/96, IPv4-IPv6 Translat.
	{IP: net.IP{0, 100, 255, 155, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Mask: net.IPMask{255, 255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},               // 64:ff9b:1::/48, IPv4-IPv6 Translat.
	{IP: net.IP{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Mask: net.IPMask{255, 255, 255, 255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0}},                 // 100::/64, Discard-Only Address Block
	{IP: net.IP{32, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Mask: net.IPMask{255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},                        // 2001::/32, IETF Protocol Assignments
	{IP: net.IP{32, 1, 0, 32, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Mask: net.IPMask{255, 255, 255, 240, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},                       // 2001:20::/28, ORCHIDv2
	{IP: net.IP{32, 1, 13, 184, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Mask: net.IPMask{255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},                     // 2001:db8::/32, Documentation
	{IP: net.IP{32, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Mask: net.IPMask{255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},                            // 2002::/16, 6to4
	{IP: net.IP{252, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Mask: net.IPMask{254, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},                             // fc00::/7, Unique-Local
	{IP: net.IP{254, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Mask: net.IPMask{255, 192, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},                         // fe80::/10, Link-Local Unicast
	{IP: net.IP{255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Mask: net.IPMask{255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},                             // ff00::/8, Multicast
}

// IsReserved6 checks if the given IPv6 address is reserved.
func IsReserved6(ip net.IP) bool {
	for i := range ReservedIPv6 {
		if ReservedIPv6[i].Contains(ip) {
			return true
		}
	}
	return false
}

// IsValid6 checks whether ip is valid IPv6 address.
func IsValid6[T IPTypes](ip T) bool {

	// Check the string. IPv6 address should contains ':'.
	if v, ok := any(ip).(string); ok {
		if !strings.Contains(v, ":") {
			return false
		}
	}
	if v, ok := any(ip).(*string); ok {
		if !strings.Contains(*v, ":") {
			return false
		}
	}

	i := convertToIP(ip)
	if i == nil {
		return false
	}

	return i.To16() != nil
}

// GetRandom6 is return a random IPv6 address.
// The returned IP *can be* a reserved address.
func GetRandom6() net.IP {

	bytes := make([]byte, 16)

	rand.Read(bytes)

	return net.IP{
		bytes[0], bytes[1], bytes[2], bytes[3],
		bytes[4], bytes[5], bytes[6], bytes[7],
		bytes[8], bytes[9], bytes[10], bytes[11],
		bytes[12], bytes[13], bytes[14], bytes[15]}
}

// GetPublic6 is return a *non reserved* IPv6 address.
func GetPublic6() net.IP {

	for {
		ip := GetRandom6()

		if !IsReserved6(ip) {
			return ip
		}
	}
}
