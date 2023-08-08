package validator

import "testing"

func TestIP(t *testing.T) {

	cases := []struct {
		IP     string
		Result bool
	}{
		{IP: "0.0.0.0", Result: true},
		{IP: "255.255.255.255", Result: true},
		{IP: "0.0.0", Result: false},
		{IP: "0.0.0.256", Result: false},
		{IP: "::", Result: true},
		{IP: "0::", Result: true},
		{IP: "::x", Result: false},
		{IP: ":", Result: false},
	}

	for i := range cases {
		if r := IP(cases[i].IP); r != cases[i].Result {
			t.Fatalf("FAIL: %s is %v", cases[i].IP, r)
		}
	}
}

func TestIPv4(t *testing.T) {

	cases := []struct {
		IP     string
		Result bool
	}{
		{IP: "0.0.0.0", Result: true},
		{IP: "255.255.255.255", Result: true},
		{IP: "0.0.0", Result: false},
		{IP: "0.0.0.256", Result: false},
	}

	for i := range cases {
		if r := IPv4(cases[i].IP); r != cases[i].Result {
			t.Fatalf("FAIL: %s is %v", cases[i].IP, r)
		}
	}
}

func TestIPv6(t *testing.T) {

	cases := []struct {
		IP     string
		Result bool
	}{
		{IP: "::", Result: true},
		{IP: "0::", Result: true},
		{IP: "::x", Result: false},
		{IP: ":", Result: false},
	}

	for i := range cases {
		if r := IPv6(cases[i].IP); r != cases[i].Result {
			t.Fatalf("FAIL: %s is %v", cases[i].IP, r)
		}
	}
}

func TestPort(t *testing.T) {

	{
		cases := []struct {
			Port   string
			Result bool
		}{
			{Port: "-1", Result: false},
			{Port: "0", Result: true},
			{Port: "80", Result: true},
			{Port: "65535", Result: true},
			{Port: "65536", Result: false},
		}

		for i := range cases {
			if r := Port(cases[i].Port); r != cases[i].Result {
				t.Fatalf("FAIL: %s is %v", cases[i].Port, r)
			}
		}
	}

	{
		cases := []struct {
			Port   int
			Result bool
		}{
			{Port: -1, Result: false},
			{Port: 0, Result: true},
			{Port: 80, Result: true},
			{Port: 65535, Result: true},
			{Port: 65536, Result: false},
		}

		for i := range cases {
			if r := Port(cases[i].Port); r != cases[i].Result {
				t.Fatalf("FAIL: %d is %v", cases[i].Port, r)
			}
		}
	}

}
