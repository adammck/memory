package memory

import (
	"github.com/adammck/memory/internal/platforms"
)

// IsNoLimit returns true if the given limit (returned by Limit) is the magic number which indicates that no limit is set.
func IsNoLimit(limit int64) bool {
	return platforms.IsNoLimit(limit)
}

// Usage returns the current memory usage.
func Usage() (int64, error) {
	return platforms.Usage()
}

// Limit returns the current memory limit.
func Limit() (int64, error) {
	return platforms.Limit()
}
