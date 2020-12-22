// +build !linux

package platforms

import (
	"errors"
)

// IsNoLimit returns true if the given limit (returned by Limit) is the magic number which indicates that no limit is set.
func IsNoLimit(limit int64) bool {
	return false
}

// Usage returns the current memory usage.
func Usage() (int64, error) {
	return 0, errors.New("platform not supported")
}

// Limit returns the current memory limit.
func Limit() (int64, error) {
	return 0, errors.New("platform not supported")
}
