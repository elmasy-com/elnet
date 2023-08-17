package dns

import (
	"testing"
	"time"
)

func TestQueryAll(t *testing.T) {

	TestDomain := "elmasy.com"

	rr, errs := QueryAll(TestDomain)
	for i := range errs {
		t.Fatalf("FAIL: %s\n", errs[i])
	}

	for i := range rr {
		t.Logf("%s %s -> %s\n", TestDomain, TypeToString(rr[i].Type), rr[i].Value)
	}

}

func BenchmarkQueryAll(b *testing.B) {

	// Sleep 2 sec to not overflow the DNS server
	time.Sleep(2 * time.Second)

	srvs, err := NewServersStr(3, 1*time.Second, "8.8.8.8", "8.8.8.8", "8.8.8.8")
	if err != nil {
		b.Fatalf("FAIL: Failed to create servers: %s\n", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		srvs.QueryAll("example.com")
	}
}
