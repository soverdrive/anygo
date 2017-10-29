package simple

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	test := []struct {
		inputA   int64
		inputB   int64
		expected uint64
	}{
		{1, 1, 2},
		{math.MaxInt64, 1, uint64(math.MaxInt64) + 1},
	}

	for _, v := range test {
		output := Add(v.inputA, v.inputB)
		if uint64(output) != v.expected {
			t.Errorf("Error while testing, expected %d, got %d", v.expected, output)
		}
	}
}

func TestAddUnsigned(t *testing.T) {
	test := []struct {
		inputA   uint64
		inputB   uint64
		expected uint64
	}{
		{1, 1, 2},
		{math.MaxInt64 + 1, math.MaxInt64 + 1, 0},
	}

	for _, v := range test {
		output := AddUnsigned(v.inputA, v.inputB)
		if output != v.expected {
			t.Errorf("Error while testing, expected %d, got %d", v.expected, output)
		}
	}
}
