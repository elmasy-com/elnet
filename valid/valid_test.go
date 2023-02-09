package valid

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

func TestDomain(t *testing.T) {

	// Generate a string with 400 x 'a'
	longStringFunc := func() string {
		s := make([]byte, 400)
		for i := 0; i < 400; i++ {
			s = append(s, 'a')
		}
		return string(s)
	}

	cases := []struct {
		Domain string
		Result bool
	}{
		{Domain: "elmasy.com", Result: true},
		{Domain: "elmasy.com.", Result: true},
		{Domain: ".elmasy.com", Result: false},
		{Domain: ".elmasy.com.", Result: false},
		{Domain: "aaaaaa", Result: false},
		{Domain: longStringFunc(), Result: false},
		{Domain: "", Result: false},
		{Domain: ".", Result: false},
		{Domain: "a a", Result: false},
		{Domain: "a=a", Result: false},
		{Domain: " ", Result: false},
		{Domain: "*", Result: false},
	}

	for i := range cases {
		if r := Domain(cases[i].Domain); r != cases[i].Result {
			t.Fatalf("FAIL: %s is %v", cases[i].Domain, r)
		}
	}
}

func TestURL(t *testing.T) {

	cases := []struct {
		URL    string
		Result bool
	}{
		{URL: "", Result: false},
		{URL: "test", Result: false},
		{URL: "test.com", Result: false},
		{URL: "https://test.com", Result: true},
		{URL: "https://test.com/path", Result: true},
	}

	for i := range cases {
		if r := URL(cases[i].URL); r != cases[i].Result {
			t.Fatalf("FAIL: %s is %v", cases[i].URL, r)
		}
	}
}

func TestEmail(t *testing.T) {

	cases := []struct {
		URL    string
		Result bool
	}{

		{URL: "simple@example.com", Result: true},
		{URL: "very.common@example.com", Result: true},
		{URL: "disposable.style.email.with+symbol@example.com", Result: true},
		{URL: "other.email-with-hyphen@example.com", Result: true},
		{URL: "fully-qualified-domain@example.com", Result: true},
		{URL: "user.name+tag+sorting@example.com", Result: true},
		{URL: "x@example.com", Result: true},
		{URL: "example-indeed@strange-example.com", Result: true},
		{URL: "test/test@test.com", Result: true},

		{URL: "test", Result: false},
		{URL: "test.com", Result: false},
		{URL: "@test.com", Result: false},
		{URL: ".a@test.com", Result: false},
		{URL: "a.@test.com", Result: false},
		{URL: "a..a@test.com", Result: false},
		{URL: "a@test", Result: false},
		{URL: "A@b@c@example.com", Result: false},
		{URL: "a\"b(c)d,e:f;g<h>i[j\\k]l@example.com", Result: false},
		{URL: "just\"not\"right@example.com", Result: false},
		{URL: "this is\"not\\allowed@example.com", Result: false},
		{URL: "1234567890123456789012345678901234567890123456789012345678901234+x@example.com", Result: false},
	}

	for i := range cases {
		if r := Email(cases[i].URL); r != cases[i].Result {
			t.Fatalf("FAIL: %s is %v", cases[i].URL, r)
		}
	}
}
