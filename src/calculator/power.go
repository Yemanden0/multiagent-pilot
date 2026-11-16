package calculator

import (
	"errors"
	"math/bits"
)

// ErrInvalidExponent indicates that the exponent is not positive.
var ErrInvalidExponent = errors.New("exponent must be positive")

// ErrNegativeExponent is an alias for ErrInvalidExponent for callers that
// specifically handle negative exponents.
var ErrNegativeExponent = ErrInvalidExponent

// Pow returns base raised to exponent. The exponent must be positive.
//
// The exact result is returned as uint64 when it fits. A negative base is
// accepted only when the exponent is even; an odd exponent produces a
// negative mathematical result, which cannot be represented as uint64 and
// returns ErrOverflow. If exponent is not positive, Pow returns
// ErrInvalidExponent.
func Pow(base, exponent int64) (uint64, error) {
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
		return 0, ErrOverflow
	}

	if base < 0 && exponent%2 != 0 {
		return 0, ErrOverflow
	}

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

	return result, nil
}

// Power is an alias for Pow.
func Power(base, exponent int64) (uint64, error) {
	return Pow(base, exponent)
}
