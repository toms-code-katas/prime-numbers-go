package primenumbers

import (
	"fmt"
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
