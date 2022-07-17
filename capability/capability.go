package capability

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// See more: linux/capability.h
func CapToMask(x uint32) uint64 {
	return (1 << ((x) & 31))
}

func Check(cap uint32) (bool, error) {

	path := fmt.Sprintf("/proc/%d/status", os.Getpid())

	out, err := os.ReadFile(path)
	if err != nil {
		return false, err
	}

	lines := strings.Split(string(out), "\n")

	for i := range lines {

		if lines[i] == "" {
			continue
		}

		fields := strings.Fields(lines[i])

		if len(fields) != 2 {
			continue
		}

		if fields[0] != "CapEff:" {
			continue
		}

		eff, err := strconv.ParseUint(fields[1], 16, 64)
		if err != nil {
			return false, err
		}

		return eff&CapToMask(cap) != 0, err
	}

	return false, fmt.Errorf("CapEff not exist in %s", path)
}
