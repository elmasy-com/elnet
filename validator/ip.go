package validator

import (
	"net"
	"net/netip"
	"strconv"
)

type IPTypes interface {
	net.IP | *net.IP | netip.Addr | *netip.Addr | string | *string
}

// Returns whether v is a valid IPv4 address.
func IPv4[T IPTypes](v T) bool {

	switch t := any(v).(type) {
	case net.IP:
		return t.To4() != nil
	case *net.IP:
		return t.To4() != nil
	case netip.Addr:
		return t.Is4()
	case *netip.Addr:
		return t.Is4()
	case string:
		i := net.ParseIP(t)
		if i == nil {
			return false
		}

		return i.To4() != nil
	case *string:
		i := net.ParseIP(*t)
		if i == nil {
			return false
		}

		return i.To4() != nil
	default:
		return false
	}
}

// Returns whether v is a valid IPv6 address.
func IPv6[T IPTypes](v T) bool {

	switch t := any(v).(type) {
	case net.IP:
		return t.To16() != nil
	case *net.IP:
		return t.To16() != nil
	case netip.Addr:
		return t.Is6()
	case *netip.Addr:
		return t.Is6()
	case string:
		i := net.ParseIP(t)
		if i == nil {
			return false
		}

		return i.To16() != nil
	case *string:
		i := net.ParseIP(*t)
		if i == nil {
			return false
		}

		return i.To16() != nil
	default:
		return false
	}
}

// Returns whether v is a valid IP address.
func IP[T IPTypes](v T) bool {
	return IPv4(v) || IPv6(v)
}

type PortTypes interface {
	int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | int | uint | uintptr | string | *string
}

func Port[T PortTypes](v T) bool {

	switch t := any(v).(type) {
	case int8:
		if t >= 0 {
			return true
		}
	case uint8:
		return true
	case int16:
		if t >= 0 {
			return true
		}
	case uint16:
		return true
	case int32:
		if t >= 0 && t <= 65535 {
			return true
		}
	case uint32:
		if t <= 65535 {
			return true
		}
	case int64:
		if t >= 0 && t <= 65535 {
			return true
		}
	case uint64:
		if t <= 65535 {
			return true
		}
	case int:
		if t >= 0 && t <= 65535 {
			return true
		}
	case uint:
		if t <= 65535 {
			return true
		}
	case uintptr:
		if t <= 65535 {
			return true
		}
	case string:
		n, err := strconv.Atoi(t)
		if err != nil {
			return false
		}

		if n >= 0 && n <= 65535 {
			return true
		}
	case *string:
		n, err := strconv.Atoi(*t)
		if err != nil {
			return false
		}

		if n >= 0 && n <= 65535 {
			return true
		}
	}

	return false
}
