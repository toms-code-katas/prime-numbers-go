package primenumbers

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/cucumber/godog"
	// "github.com/cucumber/godog"
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

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func iCalculateThePrimeNumbersBetweenAnd(arg1, arg2 int) error {
	return godog.ErrPending
}

func theCalculatedPrimeNumbersShouldBe(arg1 string) error {
	return godog.ErrPending
}

func theCalculatedPrimeNumbersShouldBeException() error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I calculate the prime numbers between (-?[0-9]{0,10}) and (-?[0-9]{0,10})$`, iCalculateThePrimeNumbersBetweenAnd)
	ctx.Step(`^the calculated prime numbers should be \d.*$`, theCalculatedPrimeNumbersShouldBe)
	ctx.Step(`^the calculated prime numbers should be Exception$`, theCalculatedPrimeNumbersShouldBeException)
}
