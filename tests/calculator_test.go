package tests

import (
	"errors"
	"math"
	"testing"

	"github.com/Yemanden0/multiagent-pilot/src/calculator"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		name       string
		a, b, want int64
		overflow   bool
	}{
		{"positive", 2, 3, 5, false},
		{"mixed", -2, 2, 0, false},
		{"upper boundary", math.MaxInt64, 0, math.MaxInt64, false},
		{"lower boundary", math.MinInt64, 0, math.MinInt64, false},
		{"upper overflow", math.MaxInt64, 1, 0, true},
		{"lower overflow", math.MinInt64, -1, 0, true},
		{"opposite signs", math.MinInt64, math.MaxInt64, -1, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := calculator.Add(tc.a, tc.b)
			if tc.overflow {
				if !errors.Is(err, calculator.ErrOverflow) {
					t.Fatalf("Add(%d, %d) error = %v, want ErrOverflow", tc.a, tc.b, err)
				}
				return
			}
			if err != nil || got != tc.want {
				t.Fatalf("Add(%d, %d) = %d, %v; want %d, nil", tc.a, tc.b, got, err, tc.want)
			}
		})
	}
}
