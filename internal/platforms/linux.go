// +build linux

package platforms

import (
	"errors"

	"github.com/adammck/memory/internal/cgroups"
)

const (
	subsysMem = "memory"
	fnLimit   = "memory.limit_in_bytes"
	fnUsage   = "memory.usage_in_bytes"
)

// IsNoLimit returns true if the given limit (returned by Limit) is the magic number which indicates that no limit is set.
func IsNoLimit(limit int64) bool {
	return limit == 9223372036854771712
}

func get(filename string) (int64, error) {
	cg, err := cgroups.NewCGroupsForCurrentProcess()
	if err != nil {
		return 0, err
	}

	memGroup, exists := cg[subsysMem]
	if !exists {
		// This shouldn't happen
		return 0, errors.New("memory subsystem not present")
	}

	usage, err := memGroup.ReadInt(filename)
	if err != nil {
		return 0, err
	}

	if usage == 0 {
		return 0, errors.New("usage not found")
	}

	return int64(usage), nil
}

// Usage returns the current memory usage.
func Usage() (int64, error) {
	return get(fnUsage)
}

// Limit returns the current memory limit.
func Limit() (int64, error) {
	return get(fnLimit)
}
