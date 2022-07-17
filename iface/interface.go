package iface

import (
	"fmt"
	"net"

	elip "github.com/elmasy-com/elnet/ip"
)

// GetIPNets returns every net.IPNet address associated with iface.
func GetIPNets(iface *net.Interface) ([]net.IPNet, error) {

	if iface == nil {
		return nil, fmt.Errorf("iface is nil")
	}

	addrs, err := iface.Addrs()
	if err != nil {
		return nil, err
	}

	res := make([]net.IPNet, 0)

	for i := range addrs {

		switch v := addrs[i].(type) {
		case *net.IPNet:
			res = append(res, *v)
		default:
			// TODO: Continue?
			return nil, fmt.Errorf("unknown type: %T", v)
		}
	}

	return res, nil
}

// GetIPNets4 returns every IPv4 address associated with iface.
func GetIPNets4(iface *net.Interface) ([]net.IPNet, error) {

	nets, err := GetIPNets(iface)
	if err != nil {
		return nil, err
	}

	res := make([]net.IPNet, 0)

	for i := range nets {

		if elip.IsValid4(nets[i]) {
			res = append(res, nets[i])
		}
	}

	if len(res) == 0 {
		return res, fmt.Errorf("not found")
	}

	return res, nil
}
