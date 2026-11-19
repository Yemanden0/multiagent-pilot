package calculator

import (
	"errors"
	"math"
	"math/bits"
)

// ErrInvalidExponent indicates that the exponent is not positive.
var ErrInvalidExponent = errors.New("exponent must be positive")

// ErrNegativeExponent is an alias for ErrInvalidExponent for callers that
// specifically handle negative exponents.
var ErrNegativeExponent = ErrInvalidExponent

// Pow returns base raised to exponent, or zero and ErrOverflow if the exact
// result is outside the int64 range. The exponent must be positive.
//
// Negative bases are accepted for both even and odd exponents; an odd
// exponent produces a negative result. If exponent is not positive, Pow
// returns ErrInvalidExponent.
func Pow(base, exponent int64) (int64, error) {
	if exponent <= 0 {
		return 0, ErrInvalidExponent
	}

	switch base {
	case 0:
		return 0, nil
	case 1:
		return 1, nil
	case -1:
		if exponent%2 == 0 {
			return 1, nil
		}
		return -1, nil
	}

	negative := base < 0 && exponent%2 != 0

	var factor uint64
	if base < 0 {
		factor = uint64(-(base + 1)) + 1
	} else {
		factor = uint64(base)
	}

	result := uint64(1)
	remaining := uint64(exponent)
	for remaining > 0 {
		if remaining&1 != 0 {
			hi, lo := bits.Mul64(result, factor)
			if hi != 0 {
				return 0, ErrOverflow
			}
			result = lo
		}

		remaining >>= 1
		if remaining == 0 {
			break
		}

		hi, lo := bits.Mul64(factor, factor)
		if hi != 0 {
			return 0, ErrOverflow
		}
		factor = lo
	}

	if negative {
		if result > uint64(1)<<63 {
			return 0, ErrOverflow
		}
		if result == uint64(1)<<63 {
			return math.MinInt64, nil
		}
		return -int64(result), nil
	}

	if result > uint64(math.MaxInt64) {
		return 0, ErrOverflow
	}
	return int64(result), nil
}

// Power is an alias for Pow.
func Power(base, exponent int64) (int64, error) {
	return Pow(base, exponent)
}
