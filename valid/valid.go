package valid

import (
	"github.com/elmasy-com/elnet/domain"
	"github.com/elmasy-com/elnet/ip"
	"github.com/elmasy-com/elnet/url"
)

// Returns whether v is a valid IP address.
func IP[T ip.IPTypes](v T) bool {
	return ip.IsValid(v)
}

// Returns whether v is a valid IPv4 address.
func IPv4[T ip.IPTypes](v T) bool {
	return ip.IsValid4(v)
}

// Returns whether v is a valid IPv6 address.
func IPv6[T ip.IPTypes](v T) bool {
	return ip.IsValid6(v)
}

// Returns whether v is a valid domain.
func Domain(v string) bool {
	return domain.IsValid(v)
}

// Returns whether v is a valid URL.
func URL(v string) bool {

	_, err := url.Parse(v)

	return err == nil
}
