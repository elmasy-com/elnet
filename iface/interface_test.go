package iface

import (
	"net"
	"testing"
)

func TestGetInterfaceIPNets(t *testing.T) {

	ifaces, err := net.Interfaces()
	if err != nil {
		t.Errorf("Failed to get ifcae: %s\n", err)
	}

	for i := range ifaces {
		nets, err := GetIPNets(&ifaces[i])
		if err != nil {
			t.Errorf("Failed to get nets: %s\n", err)
		}

		for v := range nets {
			t.Logf("%s -> %s\n", ifaces[i].Name, nets[v].String())
		}
	}
}
