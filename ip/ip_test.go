package ip

import (
	"net"
	"testing"
)

func TestRandom(t *testing.T) {

	var (
		rV4   = 0
		pV4   = 0
		rV6   = 0
		pV6   = 0
		total = 1000000
	)

	for i := 0; i < total; i++ {

		ip := GetRandom()

		switch {
		case IsValid4(ip):
			if IsReserved4(ip) {
				rV4++
			} else {
				pV4++
			}
		case IsValid6(ip):
			if IsReserved6(ip) {
				rV6++
			} else {
				pV6++
			}
		default:
			t.Errorf("Not a valid ip: %#v\n", ip)
		}
	}

	t.Logf("Total -> v4 %d (%.3f%%) | v6 %d (%.3f%%) of %d", rV4+pV4, float64(rV4+pV4)/float64(total)*100, rV6+pV6, float64(rV6+pV6)/float64(total)*100, total)
	t.Logf("IPv4  -> reserved %d / public %d = (%.2f%%)", rV4, pV4, float64(rV4)/float64(pV4)*100)
	t.Logf("IPv6  -> reserved %d / public %d = (%.2f%%)", rV6, pV6, float64(rV6)/float64(pV6)*100)
}

func TestIsLAN(t *testing.T) {

	ip := net.IPv4(1, 1, 1, 1)
	if ip == nil {
		t.Errorf("ip is nil\n")
	}

	if IsLAN(ip) {
		t.Errorf("%s is not in LAN\n", ip)
	}
}