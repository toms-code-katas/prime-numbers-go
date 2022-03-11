package primenumbers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCalculatePrimes(t *testing.T) {
	expected_primes := []uint64{2, 3, 5, 7}
	primes := CalculatePrimes(1, 10)
	for i, prime := range primes {
		if expected_primes[i] != prime {
			t.Fatal(fmt.Sprintf("Expected primes %v and calculated primes %v are not equal", expected_primes, primes))
		}
	}
}

func TestCalculatePrimesUsingTable(t *testing.T) {

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
		calculatedPrimes := CalculatePrimes(tc.start, tc.stop)
		if !reflect.DeepEqual(tc.primes, calculatedPrimes) {
			t.Fatalf("expected: %v, got: %v", tc.primes, calculatedPrimes)
		}
	}
}
