package capability

import (
	"os"
	"testing"
)

func TestCheck(t *testing.T) {

	ok, err := Check(CAP_CHOWN)
	if err != nil {
		t.Errorf("Failed to check: %s\n", err)
	}

	if os.Geteuid() == 0 && !ok {
		t.Errorf("CAP_CHOWN must be set")
	}
}
