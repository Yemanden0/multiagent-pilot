package tests

import (
	"errors"
	"math"
	"testing"

	"github.com/Yemanden0/multiagent-pilot/src/calculator"
)

func TestPow(t *testing.T) {
	cases := []struct {
		name     string
		base     int64
		exponent int64
		want     uint64
		wantErr  error
	}{
		{"positive exponent", 2, 10, 1024, nil},
		{"zero base", 0, 5, 0, nil},
		{"one base", 1, math.MaxInt64, 1, nil},
		{"even negative base", -2, 4, 16, nil},
		{"max int64 first power", math.MaxInt64, 1, uint64(math.MaxInt64), nil},
		{"uint64 boundary", 2, 63, 1 << 63, nil},
		{"negative exponent", 2, -1, 0, calculator.ErrInvalidExponent},
		{"zero exponent", 2, 0, 0, calculator.ErrInvalidExponent},
		{"uint64 overflow", 2, 64, 0, calculator.ErrOverflow},
		{"odd negative base", -2, 3, 0, calculator.ErrOverflow},
		{"max int64 squared", math.MaxInt64, 2, 0, calculator.ErrOverflow},
		{"min int64 first power", math.MinInt64, 1, 0, calculator.ErrOverflow},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := calculator.Pow(tc.base, tc.exponent)
			if tc.wantErr != nil {
				if !errors.Is(err, tc.wantErr) {
					t.Fatalf("Pow(%d, %d) error = %v, want %v", tc.base, tc.exponent, err, tc.wantErr)
				}
				if got != 0 {
					t.Fatalf("Pow(%d, %d) = %d on error, want 0", tc.base, tc.exponent, got)
				}
				return
			}
			if err != nil || got != tc.want {
				t.Fatalf("Pow(%d, %d) = %d, %v; want %d, nil", tc.base, tc.exponent, got, err, tc.want)
			}
		})
	}
}

func TestPowerAlias(t *testing.T) {
	got, err := calculator.Power(3, 4)
	if err != nil || got != 81 {
		t.Fatalf("Power(3, 4) = %d, %v; want 81, nil", got, err)
	}
}

func TestPowNegativeExponentAlias(t *testing.T) {
	_, err := calculator.Pow(2, -1)
	if !errors.Is(err, calculator.ErrNegativeExponent) {
		t.Fatalf("Pow(2, -1) error = %v, want ErrNegativeExponent", err)
	}
}
