// Package calculator provides checked arithmetic on signed 64-bit integers.
package calculator

import (
	"errors"
	"math"
)

// ErrOverflow indicates that the exact result cannot be represented as int64.
var ErrOverflow = errors.New("integer overflow")

// Add returns the exact sum, or zero and ErrOverflow if it is out of range.
func Add(a, b int64) (int64, error) {
	if (b > 0 && a > math.MaxInt64-b) || (b < 0 && a < math.MinInt64-b) {
		return 0, ErrOverflow
	}
	return a + b, nil
}
