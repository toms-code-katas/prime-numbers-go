package primenumbers_test

import (
	"reflect"
	"testing"

	primenumbers "prime-numbers-go"
)

func TestCalculatePrimes(t *testing.T) {
	t.Parallel()

	primes := []uint64{2, 3, 5, 7}
	calculatedPrimes := primenumbers.CalculatePrimes(1, 10)

	if !reflect.DeepEqual(primes, calculatedPrimes) {
		t.Fatalf("expected: %v, got: %v", primes, calculatedPrimes)
	}
}

func TestCalculatePrimesUsingTable(t *testing.T) {
	t.Parallel()

	type test struct {
		start  uint64
		stop   uint64
		primes []uint64
	}

	tests := []test{
		{start: 1, stop: 10, primes: []uint64{2, 3, 5, 7}},
		{start: 1, stop: 25, primes: []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23}},
	}

	for _, tc := range tests {
		calculatedPrimes := primenumbers.CalculatePrimes(tc.start, tc.stop)
		if !reflect.DeepEqual(tc.primes, calculatedPrimes) {
			t.Fatalf("expected: %v, got: %v", tc.primes, calculatedPrimes)
		}
	}
}
